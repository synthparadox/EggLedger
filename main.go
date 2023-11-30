package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/DavidArthurCole/EggLedger/db"
	"github.com/DavidArthurCole/EggLedger/ei"
	"github.com/DavidArthurCole/EggLedger/forkedlorca"
	humanize "github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/sync/semaphore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	//go:embed VERSION
	_appVersion string

	//go:embed www
	_fs embed.FS

	_rootDir     string
	_internalDir string

	_appIsInForbiddenDirectory bool

	// macOS Gateway security feature which executed apps with xattr
	// com.apple.quarantine in certain locations (like ~/Downloads) in a jailed
	// readonly environment. The jail looks like:
	// /private/var/folders/<...>/<...>/T/AppTranslocation/<UUID>/d/internal
	_appIsTranslocated bool

	_devMode   = os.Getenv("DEV_MODE") != ""
	dateFormat = "%02d:%02d:%02d"
)

const (
	_requestInterval = 3 * time.Second
)

type UI struct {
	forkedlorca.UI
}

func (u UI) MustLoad(url string) {
	err := u.Load(url)
	if err != nil {
		log.Fatal(err)
	}
}

func (u UI) MustBind(name string, f interface{}) {
	err := u.Bind(name, f)
	if err != nil {
		log.Fatal(err)
	}
}

type AppState string

const (
	AppState_AWAITING_INPUT    AppState = "AwaitingInput"
	AppState_FETCHING_SAVE     AppState = "FetchingSave"
	AppState_FETCHING_MISSIONS AppState = "FetchingMissions"
	AppState_EXPORTING_DATA    AppState = "ExportingData"
	AppState_SUCCESS           AppState = "Success"
	AppState_FAILED            AppState = "Failed"
	AppState_INTERRUPTED       AppState = "Interrupted"
)

type MissionProgress struct {
	Total                   int     `json:"total"`
	Finished                int     `json:"finished"`
	FinishedPercentage      string  `json:"finishedPercentage"`
	ExpectedFinishTimestamp float64 `json:"expectedFinishTimestamp"`
}

type worker struct {
	*semaphore.Weighted
	ctx     context.Context
	cancel  context.CancelFunc
	ctxlock sync.Mutex
}

type ExportedFile struct {
	File     string `json:"file"`
	Count    int    `json:"count"`
	DateTime string `json:"datetime"`
	EID      string `json:"eid"`
}

type LoadedMission struct {
	LaunchDay    int32                        `json:"launchDay"`
	LaunchMonth  int32                        `json:"launchMonth"`
	LaunchYear   int32                        `json:"launchYear"`
	LaunchTime   string                       `json:"launchTime"`
	ReturnDay    int32                        `json:"returnDay"`
	ReturnMonth  int32                        `json:"returnMonth"`
	ReturnYear   int32                        `json:"returnYear"`
	ReturnTime   string                       `json:"returnTime"`
	MissiondId   string                       `json:"missionId"`
	Ship         *ei.MissionInfo_Spaceship    `json:"ship"`
	DurationType *ei.MissionInfo_DurationType `json:"durationType"`
	Level        int32                        `json:"level"`
	Capacity     int32                        `json:"capacity"`
	Target       string                       `json:"target"`
}

type MissionDrop struct {
	SpecType string `json:"specType"`
	Name     string `json:"name"`
	Level    int32  `json:"level"`
	Rarity   int32  `json:"rarity"`
}

type ExportAccount struct {
	Id           string `json:"id"`
	Nickname     string `json:"nickname"`
	MissionCount int    `json:"missionCount"`
}

func init() {
	log.SetLevel(log.InfoLevel)
	// Send a copy of logs to $TMPDIR/EggLedger.log in case the app crashes
	// before we can even set up persistent logging.
	tmplog, err := os.OpenFile(filepath.Join(os.TempDir(), "EggLedger.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err)
	}
	log.AddHook(&writer.Hook{
		Writer:    tmplog,
		LogLevels: log.AllLevels,
	})

	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		log.Fatal(err)
	}
	_rootDir = filepath.Dir(path)
	if runtime.GOOS == "darwin" {
		// Locate parent dir of app bundle if we're inside a Mac app.
		parent, dir1 := filepath.Split(_rootDir)
		parent = filepath.Clean(parent)
		parent, dir2 := filepath.Split(parent)
		parent = filepath.Clean(parent)
		if dir1 == "MacOS" && dir2 == "Contents" && strings.HasSuffix(parent, ".app") {
			_rootDir = filepath.Dir(parent)
		}
	}
	log.Infof("root dir: %s", _rootDir)

	_internalDir = filepath.Join(_rootDir, "internal")

	// Make sure the app isn't located in the system/user app directory or
	// downloads dir.
	if runtime.GOOS == "darwin" {
		if _rootDir == "/Applications" {
			_appIsInForbiddenDirectory = true
		} else {
			pattern := regexp.MustCompile(`^/Users/[^/]+/(Applications|Downloads)$`)
			if pattern.MatchString(_rootDir) {
				_appIsInForbiddenDirectory = true
			}
		}
	} else {
		// On non-macOS platforms, just check whether the root dir ends in "/Downloads".
		pattern := regexp.MustCompile(`[\/]Downloads$`)
		if pattern.MatchString(_rootDir) {
			_appIsInForbiddenDirectory = true
		}
	}
	if _appIsInForbiddenDirectory {
		log.Error("app is in a forbidden directory")
		return
	}

	if runtime.GOOS == "darwin" {
		if strings.HasPrefix(_rootDir, "/private/var/folders/") {
			_appIsTranslocated = true
		}
	}
	if _appIsTranslocated {
		log.Error("app is translocated")
		return
	}

	if err := os.MkdirAll(_internalDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := hide(_internalDir); err != nil {
		log.Errorf("error hiding internal directory: %s", err)
	}

	// Set up persistent logging.
	logdir := filepath.Join(_rootDir, "logs")
	if err := os.MkdirAll(logdir, 0755); err != nil {
		log.Error(err)
	} else {
		logfile := filepath.Join(logdir, "app.log")
		logger := &lumberjack.Logger{
			Filename:  logfile,
			MaxSize:   5, // megabytes
			MaxAge:    7, // days
			LocalTime: true,
			Compress:  true,
		}
		log.AddHook(&writer.Hook{
			Writer:    logger,
			LogLevels: log.AllLevels,
		})
	}

	storageInit()
	dataInit()
}

func viewMissionsOfId(eid string) (string, error) {

	//Get list of complete missions from the DB
	completeMissions, err := db.RetrievePlayerCompleteMissions(eid)
	if err != nil {
		log.Error(err)
		return "", err
	}

	//Array of FileMission
	missionArr := []LoadedMission{}

	for _, completeMission := range completeMissions {

		launchDateTimeObject := time.Unix(int64(*completeMission.Info.StartTimeDerived), 0)
		ltH, ltM, ltS := launchDateTimeObject.Clock()
		launchTime := fmt.Sprintf(dateFormat, ltH, ltM, ltS)
		returnTimeObject := launchDateTimeObject.Add(time.Duration(*completeMission.Info.DurationSeconds * float64(time.Second)))
		rtH, rtM, rtS := returnTimeObject.Clock()
		returnTime := fmt.Sprintf(dateFormat, rtH, rtM, rtS)

		missionInst := LoadedMission{
			LaunchDay:   int32(launchDateTimeObject.Day()),
			LaunchMonth: int32(launchDateTimeObject.Month()),
			LaunchYear:  int32(launchDateTimeObject.Year()),
			LaunchTime:  launchTime,

			ReturnDay:   int32(returnTimeObject.Day()),
			ReturnMonth: int32(returnTimeObject.Month()),
			ReturnYear:  int32(returnTimeObject.Year()),
			ReturnTime:  returnTime,

			MissiondId:   *completeMission.Info.Identifier,
			Ship:         completeMission.Info.Ship,
			DurationType: completeMission.Info.DurationType,
			Level:        int32(*completeMission.Info.Level),
			Capacity:     int32(*completeMission.Info.Capacity),
			Target:       properTargetName(completeMission.Info.TargetArtifact),
		}
		missionArr = append(missionArr, missionInst)
	}

	// Convert the array of FileMissionYear to a JSON string
	jsonData, err := json.Marshal(missionArr)
	if err != nil {
		return "", err
	}

	// Return the JSON string
	return string(jsonData), nil
}

func properTargetName(name *ei.ArtifactSpec_Name) string {
	if name == nil || *name == ei.ArtifactSpec_UNKNOWN {
		return ""
	} else {
		return ei.ArtifactSpec_Name_name[int32(*name)]
	}
}

func main() {
	if _devMode {
		log.Info("starting app in dev mode")
	}

	chrome := forkedlorca.LocateChrome()
	if chrome == "" {
		forkedlorca.PromptDownload()
		log.Fatal("unable to locate Chrome")
		return
	}

	args := []string{}
	args = append(args, "--disable-features=TranslateUI,BlinkGenPropertyTrees")
	args = append(args, "--remote-allow-origins=*")
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	u, err := forkedlorca.New("", "", 900, 800, args...)
	if err != nil {
		log.Fatal(err)
	}
	ui := UI{u}
	defer ui.Close()

	updateKnownAccounts := func(accounts []Account) {
		encoded, err := json.Marshal(accounts)
		if err != nil {
			log.Error(err)
			return
		}
		ui.Eval(fmt.Sprintf("window.updateKnownAccounts(%s)", encoded))
	}
	updateState := func(state AppState) {
		ui.Eval(fmt.Sprintf("window.updateState('%s')", state))
	}
	updateMissionProgress := func(progress MissionProgress) {
		encoded, err := json.Marshal(progress)
		if err != nil {
			log.Error(err)
			return
		}
		ui.Eval(fmt.Sprintf("window.updateMissionProgress(%s)", encoded))
	}
	updateExportedFiles := func(files []string) {
		encoded, err := json.Marshal(files)
		if err != nil {
			log.Error(err)
			return
		}
		ui.Eval(fmt.Sprintf("window.updateExportedFiles(%s)", encoded))
	}
	emitMessage := func(message string, isError bool) {
		encoded, err := json.Marshal(message)
		if err != nil {
			log.Error(err)
			return
		}
		ui.Eval(fmt.Sprintf("window.emitMessage(%s, %t)", encoded, isError))
	}

	pinfo := func(args ...interface{}) {
		log.Info(args...)
		emitMessage(fmt.Sprint(args...), false)
	}
	perror := func(args ...interface{}) {
		log.Error(args...)
		emitMessage(fmt.Sprint(args...), true)
	}

	ui.MustBind("appVersion", func() string {
		return _appVersion
	})

	ui.MustBind("appDirectory", func() string {
		return _rootDir
	})

	ui.MustBind("appIsInForbiddenDirectory", func() bool {
		return _appIsInForbiddenDirectory
	})

	ui.MustBind("appIsTranslocated", func() bool {
		return _appIsTranslocated
	})

	ui.MustBind("knownAccounts", func() []Account {
		_storage.Lock()
		defer _storage.Unlock()
		return _storage.KnownAccounts
	})

	w := &worker{
		Weighted: semaphore.NewWeighted(1),
	}
	ui.MustBind("fetchPlayerData", func(playerId string) {
		go func() {
			if !w.TryAcquire(1) {
				perror("already fetching player data, cannot accept new work")
				return
			}
			defer w.Release(1)

			ctx, cancel := context.WithCancel(context.Background())
			w.ctxlock.Lock()
			w.ctx = ctx
			w.cancel = cancel
			w.ctxlock.Unlock()

			checkInterrupt := func() bool {
				select {
				case <-ctx.Done():
					perror("interrupted")
					updateState(AppState_INTERRUPTED)
					return true
				default:
					return false
				}
			}

			updateState(AppState_FETCHING_SAVE)
			fc, err := fetchFirstContactWithContext(w.ctx, playerId)
			if err != nil {
				perror(err)
				if !checkInterrupt() {
					updateState(AppState_FAILED)
				}
				return
			}
			nickname := fc.GetBackup().GetUserName()
			msg := fmt.Sprintf("successfully fetched backup for %s", playerId)
			if nickname != "" {
				msg += fmt.Sprintf(" (%s)", nickname)
			}
			pinfo(msg)
			lastBackupTime := fc.GetBackup().GetSettings().GetLastBackupTime()
			if lastBackupTime != 0 {
				t := unixToTime(lastBackupTime)
				now := time.Now()
				if t.After(now) {
					t = now
				}
				msg := fmt.Sprintf("backup is from %s", humanize.Time(t))
				pinfo(msg)
			} else {
				perror("backup is from unknown time")
			}
			_storage.AddKnownAccount(Account{Id: playerId, Nickname: nickname})
			_storage.Lock()
			updateKnownAccounts(_storage.KnownAccounts)
			_storage.Unlock()
			if checkInterrupt() {
				return
			}

			missions := fc.GetCompletedMissions()
			existingMissionIds, err := db.RetrievePlayerCompleteMissionIds(playerId)
			if err != nil {
				perror(err)
				updateState(AppState_FAILED)
				return
			}
			seen := make(map[string]struct{})
			for _, id := range existingMissionIds {
				seen[id] = struct{}{}
			}
			var newMissionIds []string
			var newMissionStartTimestamps []float64
			for _, mission := range missions {
				id := mission.GetIdentifier()
				if _, exists := seen[id]; !exists {
					newMissionIds = append(newMissionIds, id)
					newMissionStartTimestamps = append(newMissionStartTimestamps, mission.GetStartTimeDerived())
				}
			}
			pinfo(fmt.Sprintf("found %d completed missions, need to fetch %d",
				len(missions), len(newMissionIds)))

			total := len(newMissionIds)
			if total > 0 {
				updateState(AppState_FETCHING_MISSIONS)
				reportProgress := func(finished int) {
					updateMissionProgress(MissionProgress{
						Total:                   total,
						Finished:                finished,
						FinishedPercentage:      fmt.Sprintf("%.1f%%", float64(finished)/float64(total)*100),
						ExpectedFinishTimestamp: timeToUnix(time.Now().Add(time.Duration(total-finished) * _requestInterval)),
					})
				}
				reportProgress(0)
				finishedCh := make(chan struct{}, total)
				go func() {
					finished := 0
					for range finishedCh {
						finished++
						reportProgress(finished)
					}
				}()
				errored := 0
				var wg sync.WaitGroup
			MissionsLoop:
				for i := 0; i < total; i++ {
					if i != 0 {
						select {
						case <-ctx.Done():
							break MissionsLoop
						case <-time.After(_requestInterval):
						}
					}
					wg.Add(1)
					go func(missionId string, startTimestamp float64) {
						defer wg.Done()
						_, err := fetchCompleteMissionWithContext(w.ctx, playerId, missionId, startTimestamp)
						if err != nil {
							perror(err)
							errored++
						}
						finishedCh <- struct{}{}
					}(newMissionIds[i], newMissionStartTimestamps[i])
				}
				wg.Wait()
				close(finishedCh)
				if checkInterrupt() {
					return
				}
				if errored > 0 {
					perror(fmt.Sprintf("%d of %d missions failed to fetch", errored, total))
					updateState(AppState_FAILED)
					return
				} else {
					pinfo(fmt.Sprintf("successfully fetched %d missions", total))
				}
			}

			updateState(AppState_EXPORTING_DATA)
			completeMissions, err := db.RetrievePlayerCompleteMissions(playerId)
			if err != nil {
				perror(err)
				updateState(AppState_FAILED)
				return
			}
			var exportMissions []*mission
			for _, m := range completeMissions {
				exportMissions = append(exportMissions, newMission(m))
			}
			if checkInterrupt() {
				return
			}

			exportDir := filepath.Join(_rootDir, "exports", "missions")
			if err := os.MkdirAll(exportDir, 0755); err != nil {
				perror(errors.Wrap(err, "failed to create export directory"))
				updateState(AppState_FAILED)
				return
			}

			// Determine the last exported pair of xlsx and csv for future comparison.
			filenamePattern := regexp.QuoteMeta(playerId) + `\.\d{8}_\d{6}`
			lastExportedXlsxFile, err := findLastMatchingFile(exportDir, filenamePattern+`\.xlsx`)
			if err != nil {
				log.Errorf("error locating last exported .xlsx file: %s", err)
			}
			lastExportedCsvFile, err := findLastMatchingFile(exportDir, filenamePattern+`\.csv`)
			if err != nil {
				log.Errorf("error locating last exported .csv file: %s", err)
			}
			if filenameWithoutExt(lastExportedXlsxFile) != filenameWithoutExt(lastExportedCsvFile) {
				// If the xlsx and csv files aren't a pair, just leave them alone.
				lastExportedXlsxFile = ""
				lastExportedCsvFile = ""
			}

			filenameTimestamp := time.Now().Format("20060102_150405")

			xlsxFile := filepath.Join(exportDir, playerId+"."+filenameTimestamp+".xlsx")
			if err := exportMissionsToXlsx(exportMissions, xlsxFile); err != nil {
				perror(err)
				updateState(AppState_FAILED)
				return
			}
			if checkInterrupt() {
				return
			}

			csvFile := filepath.Join(exportDir, playerId+"."+filenameTimestamp+".csv")
			if err := exportMissionsToCsv(exportMissions, csvFile); err != nil {
				perror(err)
				updateState(AppState_FAILED)
				return
			}
			if checkInterrupt() {
				return
			}

			// Check if both exports are unchanged compared to the last exported pair.
			exportsUnchanged := lastExportedXlsxFile != "" && lastExportedCsvFile != "" && func() bool {
				xlsxUnchanged, err := cmpZipFiles(xlsxFile, lastExportedXlsxFile)
				if err != nil {
					log.Error(err)
					return false
				}
				if !xlsxUnchanged {
					return false
				}
				csvUnchanged, err := cmpFiles(csvFile, lastExportedCsvFile)
				if err != nil {
					log.Error(err)
					return false
				}
				if !csvUnchanged {
					return false
				}
				return true
			}()

			if exportsUnchanged {
				log.Info("exports unchanged, using last exported files and deleting new ones")
				emitMessage("exports identical with existing data files, reusing", false)
				err = os.Remove(xlsxFile)
				if err != nil {
					log.Errorf("error removing %s: %s", xlsxFile, err)
				}
				err = os.Remove(csvFile)
				if err != nil {
					log.Errorf("error removing %s: %s", csvFile, err)
				}

				xlsxFile = lastExportedXlsxFile
				csvFile = lastExportedCsvFile
			}
			xlsxFileRel, _ := filepath.Rel(_rootDir, xlsxFile)
			csvFileRel, _ := filepath.Rel(_rootDir, csvFile)
			updateExportedFiles([]string{xlsxFileRel, csvFileRel})

			pinfo("done.")
			updateState(AppState_SUCCESS)
		}()
	})

	ui.MustBind("stopFetchingPlayerData", func() {
		w.ctxlock.Lock()
		defer w.ctxlock.Unlock()
		if w.cancel != nil {
			w.cancel()
		}
	})

	ui.MustBind("doesExportExist", func() bool {
		for _, knownAccount := range _storage.KnownAccounts {
			ids, err := db.RetrievePlayerCompleteMissionIds(knownAccount.Id)
			if err != nil {
				log.Error(err)
			} else if len(ids) > 0 {
				return true
			}
		}
		return false
	})

	ui.MustBind("findExistingExports", func() []ExportAccount {
		knownAccounts := []ExportAccount{}
		for _, knownAccount := range _storage.KnownAccounts {
			ids, err := db.RetrievePlayerCompleteMissionIds(knownAccount.Id)
			if err != nil {
				log.Error(err)
			} else if len(ids) > 0 {
				knownAccounts = append(knownAccounts, ExportAccount{Id: knownAccount.Id, Nickname: knownAccount.Nickname, MissionCount: len(ids)})
			}
		}
		return knownAccounts
	})

	ui.MustBind("viewEidGo", func(eid string) string {
		if fileMissionYears, err := viewMissionsOfId(eid); err != nil {
			log.Error(err)
			return ""
		} else {
			return fileMissionYears
		}
	})

	ui.MustBind("getTargetName", func(target int) string {
		return ei.ArtifactSpec_Name_name[int32(target)]
	})

	ui.MustBind("getShipName", func(ship int) string {
		return ei.MissionInfo_Spaceship_name[int32(ship)]
	})

	ui.MustBind("getDurationName", func(duration int) string {
		return ei.MissionInfo_DurationType_name[int32(duration)]
	})

	/*
		Return a JSON array of the drops from a given mission
	*/
	ui.MustBind("getShipDrops", func(playerId string, shipId string) string {
		//Get the mission from the database
		completeMission, err := db.RetrieveCompleteMission(playerId, shipId)
		if err != nil {
			log.Error(err)
			return ""
		}

		shipDrops := []MissionDrop{} //Array of drops
		for _, drop := range completeMission.Artifacts {
			spec := drop.GetSpec()
			missionDrop := MissionDrop{
				Name:   ei.ArtifactSpec_Name_name[int32(spec.GetName())],
				Level:  int32(*drop.Spec.Level),
				Rarity: int32(*drop.Spec.Rarity),
			}
			switch {
			case strings.Contains(missionDrop.Name, "_FRAGMENT"):
				missionDrop.SpecType = "StoneFragment"
			case strings.Contains(missionDrop.Name, "_STONE"):
				missionDrop.SpecType = "Stone"
			case strings.Contains(missionDrop.Name, "GOLD_METEORITE"),
				strings.Contains(missionDrop.Name, "SOLAR_TITANIUM"),
				strings.Contains(missionDrop.Name, "TAU_CETI_GEODE"):
				missionDrop.SpecType = "Ingredient"
			default:
				missionDrop.SpecType = "Artifact"
			}
			shipDrops = append(shipDrops, missionDrop)
		}

		//Convert to JSON
		jsonData, err := json.Marshal(shipDrops)
		if err != nil {
			log.Error(err)
			return ""
		}

		//Return the JSON
		return string(jsonData)
	})

	ui.MustBind("getShipInfo", func(playerId string, shipId string) string {
		//Get the mission from the database
		completeMission, err := db.RetrieveCompleteMission(playerId, shipId)
		if err != nil {
			log.Error(err)
			return ""
		}

		launchDateTimeObject := time.Unix(int64(*completeMission.Info.StartTimeDerived), 0)
		ltH, ltM, ltS := launchDateTimeObject.Clock()
		launchTime := fmt.Sprintf(dateFormat, ltH, ltM, ltS)
		returnTimeObject := launchDateTimeObject.Add(time.Duration(*completeMission.Info.DurationSeconds * float64(time.Second)))
		rtH, rtM, rtS := returnTimeObject.Clock()
		returnTime := fmt.Sprintf(dateFormat, rtH, rtM, rtS)

		missionInst := LoadedMission{
			LaunchDay:   int32(launchDateTimeObject.Day()),
			LaunchMonth: int32(launchDateTimeObject.Month()),
			LaunchYear:  int32(launchDateTimeObject.Year()),
			LaunchTime:  launchTime,

			ReturnDay:   int32(returnTimeObject.Day()),
			ReturnMonth: int32(returnTimeObject.Month()),
			ReturnYear:  int32(returnTimeObject.Year()),
			ReturnTime:  returnTime,

			MissiondId:   *completeMission.Info.Identifier,
			Ship:         completeMission.Info.Ship,
			DurationType: completeMission.Info.DurationType,
			Level:        int32(*completeMission.Info.Level),
			Capacity:     int32(*completeMission.Info.Capacity),
			Target:       properTargetName(completeMission.Info.TargetArtifact),
		}

		// Convert the single mission to a JSON string
		jsonData, err := json.Marshal(missionInst)
		if err != nil {
			log.Error(err)
			return ""
		}

		// Return the JSON string
		return string(jsonData)
	})

	ui.MustBind("openFile", func(file string) {
		path := filepath.Join(_rootDir, file)
		if err := open.Start(path); err != nil {
			log.Errorf("opening %s: %s", path, err)
		}
	})

	ui.MustBind("openFileInFolder", func(file string) {
		path := filepath.Join(_rootDir, file)
		if err := openFolderAndSelect(path); err != nil {
			log.Errorf("opening %s in folder: %s", path, err)
		}
	})

	ui.MustBind("openURL", func(url string) {
		if err := open.Start(url); err != nil {
			log.Errorf("opening %s: %s", url, err)
		}
	})

	ui.MustBind("checkForUpdates", func() bool {
		log.Info("checking for updates...")
		newVersion, err := checkForUpdates()
		if err != nil {
			log.Error(err)
			return false
		}
		if newVersion == "" {
			log.Infof("no new version found")
			return false
		} else {
			log.Infof("new version found: %s", newVersion)
			return true
		}
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go func() {
		var httpfs http.FileSystem
		if _devMode {
			httpfs = http.Dir("www")
		} else {
			wwwfs, err := fs.Sub(_fs, "www")
			if err != nil {
				log.Fatal(err)
			}
			httpfs = http.FS(wwwfs)
		}
		err := http.Serve(ln, http.FileServer(httpfs))
		if err != nil {
			log.Fatal(err)
		}
	}()
	ui.MustLoad(fmt.Sprintf("http://%s/", ln.Addr()))

	// Wait until the interrupt signal arrives or browser window is closed.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
}

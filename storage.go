package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type AppStorage struct {
	sync.Mutex

	KnownAccounts []Account `json:"known_accounts"`

	LastMennoDataRefreshAt  time.Time `json:"last_menno_data_refresh_at"`
	LastUpdateCheckAt       time.Time `json:"last_update_check_at"`
	KnownLatestReleaseNotes string    `json:"known_latest_release_notes"`
	KnownLatestVersion      string    `json:"known_latest_version"`
	FilterWarningRead       bool      `json:"filter_warning_read"`
	PreferredChromiumPath   string    `json:"preferred_chromium_path"`
	AutoRefreshMennoPref    bool      `json:"auto_refresh_menno_pref"`
	UseGifsForRarity        bool      `json:"use_gifs_for_rarity"`
	DefaultViewMode         string    `json:"default_view_mode"`
}

type Account struct {
	Id            string  `json:"id"`
	Nickname      string  `json:"nickname"`
	EarningsBonus float64 `json:"earningsBonus"`
	EBString      string  `json:"ebString"`
	AccountColor  string  `json:"accountColor"`
}

var (
	_storageFile string
	_storage     AppStorage
)

func storageInit() {
	_storageFile = filepath.Join(_internalDir, "storage.json")
	_storage.Load()
}

func (s *AppStorage) Load() {
	s.Lock()
	defer s.Unlock()
	encoded, err := os.ReadFile(_storageFile)
	if err != nil {
		log.Errorf("error loading storage.json: %s", err)
		return
	}
	if err := json.Unmarshal(encoded, &s); err != nil {
		log.Errorf("error parsing storage.json: %s", err)
		return
	}
}

func (s *AppStorage) Persist() {
	s.Lock()
	defer s.Unlock()
	encoded, err := json.Marshal(s)
	if err != nil {
		log.Errorf("error serializing app storage: %s", err)
		return
	}
	if err := os.WriteFile(_storageFile, encoded, 0644); err != nil {
		log.Errorf("error writing app storage: %s", err)
	}
}

func (s *AppStorage) AddKnownAccount(account Account) {
	s.Lock()
	accounts := []Account{account}
	seen := map[string]struct{}{account.Id: {}}
	for _, a := range s.KnownAccounts {
		if _, exists := seen[a.Id]; !exists {
			accounts = append(accounts, a)
			seen[a.Id] = struct{}{}
		}
	}
	s.KnownAccounts = accounts
	s.Unlock()
	go s.Persist()
}

func (s *AppStorage) SetUpdateCheck(latestVersion string, latestReleaseNotes string) {
	s.Lock()
	s.LastUpdateCheckAt = time.Now()
	s.KnownLatestVersion = latestVersion
	s.KnownLatestReleaseNotes = latestReleaseNotes
	s.Unlock()
	go s.Persist()
}

func (s *AppStorage) SetFilterWarningRead(flag bool) {
	s.Lock()
	s.FilterWarningRead = flag
	s.Unlock()
	go s.Persist()
}

func (s *AppStorage) SetLastMennoDataRefreshAt(t time.Time) {
	s.Lock()
	s.LastMennoDataRefreshAt = t
	s.Unlock()
	go s.Persist()
}

func (s *AppStorage) SetPreferredChromiumPath(path string) {
	s.Lock()
	s.PreferredChromiumPath = path
	s.Unlock()
	go s.Persist()
}

func (s *AppStorage) SetAutoRefreshMennoPref(flag bool) {
	s.Lock()
	s.AutoRefreshMennoPref = flag
	s.Unlock()
	go s.Persist()
}

func (s *AppStorage) SetUseGifsForRarity(flag bool) {
	s.Lock()
	s.UseGifsForRarity = flag
	s.Unlock()
	go s.Persist()
}

func (s *AppStorage) SetDefaultViewMode(mode string) {
	s.Lock()
	s.DefaultViewMode = mode
	s.Unlock()
	go s.Persist()
}

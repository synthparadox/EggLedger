package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type MennoData struct {
	ConfigurationItems []ConfigurationItem `json:"configurationItems"`
}

type ConfigurationItem struct {
	ShipConfiguration     ShipConfiguration     `json:"shipConfiguration"`
	ArtifactConfiguration ArtifactConfiguration `json:"artifactConfiguration"`
	TotalDrops            int                   `json:"totalDrops"`
}

type ShipConfiguration struct {
	ShipType         IdNamePair `json:"shipType"`
	ShipDurationType IdNamePair `json:"shipDurationType"`
	Level            int        `json:"level"`
	TargetArtifact   IdNamePair `json:"targetArtifact"`
}

type ArtifactConfiguration struct {
	ArtifactType   IdNamePair `json:"artifactType"`
	ArtifactRarity IdNamePair `json:"artifactRarity"`
	ArtifactLevel  int        `json:"artifactLevel"`
}

type IdNamePair struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const _mennoDateFormat = "2006-01-02-15-04-05"
const _mennoFileFormat = "menno-data-%s.json"

func loadLatestMennoData() (data MennoData, err error) {
	_storage.Lock()
	latestRefresh := _storage.LastMennoDataRefreshAt
	_storage.Unlock()

	// Get the file name for the latest data.
	filename := fmt.Sprintf(_mennoFileFormat, latestRefresh.Format(_mennoDateFormat))
	filePath := filepath.Join(_internalDir, filename)

	returnData := []ConfigurationItem{}

	// Read the file from disk.
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return MennoData{
			ConfigurationItems: returnData,
		}, err
	}

	// Unmarshal the JSON data.
	err = json.Unmarshal(file, &returnData)
	if err != nil {
		fmt.Println(err)
		return MennoData{
			ConfigurationItems: returnData,
		}, err
	}

	return MennoData{
		ConfigurationItems: returnData,
	}, nil
}

func checkIfRefreshMennoDataIsNeeded() bool {
	if !_storage.AutoRefreshMennoPref {
		return false
	}

	// Update every 5 days.
	_storage.Lock()
	lastMennoRefesh := _storage.LastMennoDataRefreshAt
	_storage.Unlock()

	return (time.Since(lastMennoRefesh) > time.Hour*24*5)
}

func refreshMennoData() (err error) {

	// Fetch the data from the Menno server.
	resp, err := http.Get("https://eggincdatacollection.azurewebsites.net/api/GetAllData")
	if err != nil {
		// Log the error, but don't return it.
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// Read the data from the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Log the error, but don't return it.
		fmt.Println(err)
		return err
	}

	// Unmarshal
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		// Log the error, but don't return it.
		fmt.Println(err)
		return err
	}

	_storage.Lock()
	oldDataTime := _storage.LastMennoDataRefreshAt
	_storage.Unlock()

	// Save the JSON to a file with a date-time stamp.
	newTime := time.Now()
	filename := fmt.Sprintf(_mennoFileFormat, newTime.Format(_mennoDateFormat))
	filePath := filepath.Join(_internalDir, filename)

	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		// Log the error, but don't return it.
		fmt.Println(err)
		return err
	}

	// Update the last refresh time.
	_storage.SetLastMennoDataRefreshAt(newTime)

	//Remove old file
	oldFileName := fmt.Sprintf(_mennoFileFormat, oldDataTime.Format(_mennoDateFormat))
	oldFilePath := filepath.Join(_internalDir, oldFileName)
	err = os.Remove(oldFilePath)
	if err != nil {
		// Log the error, but don't return it.
		fmt.Println(err)
	}

	// Return nil if everything went well.
	return nil
}

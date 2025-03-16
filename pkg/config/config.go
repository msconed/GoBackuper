package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type MailRuConfig struct {
	// https://account.mail.ru/user/2-step-auth/passwords
	MAILRU_WEBDAV3_HOSTNAME string   `json:"MAILRU_WEBDAV3_HOSTNAME"` // Password
	MAILRU_WEBDAV3_USERNAME string   `json:"MAILRU_WEBDAV3_USERNAME"` // Username
	MAILRU_WEBDAV3_PASSWORD string   `json:"MAILRU_WEBDAV3_PASSWORD"` // Password
	MailRuDirName           string   `json:"MailRuDirName"`           // Save to Directory
	DirsToBackup            []string `json:"DirsToBackup"`            // Directories to backup
	FilesToBackup           []string `json:"FilesToBackup"`           // Files to backup
}

var config_mailru *MailRuConfig
var countTotalFiles int

func InitMailRuConfig() {
	if config_mailru == nil {
		ParseJSON_MailRu()
	}
}

func ParseJSON_MailRu() error {
	jsonFile, err := os.Open("config_mailru.json")

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	byteValue, _ := io.ReadAll(jsonFile)

	err2 := json.Unmarshal(byteValue, &config_mailru)

	if err2 != nil {
		fmt.Println("Cannot parse json ERROR: ", err2)
		return err2
	}

	defer jsonFile.Close()
	return nil
}

func GetConfig() *MailRuConfig {
	return config_mailru
}

func GetCountTotalFiles() int {
	return countTotalFiles
}

func AddCountTotalFiles(count int) {
	countTotalFiles += count
}

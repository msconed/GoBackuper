package MailRuBackuper

import (
	"GoBackuper/pkg/config"
	"GoBackuper/pkg/file_c"
	"GoBackuper/pkg/gowebdav"
	"GoBackuper/pkg/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

// https://github.com/studio-b12/gowebdav?tab=readme-ov-file#pkg-examples

var client *gowebdav.Client

func UploadFile(directory string, file string) {
	if !(DirectoryExists(directory)) {
		client.Mkdir(directory, 0644)
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := fileInfo.Size()

	bar := progressbar.DefaultBytes(
		fileSize,
		fmt.Sprintf("Uploading %s", filepath.Base(file)),
	)

	reader := io.TeeReader(f, bar)

	uploadPath := directory + "/" + filepath.Base(file)

	client.WriteStream(uploadPath, reader, 0644)

	bar.Finish()

	fmt.Println("File successfully uploaded:", uploadPath)

	defer file_c.RemoveFileIfExist(file)
	defer f.Close()
}

func GetClient() *gowebdav.Client {
	return client
}

func InitMailRuClient() {
	err := config.ParseJSON_MailRu()

	if err != nil {
		fmt.Println("Error: ", err)
	}
	client = gowebdav.NewClient(config.GetConfig().MAILRU_WEBDAV3_HOSTNAME, config.GetConfig().MAILRU_WEBDAV3_USERNAME, config.GetConfig().MAILRU_WEBDAV3_PASSWORD)
	client.Connect()
}

func DirectoryExists(dirName string) bool {
	_, err := client.Stat(dirName)
	return err == nil
}

func RunBackup() {
	config.InitMailRuConfig()

	dirs := config.GetConfig().DirsToBackup

	for _, dir := range dirs {
		count, _ := file_c.CountFilesInDirectory(dir)
		config.AddCountTotalFiles(count)
	}

	InitMailRuClient()
	zipfile := zip.ZipWriter(dirs)

	UploadFile(config.GetConfig().MailRuDirName, zipfile)
}

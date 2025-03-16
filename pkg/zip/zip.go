package zip

import (
	"GoBackuper/pkg/config"
	"GoBackuper/pkg/pb"
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func ZipWriter(dirs []string) string {

	pb.InitProgressBar(int64(config.GetCountTotalFiles()), "Packing files...")

	path, _ := os.Getwd()

	packingTime := time.Now().Format("2006-01-02_15-04-05")

	// Get a Buffer to Write To
	outFile, err := os.Create(path + "/" + packingTime + "_backup.zip")
	if err != nil {
		fmt.Println(err)
	}

	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	for _, path := range dirs {
		path = path + "/"
		dirName := filepath.Base(path)

		// Add some files to the archive.
		AddFiles(w, path, dirName+"/")
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}

	return outFile.Name()
}

func AddFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		//	fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := os.ReadFile(basePath + file.Name())
			if err != nil {
				//	fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}

			pb.GetProgressBar().Add(1)
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			//	fmt.Println("Recursing and Adding SubDir: " + file.Name())
			//	fmt.Println("Recursing and Adding SubDir: " + newBase)

			AddFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}

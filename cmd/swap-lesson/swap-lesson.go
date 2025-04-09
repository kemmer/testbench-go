package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const DownloadDir = "download"
const GameDir = "gamedir"

func main() {
	files := []string{"game.zip", "gameset.zip", "sounds.zip"}
	for _, a := range files {
		unzip(a)
		fmt.Println()
	}
	for _, a := range files {
		listzip(a)
		fmt.Println()
	}
}

func listzip(filename string) {
	archive, err := zip.OpenReader(filepath.Join(DownloadDir, filename))
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		fmt.Println(f.Name)
	}
}

func unzip(filename string) {

	filenameParts := strings.Split(filename, ".")
	if len(filenameParts) < 1 {
		return
	}

	dirname := filenameParts[0]
	dirpath := filepath.Join(GameDir, dirname)
	err := os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("unzip file: %s, destination folder: %s\n", filename, dirpath)

	archive, err := zip.OpenReader(filepath.Join(DownloadDir, filename))
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		dstPath := filepath.Join(GameDir, dirname, f.Name)

		fmt.Printf("unzipping file %s -> %s\n", strings.TrimRight(f.Name, "/"), dstPath)

		// check for invalid chars in file name
		if !strings.HasPrefix(dstPath, filepath.Clean(GameDir)+string(os.PathSeparator)) {
			fmt.Printf("invalid file path: '%s'\n", dstPath)
			return
		}

		err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		if err != nil {
			panic(err)
		}

		if f.FileInfo().IsDir() {
			continue
		}

		dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		srcFile, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			panic(err)
		}

		dstFile.Close()
		srcFile.Close()
	}
}

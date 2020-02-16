package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	zipName = os.Args[1]
	fileName = os.Args[2]

	f, err := os.Create(zipName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cpFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer cpFile.Close()

	zipWriter = zip.NewWriter(f)
	writer, err := zipWriter.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer zipWriter.Close()
	io.Copy(writer, cpFile)
}

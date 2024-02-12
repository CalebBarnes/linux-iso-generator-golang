package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ProgressWriter struct {
	Total    int64
	Written  int64
	Progress int
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Written += int64(n)
	newProgress := int(100 * pw.Written / pw.Total)
	// Update progress if it has increased.
	if newProgress > pw.Progress {
		pw.Progress = newProgress
		fmt.Printf("\rDownloading... %d%% complete", pw.Progress)
	}
	return n, nil
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if we can determine the progress
	total := resp.ContentLength
	if total <= 0 {
		fmt.Println("Cannot determine file size, progress will not be shown.")
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create a ProgressWriter
	pw := &ProgressWriter{Total: total}

	// Write the body to file, through our ProgressWriter
	_, err = io.Copy(out, io.TeeReader(resp.Body, pw))
	if err != nil {
		return err
	}
	println("\nDownload completed.\n")
	return err
}

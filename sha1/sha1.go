package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// $ cat http.log.gz
// When opening a file, have to make sure the file is always closed at the end.
// Defer only happens when the function exists
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}

	defer file.Close() // Always in the function level, shouldn't be in a for loop
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gz.Close()

		r = gz
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(sig)
}

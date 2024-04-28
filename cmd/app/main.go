package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

type pair struct {
	hash, path string
}

type fileList []string

// hash is a map key
type results map[string]fileList

// Hashing file
func hashFile(path string) pair {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

// Walking through the directory recursively
func searchTree(dir string) (results, error) {
	hashes := make(results)

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.Mode().IsRegular() && info.Size() > 0 {
			p := hashFile(path)
			hashes[p.hash] = append(hashes[p.hash], p.path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return hashes, nil
}

func sequentialFileProcessing(filePath string) {
	start := time.Now()

	if hashes, err := searchTree(filePath); err == nil {
		for hash, files := range hashes {
			if len(files) > 1 {
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files {
					fmt.Println("\t", file)
				}
			}
		}
	}

	fmt.Println("took", time.Since(start).Seconds(), "seconds")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide dir name!")
	}

	sequentialFileProcessing(os.Args[1])
}

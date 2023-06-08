package ooo

import (
	"fmt"
	"os"
)

func CountFiles(stDirectory string, ext string) int {
	if !Exists(stDirectory) {
		return 0
	}
	u := 0
	f, err := os.Open(stDirectory)
	if err != nil {
		return u
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return u
	}
	for _, v := range files {
		if !v.IsDir() {
			nam := v.Name()
			if ext == "" {
				u++
			} else {
				if Rep(nam, ext, "") != nam {
					u++
				}
			}
		}
	}
	return u
}

func CountFolders(stDirectory string, ext string) int {
	if !Exists(stDirectory) {
		return 0
	}
	u := 0
	f, err := os.Open(stDirectory)
	if err != nil {
		return u
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return u
	}
	for _, v := range files {
		if v.IsDir() {
			nam := v.Name()
			if ext == "" {
				u++
			} else {
				if Rep(nam, ext, "") != nam {
					u++
				}
			}
		}
	}
	return u
}

func Exists(stpath string) bool {
	if _, err := os.Stat(stpath); err == nil {
		return true
	} else {
		return false
	}
}

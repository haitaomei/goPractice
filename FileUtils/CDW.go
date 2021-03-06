package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
)

//CreateFile create a new file, if directory doesn't exist, create one
func CreateFile(path string) {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)

	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if checkErr(err) {
			return
		}
		defer file.Close()
	}
}

//WriteFile writes a file
func WriteFile(path string, data *[]byte) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if checkErr(err) {
		return
	}
	defer file.Close()

	_, err = file.Write(*data)
	if checkErr(err) {
		return
	}
	err = file.Sync()
	if checkErr(err) {
		return
	}
}

//DeleteFile deletes a file
func DeleteFile(path string) {
	var err = os.Remove(path)
	if checkErr(err) {
		return
	}
}

//DeleteDir delete a dir
func DeleteDir(path string) {
	fi, err := os.Stat(path)
	if checkErr(err) {
		return
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		err = os.RemoveAll(path)
		if checkErr(err) {
			return
		}
	case mode.IsRegular():
	}
}

func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

//RemoveContents delete all files and sub direcotries within a dir
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

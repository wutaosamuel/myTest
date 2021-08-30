package utils

import (
	"os"
	"io"
	"io/ioutil"
	"path"
)

// IsFile check file exist
// true, nil -> is a file
// false, nil -> not exist a file
// false, err -> an err output
func IsFile(name string) (bool, error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// IsDir check is Dir
func IsDir(dir string) (bool,error) {
	d, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	return d.IsDir(), nil
}

// CopyFile copy file
func CopyFile(name, src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, srcFile); err != nil {
		return err
	}
	return nil
}

// CopyDirTree copy dir tree
// NOTE: changed
func CopyDirTree(dir, src string) error {
	// check whether dir exist or not
	isFile, err := IsFile(dir)
	if err != nil {
		return err
	}
	if isFile {
		return Err("Directory exist.")
	}
	// get properties of src
	srcDir, err := os.Stat(src)
	if err != nil {
		return err
	}

	// copy target dir
	if err := os.MkdirAll(dir, srcDir.Mode()); err != nil {
		return err
	}

	// copy tree into dir
	// sub-dirs 
	// files
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries{
		dirChild := path.Join(dir, entry.Name())
		srcChild := path.Join(src, entry.Name())
		if entry.IsDir() {
			if err := CopyDirTree(dirChild, srcChild); err != nil {
				return err
			}
		}
		if !entry.IsDir() {
			if err := CopyFile(dirChild, srcChild); err != nil {
				return err
			}
		}
	}
	return nil
}

// os.Mkdir("./dir", os.ModePerm)

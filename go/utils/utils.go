package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// IsExist check if file or dir exist
/**
 * IsExist
 * @Params
 *  - path 		 		 <- absolute file path
 *
 * @Returns
 *	- true, nil 	 -> file exist
 *	- false, nil	 -> file not exist
 *	- _, error		 -> error
 **/
func IsExist(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// IsFile check if it is a file, not a dir
/**
 * IsFile
 * @Params
 *	- name					<- file name with absolute file path
 *
 * @Returns
 *  - true, nil		 	-> is file
 *  - false, nil		-> is not file or exist
 *  - _, err 				-> error
 */
func IsFile(name string) (bool, error) {
	f, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	return !f.IsDir(), nil
}

// IsDir check if is a Dir, not a file
/**
 * IsDir
 * @Params
 * 	- dir						<- directory path with absolute path
 *
 * @Returns
 *  - true, nil 		-> is dir
 *  - false, nil		-> is not dir or exist
 *  - _, err 				-> error
 **/
func IsDir(dir string) (bool, error) {
	d, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	return d.IsDir(), nil
}

/**
 *CopyDirTree copy dir with all contents to destination
 * @Params
 *	- dst   destination
 *  - src		source dir
 **/
func CopyDirTree(dst, src string) error {
	// check whether dir exist or not
	if isDir, err := IsDir(dst); err != nil {
		return err
	} else {
		if isDir {
			return errors.New("directory exist")
		}
	}
	// get properties of src
	srcDir, err := os.Stat(src)
	if err != nil {
		return err
	}

	// copy target dir
	if err := os.MkdirAll(dst, srcDir.Mode()); err != nil {
		return err
	}

	// copy tree into dir
	// sub-dirs
	// files
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		dstChild := path.Join(dst, entry.Name())
		srcChild := path.Join(src, entry.Name())
		if entry.IsDir() {
			if err := CopyDirTree(dstChild, srcChild); err != nil {
				return err
			}
		}
		if !entry.IsDir() {
			if err := CopyFile(dstChild, srcChild); err != nil {
				return err
			}
		}
	}
	return nil
}

/**
 *	CopyFile copy file
 *	@Params
 *		- dst   destination
 *		- src   source file
 **/
func CopyFile(dst, src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, srcFile); err != nil {
		return err
	}
	return nil
}

package main

import (
	"os"
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

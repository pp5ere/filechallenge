package util

import "os"

//DirExist return if directory exists or not
func DirExist(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	if !info.IsDir() {
        return false
	}
	return true
}

//CreateDir create a new directory folder
func CreateDir(dirPath string) error{
	if !DirExist(dirPath){
		if err := os.Mkdir(dirPath, 0755); err != nil{
			return err
		}
	}
	return nil
}
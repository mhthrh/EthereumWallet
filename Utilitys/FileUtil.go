package Utilitys

import "os"

func LockFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); err == nil {
		err = os.Remove(filename)
		if err != nil {
			return nil, err
		}

	}
	return os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666)
}

func GetPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, err
}

func ExistDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return true, nil
	}
	if err := os.Mkdir(path, 7777); err == nil {
		return true, nil
	}
	return false, err

}

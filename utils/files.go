package utils

import "os"

func GetTempFileName() string {
	file, err := os.CreateTemp(os.TempDir(), "*.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return file.Name()
}

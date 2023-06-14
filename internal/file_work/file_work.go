package file_work

import (
	"log"
	"os"
)

func OpenRead(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func OpenWrite(name string) *os.File {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func CloseFunc(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
}

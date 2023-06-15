package fileWork

import (
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

const (
	nameWin = "123"
	nameLin = "000"
	nameMac = "321"
)

func Update() {
	var useName string
	switch runtime.GOOS {
	case "darwin":
		useName = nameMac
	case "linux":
		useName = nameLin
	case "windows":
		useName = nameWin
	}

	url := "https://github.com/lipipidronstudy/updateRepo/blob/main/"
	resp, err := http.Get(url + useName)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	out, err := os.Create(useName)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseFunc(out)

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}

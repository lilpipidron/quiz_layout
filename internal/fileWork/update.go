package fileWork

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	nameWin = "123"
	nameLin = "README.md"
	nameMac = "321"
)

func Update(system int) {
	var useName string
	if system == 0 {
		useName = nameWin
	} else if system == 1 {
		useName = nameLin
	} else {
		useName = nameMac
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

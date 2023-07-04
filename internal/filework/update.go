package filework

import (
	"game_with_Nikita/consts"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

func Update() {
	var useName string
	switch runtime.GOOS {
	case "darwin":
		useName = consts.NameMac
	case "linux":
		useName = consts.NameLin
	case "windows":
		useName = consts.NameWin
	}

	url := "url репо, откуда качать"
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

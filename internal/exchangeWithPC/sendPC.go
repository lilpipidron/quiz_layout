package exchangeWithPC

import (
	"fmt"
	"game_with_Nikita/internal/fileWork"
	"io"
	"log"
	"net/http"
)

func SendToIp(name string) {
	file := fileWork.OpenRead(name)
	defer fileWork.CloseFunc(file)

	url := "http://192.168.0.192:1716/upload" //пока не работает
	resp, err := http.Post(url, "multipart/form-data", file)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	fmt.Println("Деньги украдены")
}

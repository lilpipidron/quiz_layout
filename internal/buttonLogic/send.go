package buttonLogic

import (
	"GolandProjects/game_with_Nikita/internal/fileWork"
	"GolandProjects/game_with_Nikita/internal/myHash"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func send_login_sign(login string) {

	logFile := fileWork.OpenRead("log.txt")
	defer fileWork.CloseFunc(logFile)
	reader := bufio.NewReader(logFile)
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	if string(line) != login {
		log.Fatal("Incorrect login/password")
	}
}
func send_password_sign(password string) {
	hash_pass := myHash.Password_hash(password)
	logFile := fileWork.OpenRead("log.txt")
	defer fileWork.CloseFunc(logFile)
	reader := bufio.NewReader(logFile)

	_, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	lineS, err := strconv.ParseUint(string(line), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	if lineS != hash_pass {
		log.Fatal("Incorrect login/password")
	}
}
func send_login_registation(login string) {

	logFile := fileWork.OpenWrite("log.txt")
	defer fileWork.CloseFunc(logFile)
	writer := bufio.NewWriter(logFile)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(writer)
	_, err := fmt.Fprintln(writer, login)
	if err != nil {
		log.Fatal(err)
	}
}
func send_password_registration(password string) {
	hash_pass := myHash.Password_hash(password)
	logFile := fileWork.OpenWrite("log.txt")
	defer fileWork.CloseFunc(logFile)

	scanner := bufio.NewScanner(logFile)
	writer := bufio.NewWriter(logFile)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(writer)

	scanner.Scan()
	_, err := fmt.Fprintln(writer, hash_pass)
	if err != nil {
		log.Fatal(err)
	}

}

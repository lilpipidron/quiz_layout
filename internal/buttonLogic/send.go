package buttonLogic

import (
	"bufio"
	"fmt"
	"game_with_Nikita/internal/fileWork"
	"game_with_Nikita/internal/myHash"
	"log"
	"strconv"
)

func sendLoginLogIn(login string) bool {
	filename := login + "-log.txt"
	logFile := fileWork.OpenRead(filename)
	defer fileWork.CloseFunc(logFile)
	reader := bufio.NewReader(logFile)
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	if string(line) != login {
		return false
	}
	return true
}

func sendPasswordLogIn(login, password string) bool {
	hashPass := myHash.Password_hash(password)
	filename := login + "-log.txt"
	logFile := fileWork.OpenRead(filename)
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

	if lineS != hashPass {
		return false
	}
	return true
}

func sendLoginSignUp(login string) {
	filename := login + "-log.txt"
	logFile := fileWork.OpenWrite(filename)
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

func sendPasswordSignUp(login, password string) {
	hashPass := myHash.Password_hash(password)
	filename := login + "-log.txt"
	logFile := fileWork.OpenWrite(filename)
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
	_, err := fmt.Fprintln(writer, hashPass)
	if err != nil {
		log.Fatal(err)
	}

}

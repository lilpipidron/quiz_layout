package logs

import (
	"bufio"
	"fmt"
	"game_with_Nikita/internal/filework"
	"game_with_Nikita/internal/hash"
	"log"
	"strconv"
)

func SendLoginLogIn(login string) bool {
	filename := login + "-log.txt"
	logFile := filework.OpenRead(filename)
	defer filework.CloseFunc(logFile)
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

func SendPasswordLogIn(filename, password string) bool {
	if len(password) < 8 {
		return false
	}
	hashPass := hash.PasswordHash(password)
	logFile := filework.OpenRead(filename)
	defer filework.CloseFunc(logFile)
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

func SendLoginSignUp(login string) {
	filename := login + "-log.txt"
	logFile := filework.OpenWrite(filename)
	defer filework.CloseFunc(logFile)
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

func SendPasswordSignUp(login, password string) {
	hashPass := hash.PasswordHash(password)
	filename := login + "-log.txt"
	logFile := filework.OpenWrite(filename)
	defer filework.CloseFunc(logFile)

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

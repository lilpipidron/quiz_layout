package ui

import (
	"bufio"
	"fmt"
	"game_with_Nikita/internal/file_work"
	"game_with_Nikita/internal/hash"
	"log"
	"strconv"
)

func sendLoginLogIn(login string) bool {
	filename := login + "-log.txt"
	logFile := file_work.OpenRead(filename)
	defer file_work.CloseFunc(logFile)
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

func sendPasswordLogIn(filename, password string) bool {
	if len(password) < 8 {
		return false
	}
	hashPass := hash.PasswordHash(password)
	logFile := file_work.OpenRead(filename)
	defer file_work.CloseFunc(logFile)
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
	logFile := file_work.OpenWrite(filename)
	defer file_work.CloseFunc(logFile)
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
	hashPass := hash.PasswordHash(password)
	filename := login + "-log.txt"
	logFile := file_work.OpenWrite(filename)
	defer file_work.CloseFunc(logFile)

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
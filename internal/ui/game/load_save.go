package game

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetSaveData(username string) int {
	maxStage := 1
	filename := username + "-log.txt"

	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		const stringsToKeep = 3
		if len(lines) >= stringsToKeep {
			if lines[len(lines)-1] == "" {
				lines = lines[:len(lines)-2] //fixme
				lines = append(lines, strconv.Itoa(maxStage))
				if err := file.Close(); err != nil {
					log.Fatal(err)
				}
				if err := os.Remove(filename); err != nil {
					log.Fatal(err)
				}
				file, err = os.Create(filename)
				if err != nil {
					log.Fatal(err)
				}
				writer := bufio.NewWriter(file)
				for _, line := range lines {
					if _, err := writer.WriteString(line + "\n"); err != nil {
						log.Fatal(err)
					}
				}
				if err := writer.Flush(); err != nil {
					log.Fatal(err)
				}
			} else {
				maxStage, _ = strconv.Atoi(lines[len(lines)-1])
			}
		}
	} else {
		log.Fatal(err)
	}
	return maxStage
}

func SetSaveData(username string) {
	filename := username + "-log.txt"
	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		const stringsToKeep = 3
		if len(lines) >= stringsToKeep {
			lines = lines[:len(lines)-2] //fixme
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
			if err := os.Remove(filename); err != nil {
				log.Fatal(err)
			}
			file, err = os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			writer := bufio.NewWriter(file)
			for _, line := range lines {
				if _, err := writer.WriteString(line + "\n"); err != nil {
					log.Fatal(err)
				}
			}
			if err := writer.Flush(); err != nil {
				log.Fatal(err)
			}

		}
	} else {
		log.Fatal(err)
	}
}

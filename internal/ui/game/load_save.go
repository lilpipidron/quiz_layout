package game

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func getSaveData(username string, currentMaxStage int) int {
	maxStage := currentMaxStage
	if currentMaxStage == 0 {
		maxStage = 1
	}
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
		const lenIfHasMaxStageData = 3
		data, err := strconv.Atoi(lines[len(lines)-1])
		if err != nil {
			log.Fatal(err)
		}
		if (len(lines) != lenIfHasMaxStageData) || (len(lines) == lenIfHasMaxStageData && data < currentMaxStage) {
			if len(lines) != lenIfHasMaxStageData {
				lines = append(lines, strconv.Itoa(maxStage))
			} else {
				lines[len(lines)-1] = strconv.Itoa(currentMaxStage)
			}
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
		} else if len(lines) == lenIfHasMaxStageData && data > currentMaxStage {
			maxStage = data
		}
	} else {
		log.Fatal(err)
	}
	return maxStage
}

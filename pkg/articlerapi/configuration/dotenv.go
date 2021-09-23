package configuration

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func ProcessDotEnv(dirPaths ...string) error {

	for _, path := range dirPaths {
		if err := parseDotenvDir(path); err != nil {
			return err
		}
	}

	return nil
}

func parseDotenvDir(dirPath string) error {
	files, err := os.ReadDir(dirPath)

	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".env") {
			if err := processDotenvFile(dirPath, file); err != nil {
				return err
			}
		}
	}

	return nil
}

func processDotenvFile(dirPath string, file os.DirEntry) error {
	fileReader, err := os.Open(strings.TrimSuffix(dirPath, "/") + "/" + file.Name())

	defer func() {
		if err := fileReader.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(fileReader)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		if utf8.RuneCount([]byte(trimmedLine)) >= 3 {
			splittedLine := strings.Split(trimmedLine, "=")

			if len(splittedLine) == 2 {
				err := os.Setenv(strings.TrimSpace(splittedLine[0]), strings.TrimSpace(splittedLine[1]))
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

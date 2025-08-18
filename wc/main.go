package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	countBytes, countLines, countWords := getFlags()

	file, fileInfo, err := getFileAndInfo()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	bytes := fileInfo.Size()
	output := file.Name()

	if countBytes {
		output = fmt.Sprintf("%d %s", bytes, output)
	}

	if countWords {
		nrOfWords := countWordsInFile(file)
		output = fmt.Sprintf("%d %s", nrOfWords, output)
	}

	if countLines {
		nrOfLines := countLinesInFile(file)
		output = fmt.Sprintf("%d %s", nrOfLines, output)
	}

	fmt.Println(output)
}

func getFlags() (bool, bool, bool) {
	var countBytes bool
	var countLines bool
	var countWords bool

	flag.BoolVar(&countBytes, "c", false, "count number of bytes in file")
	flag.BoolVar(&countLines, "l", false, "count number of lines in file")
	flag.BoolVar(&countWords, "w", false, "count number of words in file")

	flag.Parse()

	noFlagsSet := !countBytes && !countLines && !countWords
	if noFlagsSet {
		countBytes = true
		countLines = true
		countWords = true
	}

	return countBytes, countLines, countWords
}

func getFileAndInfo() (*os.File, os.FileInfo, error) {
	filename := flag.Arg(0)
	if filename == "" {
		return nil, nil, fmt.Errorf("Missing filename.\n")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"Unable to open file '%s' for reading. Does it exist?",
			filename,
		)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf(
			"Unable to query fileinfo for file '%s'.",
			filename,
		)
	}

	return file, fileInfo, nil
}

func countLinesInFile(file *os.File) int {
	nrOfLines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nrOfLines++
	}

	return nrOfLines
}

func countWordsInFile(file *os.File) int {
	nrOfWords := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tokens := []rune(line)
		inWord := false
		for _, token := range tokens {
			if (unicode.IsLetter(token)) && !inWord {
				inWord = true
				continue
			}

			if !(unicode.IsLetter(token)) && inWord {
				nrOfWords++
				inWord = false
			}
		}
	}

	return nrOfWords
}

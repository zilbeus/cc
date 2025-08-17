package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	countBytes, countLines := getFlags()

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

	if countLines {
		nrOfLines := countLinesInFile(file)
		output = fmt.Sprintf("%d %s", nrOfLines, output)
	}

	fmt.Println(output)
}

func getFlags() (bool, bool) {
	var countBytes bool
	var countLines bool

	flag.BoolVar(&countBytes, "c", false, "count number of bytes in file")
	flag.BoolVar(&countLines, "l", false, "count number of lines in file")

	flag.Parse()

	noFlagsSet := !countBytes && !countLines
	if noFlagsSet {
		countBytes = true
		countLines = true
	}

	return countBytes, countLines
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

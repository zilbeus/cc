package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
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

	filename := flag.Arg(0)
	if filename == "" {
		fmt.Printf("Missing filename.\n")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf(
			"Unable to open file '%s' for reading. Does it exist?",
			filename,
		)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf(
			"Unable to query fileinfo for file '%s'.",
			filename,
		)
		return
	}

	bytes := fileInfo.Size()

	output := filename

	if countBytes {
		output = fmt.Sprintf("%d %s", bytes, output)
	}

	nrOfLines := 0
	if countLines {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			nrOfLines++
		}

		output = fmt.Sprintf("%d %s", nrOfLines, output)
	}

	fmt.Println(output)
}

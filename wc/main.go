package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var countBytes bool
	flag.BoolVar(&countBytes, "c", false, "count number of bytes in file")
	flag.Parse()

	noFlagsSet := !countBytes
	if noFlagsSet {
		fmt.Printf("No count option given. Exiting.\n")
		return
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

	if countBytes {
		fmt.Printf("%d %s\n", bytes, filename)
		return
	}
}

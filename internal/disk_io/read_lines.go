package disk_io

import (
	"bufio"
	"os"
)

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	closeFile := func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}
	defer closeFile()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 5000)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the "gridChallenge" function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
	clength := len(grid)
	rlength := len(grid[0])
	for i := 0; i < clength; i++ {
		swap := []rune(grid[i])
		for j := 0; j < rlength; j++ {
			for k := 0; k < rlength; k++ {
				if swap[j] < swap[k] {
					swap[j], swap[k] = swap[k], swap[j]
				}
			}
			grid[i] = string(swap)
		}
	}

	for i := 0; i < clength-1; i++ {
		for j := 0; j < rlength; j++ {
			if grid[i][j] > grid[i+1][j] {
				return "NO"
			}
		}
	}

	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var grid []string

		for i := 0; i < int(n); i++ {
			gridItem := readLine(reader)
			grid = append(grid, gridItem)
		}
		result := gridChallenge(grid)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

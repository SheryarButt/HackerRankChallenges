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
 * Complete the 'maxCost' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. STRING_ARRAY labels
 *  3. INTEGER dailyCount
 */

func maxCost(cost []int32, labels []string, dailyCount int32) int32 {
	// Write your code here
	var count int32
	var Cost int32
	var totalCost int32
	for i, v := range cost {
		Cost += v
		if labels[i] == "legal" {
			count++
			if count == dailyCount {
				if Cost > totalCost {
					totalCost = Cost
				}
				Cost = 0
				count = 0
			}
		}
	}
	return totalCost
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	costCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var cost []int32

	for i := 0; i < int(costCount); i++ {
		costItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		costItem := int32(costItemTemp)
		cost = append(cost, costItem)
	}

	labelsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var labels []string

	for i := 0; i < int(labelsCount); i++ {
		labelsItem := readLine(reader)
		labels = append(labels, labelsItem)
	}

	dailyCountTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	dailyCount := int32(dailyCountTemp)

	result := maxCost(cost, labels, dailyCount)

	fmt.Fprintf(writer, "%d\n", result)

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

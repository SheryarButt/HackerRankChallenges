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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {

	sum := make([]uint64, len(arr))
	var min, max uint64
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j == i {
				continue
			}
			sum[i] += uint64(arr[j])
		}
		if i == 0 {
			min = sum[i]
		}
		if max < sum[i] {
			max = sum[i]
		}
		if min > sum[i] {
			min = sum[i]
		}
	}
	fmt.Println(min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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

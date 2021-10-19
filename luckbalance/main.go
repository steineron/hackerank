package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'luckBalance' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. 2D_INTEGER_ARRAY contests
 */

type Int32Slice [][]int32

func (x Int32Slice) Len() int           { return len(x) }
func (x Int32Slice) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x Int32Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func luckBalance(k int32, contests [][]int32) int32 {
	// find the K lowest luck tournaments to loose

	sort.Sort(Int32Slice(contests))

	maxLuck := int32(0)

	j := int32(0)
	for i :=  int32(len(contests)-1); i>=0;  i-- {

		if contests[i][1] == 0 || j < k { // loose it
			maxLuck += contests[i][0]
			if contests[i][1] == 1 {
				j++ // count that loss towards K losses
			}
		} else {
			maxLuck -= contests[i][0]
		}
	}
	return maxLuck
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	var contests [][]int32
	for i := 0; i < int(n); i++ {
		contestsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var contestsRow []int32
		for _, contestsRowItem := range contestsRowTemp {
			contestsItemTemp, err := strconv.ParseInt(contestsRowItem, 10, 64)
			checkError(err)
			contestsItem := int32(contestsItemTemp)
			contestsRow = append(contestsRow, contestsItem)
		}

		if len(contestsRow) != 2 {
			panic("Bad input")
		}

		contests = append(contests, contestsRow)
	}

	result := luckBalance(k, contests)

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

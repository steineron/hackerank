// https://www.hackerrank.com/challenges/counter-game/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign&h_r=next-challenge&h_v=zen

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
 * Complete the 'counterGame' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts LONG_INTEGER n as parameter.
 */
var next map[string]string = map[string]string{
	"Louise":"Richard",
	"Richard": "Louise",
}


func subroutine(n int64, player string) string {
	if n ==1  {
		return next[player]
	}

	if n ==2  {
		return player
	}
	for i := 1; i < 64; i++ {
		p := int64(1 << i)
		if n == p {
			return subroutine(n/2, next[player])
		} else if n < p {
			return subroutine(n-(1<<(i-1)), next[player])
		}
	}
	return "ERROR"
}
func counterGame(n int64) string {
	// Write your code here

	return subroutine(n, "Louise")

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
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := counterGame(n)

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

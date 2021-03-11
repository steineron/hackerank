package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"math"
)

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
// return the next jusmp size or error if not possible
func nextJump(c []int32, location int) (int32) {
	// when on or past the before-last index, which is known to be 0 - return 0
	if location >= len(c)-2 {
		return 1
	}

	// otherwise - there are 4 possibilities for the next 2 steps
	theNext2 :=  c[location+1]<<1 + c[location+2]
	switch theNext2 {
	case 0:
		return 1 + min(nextJump(c,location+1), nextJump(c, location+2))
	case 1:
		return 1 + nextJump(c, location+1)
	case 2:
		return 1 + nextJump(c, location+2)
		// case 0x11:

	}
	return math.MaxInt32
}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	return nextJump(c,0)


}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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

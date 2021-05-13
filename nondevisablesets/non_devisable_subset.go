package main

/*
https://www.hackerrank.com/challenges/non-divisible-subset/problem
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {

	// count members of `s` into set of K-modulo remainders
	// e.g. K=4, S={5,10,20,25} will give [k:c} { 0:1, 1:2, 2:2, 3:0}

	if k == int32(0) || len(s) == 0 {
		return 0
	}

	mods := make([]int32, k)
	for _, v := range s {
		mods[v%k]++
	}

	var result int32
	// find the modulo remainder and its complement and drop the one with smallest count
	for rem, c := range mods {
		comp := (k - int32(rem)) % k     // i.e [0:k-1]
		if comp == int32(rem) && c > 0 { // can only use on number from that modulo group e.g k=8 and rem = 0 or rem =4
			result++
		} else if c > mods[comp] {
			result += c
		} else {
			result += mods[comp]
		}
		if rem >= int(k/2) { // do no repeat comparison for indexes larger than k/2 because the are the complements
			break
		}
	}

	return result
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

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	/*
	   iterate arr once - collect infromation about each encountered number n.
	   first - record and update it as a singlete - a candidate for the first occurance
	   in a triplet (n,n*r, n*r^2)
	   then check to see if it increases the number of tuples (n/r,n)
	   then check to see if it increases teh number of triplets (n/r^2, n/r, n)
	   tuples and triplets
	*/
	// if a triplet is (i,j,k) where j=i*r and k=j*r then maintain the folowing records:

	// if a number n apears in arr it may be the first in a triplet (n,n*r, n*r^2)
	// keep the number of occurences as first/i
	theIs := map[int64] int64{}

	// if a number n apears in arr it may be the second in a triplet (n/r, n, n*r)
	// keep the number of occurences as second/j
	theJs := map[int64] int64{}

	// if a number n apears in arr it may be the last in a triplet (n/r^2, n/r, n)
	// keep the number of occurences as third/j
	theKs := map[int64] int64{}

	rr := r*r
	for _,n := range arr {
		if n%r == 0 {
			if n/rr > 0 {
				theKs[n] = theKs[n]+theJs[n/r]
			}
			if n/r > 0 {
				theJs[n] = theJs[n]+theIs[n/r]
			}
		}
		theIs[n] = theIs[n]+1
	}

	var triplets int64
	for _,v := range theKs {
		triplets+=v
	}

	return triplets
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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

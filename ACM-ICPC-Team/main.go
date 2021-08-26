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

There are a number of people who will be attending ACM-ICPC World Finals.
Each of them may be well versed in a number of topics.
Given a list of topics known by each attendee, presented as binary strings, determine the maximum number of topics a
2-person team can know. Each subject has a column in the binary string, and a '1' means the subject is known
while '0' means it is not.
Also determine the number of teams that know the maximum number of topics.

Return an integer array with two elements. The first is the maximum number of topics known, and the second is the
number of teams that know that number of topics.

Example:
The attendee data is aligned for clarity below:

10101
11110
00010

These are all possible teams that can be formed:

Members Subjects
(1,2)   [1,2,3,4,5]
(1,3)   [1,3,4,5]
(2,3)   [1,2,3,4]
In this case, the first team will know all 5 subjects. They are the only team that can be created that knows that
many subjects, so [5,1] is returned.


 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
*/

func acmTeam(teamTopics []string) (result []int32) {
	// Write your code here
	result = []int32{0, 0}

	if len(teamTopics) <= 0 || len(teamTopics[0]) <= 0 {
		return result
	}
	teams := len(teamTopics)     // the number of teams
	topics := len(teamTopics[0]) // the number of topics

	// known topics
	known := make([][]int32, teams)
	for i := 0; i < teams; i++ {
		known[i] = make([]int32, teams)
		for j := i + 1; j < teams; j++ {
			for t := 0; t < topics; t++ {

				if teamTopics[i][t] == '1' || teamTopics[j][t] == '1' {
					known[i][j]++


					if known[i][j] == result[0] {
						result[1] ++
					}

					if known[i][j] > result[0] {
						result[0] = known[i][j]
						result[1] = 1
					}
				}
			}
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

	//mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	//checkError(err)
	//m := int32(mTemp)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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

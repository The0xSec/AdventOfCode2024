package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func errcheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	getLists()
}

func getLists() {
	fileName := "./index.txt"

	file, err := os.Open(fileName)
	errcheck(err)

	defer file.Close()

	var leftList []int
	var rightList []int

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Print("There is an error in the number of lists")
			continue
		}

		leftNum, err0 := strconv.Atoi(parts[0])
		rightNum, err1 := strconv.Atoi(parts[1])
		if err0 != nil || err1 != nil {
			fmt.Println("Conversion Err", line)
			continue
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scan.Err(); err != nil {
		fmt.Println(err)
		return
	}

	totalDistance := findDistance(leftList, rightList)
	fmt.Printf("The total distance between the two lists is: %d\n", totalDistance)
}

func findDistance(lft, rht []int) int {
	sort.Ints(lft)
	sort.Ints(rht)

	if len(lft) != len(rht) {
		fmt.Print("There is an error is the length of the lists")
		return 0
	}

	totalDist := 0
	for i := 0; i < len(lft); i++ {
		dist := int(math.Abs(float64(lft[i] - rht[i])))
		totalDist += dist
	}

	return totalDist
}

func split(s string, delim rune) []string {
	var res []string

	for _, part := range s {
		res = append(res, string(part))
	}
	return res
}
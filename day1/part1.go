package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("invalid file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		if len(parts) == 2 {
			left, err1 := strconv.Atoi(parts[0])
			right, err2 := strconv.Atoi(parts[1])

			if err1 == nil && err2 == nil {
				leftList = append(leftList, left)
				rightList = append(rightList, right)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error:", err)
		return
	}

	fmt.Println("lists:")
	fmt.Println(leftList, rightList)

	sort.Ints(leftList)
	sort.Ints(rightList)

	fmt.Println("\nlists:")
	fmt.Println(leftList, rightList)

	pairList := make([][]int, 0)
	for idx := 0; idx < len(leftList); idx++ {
		pairList = append(pairList, []int{leftList[idx], rightList[idx]})
	}

	fmt.Println("\n pairs:")
	fmt.Println(pairList)

	Distances := make([]int, 0)
	for idx := 0; idx < len(pairList); idx++ {
		Distances = append(Distances, pairList[idx][1] - pairList[idx][0])
	}

	fmt.Println("\n distances:")
	fmt.Println(Distances)


	TotalDistance := 0
	for idx := 0; idx < len(Distances); idx++ {
		TotalDistance += Distances[idx]
	}

	fmt.Println("\n total distance:")
	fmt.Println(TotalDistance)
}

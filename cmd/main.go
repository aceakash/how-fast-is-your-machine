package main

import (
	"fmt"
	"github.com/aceakash/sort"
	"log"
	"math/rand"
	"time"
)

func main() {
	const numOfLists = 64
	const lengthOfEachList = 20 * 1000

	lists := makeRandomListsToSort(lengthOfEachList, numOfLists)

	start := time.Now()
	for i, list := range lists {
		sort.BubbleSort(list)
		log.Printf("Sorted list %d of %d\n", i+1, numOfLists)
	}
	duration := time.Since(start)

	fmt.Printf("Took %v to sort %d lists of %d ints each", duration, numOfLists, lengthOfEachList)
}

func makeRandomListsToSort(lengthOfEachList int, numOfLists int) [][]int {
	lists := [][]int{}
	for i := 0; i < numOfLists; i++ {
		lists = append(lists, bigList(lengthOfEachList))
	}
	return lists
}

func bigList(size int) []int {
	rand.Seed(time.Now().UnixNano())

	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/aceakash/sort"
)

func main() {
	singleGoRoutine()
	multipleGoRoutines()
}

func multipleGoRoutines() {
	const numOfLists = 64
	const lengthOfEachList = 25 * 1000
	allLists := makeRandomListsToSort(lengthOfEachList, numOfLists)

	concurrency := runtime.NumCPU()
	batchSize := numOfLists / concurrency

	startTime := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(i int) {
			start := i * batchSize
			upto := (i + 1) * batchSize
			lists := allLists[start:upto]
			for _, list := range lists {
				sort.BubbleSort(list)
				//log.Printf("[singleGoRoutine] Sorted list %d of %d\n", i+1, numOfLists)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	duration := time.Since(startTime)

	fmt.Printf("[manyyGoroutines] Took %v to sort %d lists of %d ints each\n", duration, numOfLists, lengthOfEachList)
}

func singleGoRoutine() {
	const numOfLists = 64
	const lengthOfEachList = 25 * 1000

	lists := makeRandomListsToSort(lengthOfEachList, numOfLists)

	start := time.Now()
	for i, list := range lists {
		sort.BubbleSort(list)
		log.Printf("[singleGoRoutine] Sorted list %d of %d\n", i+1, numOfLists)
	}
	duration := time.Since(start)
	fmt.Println("====================================")
	fmt.Printf("[singleGoRoutine] Took %v to sort %d lists of %d ints each\n", duration, numOfLists, lengthOfEachList)
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

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, function string, sleepSeconds int) {
	defer wg.Done()
	fmt.Printf("Workter %v: Started\n", function)
	time.Sleep(time.Second * time.Duration(sleepSeconds))
	fmt.Printf("Workere %v: finished\n", function)
}

func main() {
	fmt.Println("hellow")
	var wg sync.WaitGroup
	for index, function := range []string{"FunctionOne", "FunctionTwo", "FunctionThree"} {
		wg.Add(1)
		go worker(&wg, function, index)
		fmt.Println(function)
	}
	wg.Wait()
}

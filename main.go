package main

import (
	"fmt"
	"sync"
)

/*
Pessimistic Locking in Go to ensure correct concurrent operations
count++ is not an atomic operation and hence it is the critical section

Perfomance takes a hit when doing pessimistic locking because all the threads have to wait for one thread to finish its execution
*/


var count int = 0
var wg sync.WaitGroup
var mu sync.Mutex

func incCount(){
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func doCount(){
	for i:=0;i<1000000;i++ {
		wg.Add(1)
		go incCount()
	}
}

func main(){
	count = 0
	doCount()
	wg.Wait()
	fmt.Println(count)
}

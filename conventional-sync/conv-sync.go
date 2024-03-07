package main

import (
	"fmt"
	"sync"
)

func do() int{
	m := make(chan bool, 1)

	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++{
		w.Add(1)
		go func(){
			m <- true
			n++
			<-m
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

func main(){
	fmt.Println(do())
}
package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	hellos := make(chan string, 20)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go getHello(i, hellos, wg)
	}
	wg.Wait()
	close(hellos)
	for greeting := range hellos {
		fmt.Println(greeting)
	}
}

func getHello(i int, hellos chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	hellos <- "Hello World: " + strconv.Itoa(i)
}

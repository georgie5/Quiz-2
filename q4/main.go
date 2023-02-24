package main

import (
	"fmt"
	"sync"
)

func downloads(id int, wg *sync.WaitGroup) {
	fmt.Printf("download %d starting\n", id)

	//finished download
	fmt.Printf("download %d finished\n", id)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go downloads(1, &wg)
	go downloads(2, &wg)
	go downloads(3, &wg)

	wg.Wait()
	fmt.Println("All downloads finished")

}

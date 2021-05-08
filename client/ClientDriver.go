package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Please enter your User ID: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userID := scanner.Text()
	wg.Add(1)
	go func() {
		var client Client = Client{}
		client.init(userID)
		client.start()
		wg.Done()
	}()
	wg.Wait()
}

package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"sync"
)

const CONPORT = "127.0.0.1:1301"
const MCGPORT = "127.0.0.1:1302"
const MONPORT = "127.0.0.1:1303"

func main() {
	var wg sync.WaitGroup
	fmt.Println("Please enter your User ID: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userID := scanner.Text()
	wg.Add(1)
	go func() {
		var library = userID[0:3]
		var address string
		if library == "CON" {
			address = CONPORT
		} else if library == "MCG" {
			address = MCGPORT
		} else if library == "MON" {
			address = MONPORT
		}
		httpclient, err := rpc.Dial("tcp", address)
		if err != nil {
			fmt.Printf("dial(%q): %s\n", address, err)
			return
		}
		var client = Client{}
		client.init(userID, httpclient)
		client.start()
		wg.Done()
	}()
	wg.Wait()
}

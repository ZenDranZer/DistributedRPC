package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	itemID    string
	library   string
	index     string
	itemName  string
	itemCount int64
}

func (i *Item) init(itemID string) {
	i.itemID = itemID
	i.library = string(itemID[0:3])
	i.index = string(itemID[3:])
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("The item id : %q \n", itemID)
	fmt.Println("Enter Item name: ")
	scanner.Scan()
	i.itemName = scanner.Text()
	fmt.Println("Enter Item count: ")
	scanner.Scan()
	i.itemCount, _ = strconv.ParseInt(scanner.Text(), 10, 64)
}

func (i Item) getItemID() string {
	return i.itemID
}

func (i Item) getItemName() string {
	return i.itemName
}

func (i Item) getItemCount() int64 {
	return i.itemCount
}

func (i *Item) setItemCount(count int64) {
	i.itemCount = count
}

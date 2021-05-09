package main

type Item struct {
	itemID    string
	library   string
	index     string
	itemName  string
	itemCount int
}

func (i *Item) init(itemID string, itemName string, itemCount int) {
	i.itemID = itemID
	i.library = itemID[0:3]
	i.index = itemID[3:]
	i.itemName = itemName
	i.itemCount = itemCount
}

func (i Item) getItemID() string {
	return i.itemID
}

func (i Item) getItemName() string {
	return i.itemName
}

func (i Item) getItemCount() int {
	return i.itemCount
}

func (i *Item) setItemCount(count int) {
	i.itemCount = count
}

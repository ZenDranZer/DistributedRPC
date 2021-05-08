package main

import "fmt"

type Client struct {
	userID   string
	library  string
	userType string
	index    string
}

func (c *Client) init(userID string) {
	c.userID = userID
	c.library = string(userID[0:3])
	c.userType = string(userID[3:4])
	c.index = string(userID[4:])
}

func (c *Client) start() {
	isValid := c.validateClient()
	fmt.Println(isValid)
}

func (c *Client) validateClient() bool {
	return true
}

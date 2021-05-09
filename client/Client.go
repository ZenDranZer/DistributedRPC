package main

import (
	"fmt"
	"net/rpc"
)

type Client struct {
	userID     string
	library    string
	userType   string
	index      string
	httpclient *rpc.Client
}

func (c *Client) init(userID string, htttpclient *rpc.Client) {
	c.userID = userID
	c.library = userID[0:3]
	c.userType = userID[3:4]
	c.index = userID[4:]
	c.httpclient = htttpclient
}

func (c *Client) start() {
	isValid := c.validateClient()
	fmt.Println(isValid)
}

func (c *Client) validateClient() bool {
	var reply = true
	c.httpclient.Call("Server.ValidateClient", c.userID, &reply)
	return reply
}

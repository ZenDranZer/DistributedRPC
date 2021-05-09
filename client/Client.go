package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
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
	if isValid && c.userType == "U" {
		c.userMenu()
	} else if isValid && c.userType == "M" {
		c.managerMenu()
	} else {
		fmt.Println("The userID is not valid Please enter a valid UserID.")
	}
}

func (c *Client) userMenu() {
	option := "Y"
	fmt.Printf("Hello, %q \n", c.userID)
	for option == "Y" {
		printUserMenu()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option = scanner.Text()
		var reply string
		var err error
		switch option {
		case "1":
			fmt.Println("Borrow Item Section :")
			fmt.Println("Enter Item ID: ")
			scanner.Scan()
			itemID := scanner.Text()
			fmt.Println("Enter for how many days you want to borrow ?")
			scanner.Scan()
			numberOfDays := scanner.Text()
			err = c.httpclient.Call("Server.BorrowItem", []string{itemID, numberOfDays, c.userID}, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "2":
			fmt.Println("Find Item by Name Section :")
			fmt.Println("Enter Item Name: ")
			scanner.Scan()
			itemName := scanner.Text()
			err = c.httpclient.Call("Server.FindItemByName", itemName, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "3":
			fmt.Println("Find Item by ID Section :")
			fmt.Println("Enter Item ID: ")
			scanner.Scan()
			itemID := scanner.Text()
			err = c.httpclient.Call("Server.FindItemByID", itemID, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "4":
			fmt.Println("Return Item Section :")
			fmt.Println("Enter Item ID: ")
			scanner.Scan()
			itemID := scanner.Text()
			err = c.httpclient.Call("Server.ReturnItem", []string{itemID, c.userID}, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "N", "n":
			fmt.Printf("User Quit : UserID : %q \n", c.userID)
		default:
			fmt.Println("Wrong Selection")
			option = "Y"

		}
		if err != nil {
			fmt.Println("Server side error. Please try again later.")
		}
	}
}

func (c *Client) managerMenu() {
	option := "Y"
	fmt.Printf("Hello, %q \n", c.userID)
	for option == "Y" {
		printManagerMenu()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option = scanner.Text()
		var reply string
		var multiReply []string
		var err error
		switch option {
		case "1":
			fmt.Println("Add Item Section :")
			fmt.Println("Enter Item ID: ")
			scanner.Scan()
			itemID := scanner.Text()
			fmt.Println("Enter Item Name: ")
			scanner.Scan()
			itemName := scanner.Text()
			fmt.Println("Enter Item Quantity: ")
			scanner.Scan()
			itemQTY := scanner.Text()
			err = c.httpclient.Call("Server.AddItem", []string{itemID, itemName, itemQTY}, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "2":
			fmt.Println("Remove Item Section :")
			fmt.Println("Enter Item ID: ")
			scanner.Scan()
			itemID := scanner.Text()
			fmt.Println("Enter Quantity to remove: ")
			scanner.Scan()
			itemQTY := scanner.Text()
			err = c.httpclient.Call("Server.RemoveItem", []string{itemID, itemQTY}, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "3":
			fmt.Println("List Item availability Section :")
			err = c.httpclient.Call("Server.ListAvailability", c.userID, &multiReply)
			fmt.Println("Reply from server")
			for _, str := range multiReply {
				fmt.Println(str)
			}
			option = "Y"
		case "4":
			fmt.Println("Add User Section :")
			err = c.httpclient.Call("Server.AddUser", c.userID, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "5":
			fmt.Println("Add Manager Section :")
			fmt.Println("Enter New Manager ID: ")
			scanner.Scan()
			managerID := scanner.Text()
			err = c.httpclient.Call("Server.AddManager", managerID, &reply)
			fmt.Printf("Reply from server %q", reply)
			option = "Y"
		case "N", "n":
			fmt.Printf("User Quit : UserID : %q\n", c.userID)
			option = "N"
		default:
			fmt.Println("Wrong Selection")
			option = "Y"

		}
		if err != nil {
			fmt.Println("Server side error. Please try again later.")
		}
	}
}

func (c *Client) validateClient() bool {
	var reply = true
	err := c.httpclient.Call("Server.ValidateClient", c.userID, &reply)
	if err != nil {
		fmt.Println("Server side error. Please try again later.")
	}
	return reply
}

func printUserMenu() {
	fmt.Println("\nFeatures :")
	fmt.Println("1) Borrow an item.")
	fmt.Println("2) Find an Item by Name.")
	fmt.Println("3) Find an Item by ID.")
	fmt.Println("4) Return an Item.")
	fmt.Println("Press 'N' or 'n' to exit.")
}

func printManagerMenu() {
	fmt.Println("\nFeatures :")
	fmt.Println("1) Add an item.")
	fmt.Println("2) Remove an Item.")
	fmt.Println("3) List Item Availability.")
	fmt.Println("4) Create user.")
	fmt.Println("4) Create Manager.")
	fmt.Println("Press 'N' or 'n' to exit.")
}

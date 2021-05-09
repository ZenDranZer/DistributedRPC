package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"sync"
)

const CON = "CON"
const MCG = "MCG"
const MON = "MON"
const CONPORT = ":1301"
const MCGPORT = ":1302"
const MONPORT = ":1303"

type Server struct {
	library       string
	port          string
	books         map[string]Item
	users         map[string]User
	managers      map[string]Manager
	borrowed      map[User]map[Item]int
	nextUserID    int
	nextManagerID int
}

func (s *Server) init(library string, port string) {
	s.managers = make(map[string]Manager)
	s.users = make(map[string]User)
	s.books = make(map[string]Item)
	s.borrowed = make(map[User]map[Item]int)
	s.library = library
	s.port = port

	initManagerID := s.library + "M" + "1001"
	var initManager = new(Manager)
	initManager.init(initManagerID, s.library)
	s.managers[initManagerID] = *initManager
	initUserID1001 := s.library + "U" + "1001"
	initUserID1002 := s.library + "U" + "1002"
	var initUser1001 = new(User)
	initUser1001.init(initUserID1001, s.library)
	var initUser1002 = new(User)
	initUser1002.init(initUserID1002, s.library)
	s.users[initUserID1001] = *initUser1001
	s.users[initUserID1002] = *initUser1002

	initItemID1001 := s.library + "1001"
	initItemID1002 := s.library + "1002"
	initItemID1003 := s.library + "1003"
	var initItem1001 = new(Item)
	var initItem1002 = new(Item)
	var initItem1003 = new(Item)
	initItem1001.init(initItemID1001, "Distributed Systems", 1)
	initItem1002.init(initItemID1002, "Parallel Programming", 6)
	initItem1003.init(initItemID1003, "Algorithm Designs", 7)
	s.books[initItemID1001] = *initItem1001
	s.books[initItemID1002] = *initItem1002
	s.books[initItemID1003] = *initItem1003

	s.nextUserID = 1003
	s.nextManagerID = 1002

	s.run()
}

func (s *Server) getLibrary() string {
	return s.library
}

func (s *Server) setLibrary(lib string) {
	s.library = lib
}

func (s *Server) getPort() string {
	return s.port
}

func (s *Server) run() {
	handler := rpc.NewServer()
	err := handler.Register(s)
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatal("error registering Servers", err)
	}
	log.Printf("%q server is up on %q", s.library, s.port)
	for true {
		connection, err := listener.Accept()
		if err != nil {
			log.Printf("listen(%q): %s\n", s.port, err)
			return
		}
		log.Printf("Server %q accepted connection to %s from %s\n", s.library, connection.LocalAddr(), connection.RemoteAddr())
		go handler.ServeConn(connection)
	}
}

func (s *Server) ValidateClient(clientID string, reply *bool) error {
	var isValid = true
	log.Printf("Validate user request from UserID: %q \n", clientID)
	if clientID == "" {
		isValid = false
	} else if clientID[3] == 'U' {
		if _, ok := s.users[clientID]; ok {
			isValid = true
			log.Printf(" The user is authenticated and logged in. \n")
		} else {
			log.Printf(" The user is not authenticated and can not log in. \n")
			isValid = false
		}
	} else if clientID[3] == 'M' {
		if _, ok := s.managers[clientID]; ok {
			isValid = true
			log.Printf(" The manager is authenticated and logged in. \n")
		} else {
			log.Printf(" The manager is not authenticated and can not log in. \n")
			isValid = false
		}
	}
	*reply = isValid
	return nil
}

func (s *Server) AddItem(args []string, reply *string) error {
	itemID := args[0]
	itemName := args[1]
	log.Printf("Add Item request for adding am item with item ID : %q \n", itemID)
	quantity, _ := strconv.Atoi(args[2])
	var newItem = new(Item)
	if quantity <= 0 {
		log.Printf("The manager is trying to add an item with 0 quatity")
		*reply = "0 or less quantity is not accepted. Please try again with positive integers."
	} else {
		if item, ok := s.books[itemID]; ok {
			item.itemCount += quantity
			log.Printf("The manager is trying to add an existing item. The quantiy is added to existing item. ")
			*reply = "Item already exist. Increased Item count."
		} else {
			newItem.init(itemID, itemName, quantity)
			s.books[itemID] = *newItem
			log.Printf("The item is added successfully to the database.")
			*reply = "Item added successfully to the server."
		}
	}
	log.Printf("Request : Add Item | Server : %q | Reply : %q", s.library, *reply)
	return nil
}

func (s *Server) RemoveItem(args []string, reply *string) error {
	itemID := args[0]
	quantity, _ := strconv.Atoi(args[1])
	if item, ok := s.books[itemID]; ok {
		if quantity == 0 || item.itemCount <= quantity {
			delete(s.books, itemID)
			log.Printf("The item %s is removed from the database.", itemID)
			*reply = "Item completely removed from the server."
		} else {
			item.itemCount -= quantity
			log.Printf("The item %s is exists in the database. The item quantity is reduced.", itemID)
			*reply = "Item already exist. Decreased Item count."
		}
	} else {
		*reply = "Item does not exist in this server. Please send valid itemID."
	}
	log.Printf("Request : Remove Item | Server : %q | Reply : %q", s.library, *reply)
	return nil
}

func (s *Server) ListAvailability(managerID string, reply *[]string) error {
	var availableItems []string
	for _, item := range s.books {
		str := fmt.Sprintf("Item Name: %s Item Quantity: %d ", item.itemName, item.itemCount)
		availableItems = append(availableItems, str)
	}
	*reply = availableItems
	log.Printf("Request : Available Items | Server : %q | Reply : %q", s.library, *reply)
	return nil
}

func (s *Server) AddUser(managerID string, reply *string) error {
	userID := fmt.Sprintf("%sU%d", s.library, s.nextUserID)
	newUser := new(User)
	newUser.init(userID, s.library)
	s.users[userID] = *newUser
	s.nextUserID++
	*reply = fmt.Sprintf("New user added with UserID: %s", userID)
	log.Printf("Request : Add User | Server : %q | Reply : %q", s.library, *reply)
	return nil
}

func (s *Server) AddManager(managerID string, reply *string) error {
	userID := fmt.Sprintf("%sM%d", s.library, s.nextManagerID)
	newManager := new(Manager)
	newManager.init(userID, s.library)
	s.managers[userID] = *newManager
	s.nextManagerID++
	*reply = fmt.Sprintf("New user added with UserID: %s", userID)
	log.Printf("Request : Add Manager | Server : %q | Reply : %q", s.library, *reply)
	return nil
}

func (s *Server) BorrowItem(args []string, reply *string) error {
	itemID := args[0]
	numberOfDays, _ := strconv.Atoi(args[1])
	userID := args[2]
	user := s.users[userID]
	//If item is available in the library then
	if item, ok := s.books[itemID]; ok && item.itemCount > 0 {
		//If item is available and user have already borrowed some books the
		if borrowedItems, ok := s.borrowed[user]; ok {
			//check whether the user has already borrowed the book
			if _, ok := borrowedItems[item]; ok {
				//can not borrow again
				*reply = "You have already borrowed this item. You can borrow an item once."
				log.Printf("The user %s has already borrowed the item %s.", userID, itemID)
			} else {
				*reply = "Congratulations! You have borrowed this item."
				log.Printf("The user %s has borrowed the item %s.", userID, itemID)
				borrowedItems[item] = numberOfDays
				item.itemCount--
				s.books[itemID] = item
				s.borrowed[user] = borrowedItems
			}
		} else {
			borrowedItem := make(map[Item]int)
			item.itemCount--
			borrowedItem[item] = numberOfDays
			s.books[itemID] = item
			s.borrowed[user] = borrowedItem
			log.Printf("The user %s has already borrowed the item %s.", userID, itemID)
			*reply = "Congratulations! You have borrowed your first item."
		}
	} else { //Item not available in library
		*reply = "Unfortunately, the item is not available at our library. Please try again later."
		log.Printf("The user %s has already borrowed the item %s.", userID, itemID)
	}

	log.Printf("Request : Borrow Item | Server : %q | Reply : %q", s.library, *reply)
	return nil
}

func (s *Server) FindItemByName(itemName string, reply *string) error {
	var found = false
	for key, item := range s.books {
		if item.itemName == itemName {
			found = true
			*reply = fmt.Sprintf("Item found. Item ID: %s | Item Quantity: %d ", key, item.itemCount)
			log.Printf("The user has query for the item %s. The item is found with %s and %d", itemName, item.itemID, item.itemCount)
		}
	}
	if !found {
		*reply = "Unfortunately, there is no item available with this name at our library. Please try again later"
		log.Printf("The user has query for the item %s. Item not available at this library.", itemName)
	}
	log.Println(*reply)
	log.Printf("Request : Find Item by name | Server : %q | Reply : %q", s.library, *reply)

	return nil
}

func (s *Server) FindItemByID(itemID string, reply *string) error {
	if item, ok := s.books[itemID]; ok {
		*reply = fmt.Sprintf("Item found. Item Name: %s | Item Quantity: %d ", item.itemName, item.itemCount)

		log.Printf("The user has query for the item %s. The item is found with %s and %d", itemID, item.itemName, item.itemCount)
	} else {
		*reply = "Unfortunately, there is no item available with this ID at our library. Please try again later"
		log.Printf("The user has query for the item %s. Item not available at this library.", itemID)
	}
	log.Println(*reply)
	return nil
}

func (s *Server) ReturnItem(args []string, reply *string) error {
	itemID := args[0]
	userID := args[1]
	user := s.users[userID]

	if borrowedItems, ok := s.borrowed[user]; ok {
		var found = false
		for item, _ := range borrowedItems {
			if item.itemID == itemID {
				found = true
				delete(borrowedItems, item)
				*reply = "You have successfully returned the item."
				log.Printf("The user %s returned the item %s", userID, itemID)
				item.itemCount++
				s.books[itemID] = item
				break
			}
		}
		if !found {
			log.Printf("The user %s wants to returned the item %s. However, they have not yet borrowed it.", userID, itemID)
			*reply = "You have not borrowed this item. Please return the item you have already borrowed."
		}

	} else {
		log.Printf("The user %s wants to returned the item %s. However, they have not yet borrowed anything from this library.", userID, itemID)
		*reply = "You have not borrowed anything from this library. Please borrow a book in order to return it."
	}
	return nil
}

var concordia = new(Server)
var mcgill = new(Server)
var montreal = new(Server)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		concordia.init(CON, CONPORT)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		mcgill.init(MCG, MCGPORT)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		montreal.init(MON, MONPORT)
		wg.Done()
	}()
	wg.Wait()
}

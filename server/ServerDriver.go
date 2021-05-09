package main

import (
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

type API int

type Server struct {
	library           string
	port              string
	books             map[string]Item
	users             map[string]User
	managers          map[string]Manager
	borrowed          map[User]map[Item]int
	borrowedItemsDays map[string]int
	nextUserID        int
	nextManagerID     int
}

func (s *Server) init(library string, port string) {
	s.managers = make(map[string]Manager)
	s.users = make(map[string]User)
	s.books = make(map[string]Item)
	s.borrowed = make(map[User]map[Item]int)
	s.borrowedItemsDays = make(map[string]int)
	s.library = library
	s.port = port

	initManagerID := s.library + "M" + "1001"
	var initManager = new(Manager)
	initManager.init(initManagerID, s.library, "1001")
	s.managers[initManagerID] = *initManager
	initUserID1001 := s.library + "U" + "1001"
	initUserID1002 := s.library + "U" + "1002"
	var initUser1001 = new(User)
	initUser1001.init(initUserID1001, s.library, "1001")
	var initUser1002 = new(User)
	initUser1002.init(initUserID1002, s.library, "1002")
	s.users[initUserID1001] = *initUser1001
	s.users[initUserID1002] = *initUser1002

	initItemID1001 := s.library + "1001"
	initItemID1002 := s.library + "1002"
	initItemID1003 := s.library + "1003"
	var initItem1001 = new(Item)
	var initItem1002 = new(Item)
	var initItem1003 = new(Item)
	initItem1001.init(initItemID1001, "Distributed Systems", 1)
	initItem1002.init(initItemID1001, "Parallel Programming", 6)
	initItem1003.init(initItemID1001, "Algorithm Designs", 7)
	s.books[initItemID1001] = *initItem1001
	s.books[initItemID1002] = *initItem1002
	s.books[initItemID1003] = *initItem1003
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
	if clientID == "" {
		isValid = false
	} else if clientID[3] == 'U' {
		if _, ok := s.users[clientID]; ok {
			isValid = true
		} else {
			isValid = false
		}
	} else if clientID[3] == 'M' {
		if _, ok := s.managers[clientID]; ok {
			isValid = true
		} else {
			isValid = false
		}
	}
	log.Printf("Validate user request : %q \n Result : %t", clientID, isValid)
	*reply = isValid
	return nil
}

func (s *Server) AddItem(args []string, reply *string) error {
	itemID := args[0]
	itemName := args[1]
	var newItem = new(Item)
	quantity, _ := strconv.Atoi(args[2])
	if item, ok := s.books[itemID]; ok {
		item.itemCount += quantity
		*reply = "Item already exist. Increased Item count."
	} else {
		newItem.init(itemID, itemName, quantity)
		s.books[itemID] = *newItem
		*reply = "Item added successfully to the server."
	}
	log.Printf("Request : Add Item | Server : %q | Reply : %q", s.library, *reply)
	return nil
}

func (s *Server) RemoveItem(args []string, reply *string) error {
	itemID := args[0]
	quantity, _ := strconv.Atoi(args[1])
	if item, ok := s.books[itemID]; ok {
		if quantity == 0 || item.itemCount < quantity {
			delete(s.books, itemID)
			*reply = "Item completely removed from the server."
		} else {
			item.itemCount -= quantity
			*reply = "Item already exist. Decreased Item count."
		}
	} else {
		*reply = "Item does not exist in this server. Please send valid itemID."
	}
	log.Printf("Request : Add Item | Server : %q | Reply : %q", s.library, *reply)
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

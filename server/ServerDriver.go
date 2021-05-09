package main

import (
	"log"
	"net"
	"net/rpc"
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

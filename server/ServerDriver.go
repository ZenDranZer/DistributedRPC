package main

import "fmt"

const CON = "CON"
const MCG = "MCG"
const MON = "MON"
const CONPORT = 1301
const MCGPORT = 1302
const MONPORT = 1303

type Server struct {
	library           string
	port              int
	books             map[string]Item
	users             map[string]User
	managers          map[string]Manager
	borrowed          map[User]map[Item]int
	waitingQueue      map[string]map[string]int
	borrowedItemsDays map[string]int
	nextUserID        int
	nextManagerID     int
}

func (s *Server) init(library string, port int) {
	s.managers = make(map[string]Manager)
	s.users = make(map[string]User)
	s.books = make(map[string]Item)
	s.borrowed = make(map[User]map[Item]int)
	s.waitingQueue = make(map[string]map[string]int)
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
}

func (s *Server) getLibrary() string {
	return s.library
}

func (s *Server) setLibrary(lib string) {
	s.library = lib
}

func (s *Server) getPort() int {
	return s.port
}

var concordia = new(Server)
var mcgill = new(Server)
var montreal = new(Server)

func main() {

	concordia.init(CON, CONPORT)
	mcgill.init(MCG, MCGPORT)
	montreal.init(MON, MONPORT)

	fmt.Println(concordia)
	fmt.Println(mcgill)
	fmt.Println(montreal)
}

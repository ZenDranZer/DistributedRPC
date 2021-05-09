package main

type User struct {
	userID     string
	library    string
	index      string
	outsourced [3]bool
}

func (u *User) init(userID string, library string) {
	u.userID = userID
	u.library = library
	u.index = userID[4:]
}

func (u User) getOutsourced() [3]bool {
	return u.outsourced
}

func (u *User) setOutsourced(b [3]bool) {
	u.outsourced = b
}

func (u User) getUserID() string {
	return u.userID
}

func (u *User) setUserID(s string) {
	u.userID = s
}

func (u User) getLibrary() string {
	return u.library
}

func (u *User) setLibrary(s string) {
	u.library = s
}

func (u User) getIndex() string {
	return u.index
}

func (u *User) setIndex(s string) {
	u.index = s
}

package main

type Manager struct {
	managerID string
	library   string
	index     string
}

func (m *Manager) init(managerID string, library string, index string) {
	m.managerID = managerID
	m.library = library
	m.index = index
}

func (m Manager) getManagerID() string {
	return m.managerID
}

func (m *Manager) setManagerID(s string) {
	m.managerID = s
}

func (m Manager) getLibrary() string {
	return m.library
}

func (m *Manager) setLibrary(s string) {
	m.library = s
}

func (m Manager) getIndex() string {
	return m.index
}

func (m *Manager) setIndex(s string) {
	m.index = s
}

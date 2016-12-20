package main

import "sync"

type record struct {
	Amount int
	Name   string
}

type database struct {
	sync.RWMutex
	records []record
}

func NewDatabase() *database {
	return &database{records: make([]record, 0, 4)}
}

func (db *database) get(index int) (r record, ok bool) {
	db.RLock()
	defer db.RUnlock()
	if !db.isValidIndex(index) {
		return record{}, false
	}

	return db.records[index], true
}

func (db *database) isValidIndex(index int) bool {
	return index >= 0 && index < db.length()
}

func (db *database) length() int {
	db.RLock()
	defer db.RUnlock()
	return len(db.records)
}

func (db *database) clear() {
	db.Lock()
	defer db.Unlock()
	db.records = make([]record, 0, 4)
}

func (db *database) remove(index int) bool {
	db.Lock()
	defer db.Unlock()
	if !db.isValidIndex(index) {
		return false
	}

	db.records = append(db.records[:index], db.records[index+1:]...)

	return true
}

func (db *database) insert(index int, r record) bool {
	db.Lock()
	defer db.Unlock()
	if index < 0 || index > db.length() {
		return false
	}

	db.records = append(db.records, record{})
	copy(db.records[index+1:], db.records[index:])
	db.records[index] = r

	return true
}

func (db *database) update(index int, r record) bool {
	db.Lock()
	defer db.Unlock()
	if !db.isValidIndex(index) {
		return false
	}

	db.records[index] = r

	return true
}

package database

import "sync"

type Record struct {
	Amount int
	Name   string
}

type Database struct {
	sync.RWMutex
	records []Record
}

func NewDatabase() *Database {
	return &Database{records: make([]Record, 0, 4)}
}

func (db *Database) Get(index int) (r Record, ok bool) {
	db.RLock()
	defer db.RUnlock()
	if !db.isValidIndex(index) {
		return Record{}, false
	}

	return db.records[index], true
}

func (db *Database) isValidIndex(index int) bool {
	return index >= 0 && index < db.length()
}

func (db *Database) Length() int {
	db.RLock()
	defer db.RUnlock()
	return db.length()
}

func (db *Database) length() int {
	return len(db.records)
}

func (db *Database) Clear() {
	db.Lock()
	defer db.Unlock()
	if len(db.records) == 0 {
		return
	}
	db.records = make([]Record, 0, 4)
}

func (db *Database) Remove(index int) bool {
	db.Lock()
	defer db.Unlock()
	if !db.isValidIndex(index) {
		return false
	}

	db.records = append(db.records[:index], db.records[index+1:]...)

	return true
}

func (db *Database) Insert(index int, r Record) bool {
	db.Lock()
	defer db.Unlock()
	if index < 0 || index > db.length() {
		return false
	}

	db.records = append(db.records, Record{})
	copy(db.records[index+1:], db.records[index:])
	db.records[index] = r

	return true
}

func (db *Database) Update(index int, r Record) bool {
	db.Lock()
	defer db.Unlock()
	if !db.isValidIndex(index) {
		return false
	}

	db.records[index] = r

	return true
}

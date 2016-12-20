package main

type record struct {
	amount int
	name   string
}

type database struct {
	records []record
}

func NewDatabase() *database {
	return &database{records: make([]record, 4)}
}

func (db *database) get(index int) (r record, ok bool) {
	if !db.isValidIndex(index) {
		return record{}, false
	}

	return db.records[index], true
}

func (db *database) isValidIndex(index int) bool {
	return index >= 0 && index < db.length()
}

func (db *database) length() int {
	return len(db.records)
}

func (db *database) clear() {
	db.records = make([]record, 4)
}

func (db *database) remove(index int) bool {
	if !db.isValidIndex(index) {
		return false
	}

	db.records = append(db.records[:index], db.records[index+1:]...)

	return true
}

func (db *database) insert(index int, r record) bool {
	if index < 0 || index > db.length() {
		return false
	}

	db.records = append(db.records, record{})
	copy(db.records[index+1:], db.records[index:])
	db.records[index] = r

	return true
}

func (db *database) update(index int, r record) bool {
	if !db.isValidIndex(index) {
		return false
	}

	db.records[index] = r

	return true
}

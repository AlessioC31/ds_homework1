package main

type Author struct {
	Name    string
	Surname string
}

type Publisher struct {
	Name string
}

type Date struct {
	Day   uint32
	Month uint32
	Year  uint32
}

type Time struct {
	Hours   uint32
	Minutes uint32
}

type State int

const (
	Available State = iota
	NotAvailable
)

type Book struct {
	Authors   []Author `xml:"Authors>Author"`
	Publisher Publisher
	Date      Date
	Pages     uint32
	State     State
}

type Library struct {
	Books       []Book `xml:"Books>Book"`
	OpeningTime Time
	ClosingTime Time
}

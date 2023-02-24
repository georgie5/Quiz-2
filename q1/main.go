package main

import "fmt"

//an interface is what we call an Abstract.
// it means that it is like a class but there is no implementation

type Book interface {
	getTitle() string
	getAuthor() string
	datePublished() string
}

// Interfaces in Go are just a set of method signatures that define a behavior
//of an object.

// ---------struct fiction books
type Fiction struct {
	title         string
	author        string
	datepublished string
	genre         string
}

func (f Fiction) getTitle() string {
	return f.title
}
func (f Fiction) getAuthor() string {
	return f.author
}
func (f Fiction) datePublished() string {
	return f.datepublished
}

// ---------struct non fiction books
type NonFiction struct {
	title         string
	author        string
	datepublished string
	topic         string
}

func (n NonFiction) getTitle() string {
	return n.title
}
func (n NonFiction) getAuthor() string {
	return n.author
}
func (n NonFiction) datePublished() string {
	return n.datepublished
}

func PrintInfo(b Book) {
	fmt.Println("Title of the book is ", b.getTitle(),
		"and the Author is", b.getAuthor(),
		" and it was published on", b.datePublished())
}

//-----------------main

func main() {
	//create a fiction book
	fictionbook := Fiction{
		title:         "The day I turned invisible",
		author:        "Jane",
		datepublished: "Jan 23, 2010",
		genre:         "Fantasy",
	}

	//create a non fiction book
	nonfictionbook := NonFiction{
		title:         "World war II",
		author:        "Doe",
		datepublished: "Feb 20, 2003",
		topic:         "Historical",
	}

	PrintInfo(fictionbook)
	PrintInfo(nonfictionbook)
	fmt.Println(fictionbook.genre)
}

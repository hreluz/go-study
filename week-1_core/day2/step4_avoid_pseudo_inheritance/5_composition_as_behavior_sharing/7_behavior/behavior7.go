package main

import "fmt"

type Saver interface {
	Save() string
}

type File struct{}

func (File) Save() string { return "saving to file" }

type Database struct{}

func (Database) Save() string { return "saving to database" }

func Persist(s Saver) {
	fmt.Println(s.Save())
}

func main() {
	Persist(File{})
	Persist(Database{})
}

/**
	Note:
	Neither File nor Database declares:

	implements Saver
	extends Saver
	inherits Saver

	Go infers interface satisfaction purely from behavior.
**/

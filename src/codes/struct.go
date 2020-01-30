package main

import "fmt"

type book struct {
	title   string
	arthur  string
	subject string
	id      int
}

func main() {
	var book1 book
	book1.title = "aaa"
	book1.arthur = "allen"
	book1.subject = "wwww"
	book1.id = 10000

	book2 := book{"bbb", "hommy", "rrr", 9999}

	fmt.Println(book1.title, book1.arthur, book1.subject, book1.id)
	fmt.Println(book2.title, book2.arthur, book2.subject, book2.id)

	printbook(book2)
	printbook2(&book2)

}

func printbook(bk book) {
	fmt.Println(bk.title, bk.arthur, bk.subject, bk.id)
}

func printbook2(bk *book) {
	fmt.Println(bk.title, bk.arthur, bk.subject, bk.id)
}

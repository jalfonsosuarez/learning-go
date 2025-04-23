package book

import "fmt"

// ! Interface con métodos obligatorios en las estructuras.
type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

//! Los atributos y métodos públicos deben tener la primera letra mayúscula:
// type Book struct {
// 	Title  string
// 	Author string
// 	Pages  int
// }

// ! Para esteblecer atributos y métodos privados, la primera letra debe ser minúscula.
type Book struct {
	title  string
	author string
	pages  int
}

// ? Constructor en Go
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTile(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b *Book) GetPages() int {
	return b.pages
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// ! Composición, permite crear una clase a partir de otra.
// ! Lo que en otros lenguajes es la herencia.
type TextBook struct {
	Book
	editorial string
	level     string
}

// ! Constructor
func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book: Book{
			title:  title,
			author: author,
			pages:  pages,
		},
		editorial: editorial,
		level:     level,
	}
}

func (tb *TextBook) SetEditorial(editorial string) {
	tb.editorial = editorial
}

func (tb *TextBook) GetEditorial() string {
	return tb.editorial
}

func (tb *TextBook) SetLevel(level string) {
	tb.level = level
}

func (tb *TextBook) GetLevel() string {
	return tb.level
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n",
		b.title, b.author, b.pages, b.editorial, b.level)
}

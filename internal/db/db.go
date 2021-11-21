package db

import "github.com/morzhanov/neptune-example/internal/config"

// TODO: implement NeptuneDB connection and operations
//		create struct and interface
//		perform connection
//		add specific (non-general) CRUD methods (AddReader, AddBook, etc.)

type DB interface {
	AddReader(r *Reader) error
	DeleteReader(id string) error
	ReadBook(rID string, title string) error
	GetReaders() ([]*Reader, error)
	AddBook(b *Book, a *Author) error
	UpdateBook(title string, b *Book) error
	DeleteBook(title string) error
	GetBooks() ([]*Book, error)
	AddAuthor(a *Author) error
	DeleteAuthor(id string) error
	GetAuthors() ([]*Author, error)
}

type db struct{}

func (d *db) AddReader(r *Reader) error {
	panic("implement me")
}

func (d *db) DeleteReader(id string) error {
	panic("implement me")
}

func (d *db) ReadBook(rID string, title string) error {
	panic("implement me")
}

func (d *db) GetReaders() ([]*Reader, error) {
	panic("implement me")
}

func (d *db) AddBook(b *Book, a *Author) error {
	panic("implement me")
}

func (d *db) UpdateBook(title string, b *Book) error {
	panic("implement me")
}

func (d *db) DeleteBook(title string) error {
	panic("implement me")
}

func (d *db) GetBooks() ([]*Book, error) {
	panic("implement me")
}

func (d *db) AddAuthor(a *Author) error {
	panic("implement me")
}

func (d *db) DeleteAuthor(id string) error {
	panic("implement me")
}

func (d *db) GetAuthors() ([]*Author, error) {
	panic("implement me")
}

func NewDB(c *config.Config) DB {
	// TODO: create db and perform connection
}

package service

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/morzhanov/neptune-example/internal/db"
	"go.uber.org/zap"
)

// TODO: implement service which:
//		creates/updates/deletes neptune vertexes and edges
//		logs all operations

// TODO: service should receive DB as a dependency

type Service interface {
	Run() error
}

type service struct {
	db  db.DB
	log *zap.Logger
}

func generateAuthor() *db.Author {
	var d db.Author
	gofakeit.Struct(&d)
	return &d
}

func generateBook() *db.Book {
	var d db.Book
	gofakeit.Struct(&d)
	return &d
}

func generateReader() *db.Reader {
	var d db.Reader
	gofakeit.Struct(&d)
	return &d
}

func (s *service) logAllResources() error {
	authors, err := s.db.GetAuthors()
	if err != nil {
		return err
	}
	books, err := s.db.GetBooks()
	if err != nil {
		return err
	}
	readers, err := s.db.GetReaders()
	if err != nil {
		return err
	}
	s.log.Info("Authors:")
	for _, a := range authors {
		s.log.Sugar().Info(a)
	}
	s.log.Info("Books:")
	for _, b := range books {
		s.log.Sugar().Info(b)
	}
	s.log.Info("Readers:")
	for _, r := range readers {
		s.log.Sugar().Info(r)
	}
	return nil
}

func (s *service) clearResources(authors []*db.Author, books []*db.Book, readers []*db.Reader) error {
	for _, a := range authors {
		if err := s.db.DeleteAuthor(a.ID); err != nil {
			return err
		}
	}
	for _, b := range books {
		if err := s.db.DeleteAuthor(b.Title); err != nil {
			return err
		}
	}
	for _, r := range readers {
		if err := s.db.DeleteAuthor(r.ID); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Run() error {
	var authors []*db.Author
	var books []*db.Book
	var readers []*db.Reader

	for i := 0; i <= 3; i++ {
		a := generateAuthor()
		authors = append(authors, a)
		if err := s.db.AddAuthor(a); err != nil {
			return err
		}
	}
	s.log.Info("Authors generated and saved to the database...")

	for i := 0; i <= 3; i++ {
		b := generateBook()
		books = append(books, b)
		if err := s.db.AddBook(b, authors[i]); err != nil {
			return err
		}
	}
	s.log.Info("Books generated and saved to the database...")

	for i := 0; i <= 3; i++ {
		r := generateReader()
		readers = append(readers, r)
		if err := s.db.AddReader(r); err != nil {
			return err
		}
	}
	s.log.Info("Readers generated and saved to the database...")
	if err := s.logAllResources(); err != nil {
		return err
	}

	updatedBook := *books[0]
	updatedBook.Title = "New Title"
	if err := s.db.UpdateBook(books[0].Title, &updatedBook); err != nil {
		return err
	}
	s.log.Info("Updated book...")
	if err := s.logAllResources(); err != nil {
		return err
	}

	if err := s.clearResources(authors, books, readers); err != nil {
		return err
	}
	s.log.Info("All Resources destroyed...")
	return nil
}

func NewService(d db.DB, l *zap.Logger) Service {
	return &service{d, l}
}

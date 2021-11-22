package service

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/morzhanov/neptune-example/internal/db"
	"github.com/northwesternmutual/grammes/model"
	"go.uber.org/zap"
)

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
	traversal, err := s.db.Traversal()
	if err != nil {
		return err
	}
	s.log.Info("Graph Traversal: ")
	for _, t := range traversal {
		s.log.Sugar().Info(string(t))
	}
	return nil
}

func (s *service) clearResources() error {
	return s.db.ClearResources()
}

func (s *service) Run() error {
	var authors []*model.Vertex
	var books []*model.Vertex
	var readers []*model.Vertex

	for i := 0; i <= 3; i++ {
		a := generateAuthor()
		res, err := s.db.AddAuthor(a)
		if err != nil {
			return err
		}
		authors = append(authors, res)
	}
	s.log.Info("Authors generated and saved to the database...")

	for i := 0; i <= 3; i++ {
		b := generateBook()
		res, err := s.db.AddBook(b, &authors[i])
		if err != nil {
			return err
		}
		books = append(books, res)
	}
	s.log.Info("Books generated and saved to the database...")

	for i := 0; i <= 3; i++ {
		r := generateReader()
		res, err := s.db.AddReader(r)
		if err != nil {
			return err
		}
		readers = append(readers, res)
	}
	s.log.Info("Readers generated and saved to the database...")
	if err := s.logAllResources(); err != nil {
		return err
	}

	updatedBook := generateBook()
	if err := s.db.UpdateBook(books[0].ID(), updatedBook); err != nil {
		return err
	}
	s.log.Info("Updated book...")
	if err := s.logAllResources(); err != nil {
		return err
	}

	if err := s.clearResources(); err != nil {
		return err
	}
	s.log.Info("All Resources destroyed...")
	return nil
}

func NewService(d db.DB, l *zap.Logger) Service {
	return &service{d, l}
}

package db

import (
	"github.com/morzhanov/neptune-example/internal/config"
	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/model"
)

type DB interface {
	AddReader(r *Reader) (*model.Vertex, error)
	DeleteReader(id string) error
	ReadBook(readerID interface{}, bookID interface{}) error
	AddBook(b *Book, authorID interface{}) (*model.Vertex, error)
	UpdateBook(id interface{}, b *Book) error
	DeleteBook(id interface{}) error
	AddAuthor(a *Author) (*model.Vertex, error)
	DeleteAuthor(id interface{}) error
	Traversal() ([][]byte, error)
	ClearResources() error
}

type db struct {
	client *grammes.Client
}

func (d *db) connect(neptuneDBUrl string) error {
	client, err := grammes.DialWithWebSocket(neptuneDBUrl)
	if err != nil {
		return err
	}
	d.client = client
	return d.client.Connect()
}

func (d *db) AddReader(r *Reader) (*model.Vertex, error) {
	v, err := d.client.AddVertex("reader", r)
	return &v, err
}

func (d *db) DeleteReader(id string) error {
	return d.client.DropVertexByID(id)
}

func (d *db) ReadBook(readerID interface{}, bookID interface{}) error {
	r, err := d.client.VertexByID(readerID)
	if err != nil {
		return err
	}
	_, err = r.AddEdge(d.client, "read", bookID)
	return err
}

func (d *db) AddBook(b *Book, authorID interface{}) (*model.Vertex, error) {
	bv, err := d.client.AddVertex("book", b)
	if err != nil {
		return nil, err
	}
	av, err := d.client.VertexByID(authorID)
	if err != nil {
		return nil, err
	}
	_, err = av.AddEdge(d.client, "authored", bv.ID())
	return &bv, err
}

func (d *db) UpdateBook(id interface{}, b *Book) error {
	bv, err := d.client.VertexByID(id)
	if err != nil {
		return err
	}
	if err := bv.DropProperties(d.client); err != nil {
		return err
	}
	if err := bv.AddProperty(d.client, "title", b.Title); err != nil {
		return err
	}
	if err := bv.AddProperty(d.client, "description", b.Title); err != nil {
		return err
	}
	return nil
}

func (d *db) DeleteBook(id interface{}) error {
	return d.client.DropVertexByID(id)
}

func (d *db) AddAuthor(a *Author) (*model.Vertex, error) {
	v, err := d.client.AddVertex("author", a)
	return &v, err
}

func (d *db) DeleteAuthor(id interface{}) error {
	return d.client.DropVertexByID(id)
}

func (d *db) Traversal() ([][]byte, error) {
	g := grammes.Traversal()
	res, err := d.client.ExecuteQuery(g.V().Label())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *db) ClearResources() error {
	return d.client.DropAll()
}

func NewDB(c *config.Config) (DB, error) {
	db := db{}
	if err := db.connect(c.AWSNeptuneDBUrl); err != nil {
		return nil, err
	}
	return &db, nil
}

package db

type Reader struct {
	ID   string `fake:"{uuid}" json:"id"`
	Name string `fake:"{firstname}" json:"name"`
}

type Book struct {
	Title       string `fake:"{sentence:3}" json:"title"`
	Description string `fake:"{paragraph:2,2,2, }" json:"description"`
}

type Author struct {
	ID   string `fake:"{uuid}" json:"id"`
	Name string `fake:"{firstname}" json:"name"`
}

package domain

type Book struct {
	Id          int32
	Title       string
	Author      int32
	Description string
}


type BookWithAuthor struct {
    Id          int32  `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    AuthorId    int32  `json:"author_id"`
    AuthorName  string `json:"author_name"`
}
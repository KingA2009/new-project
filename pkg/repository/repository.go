package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/model"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type Book interface {
	Create(book model.CreateBook) (int, error)
	GetBookList(pagination model.Pagination, order string) (book []model.Book, err error)
	GetBookById(bookID, userID int) (model.GetBook, error)
	GetBooksByGenre(genre string) (book model.UpdateGenreOutput, err error)
	DeleteBook(bookID int) error
	UpdateBook(bookID int, input model.UpdateBookInput) error
}
type Genre interface {
	Create(genre model.Genre) (id int, err error)
	GetGenreList(pagination model.Pagination) (genre []model.Genre, err error)
	GetGenreById(genreID int) (genre model.Genre, err error)
	DeleteGenre(genreID int) error
	UpdateGenre(genreID int, input model.UpdateGenreInput) error
}
type Like interface {
	Create(like model.LikedBook) (id int, err error)
	GetLikedList(userID int) (like []model.LikeOutput, err error)
}
type Rating interface {
	CreateRating(rating model.RatingInput) (id int, err error)
}
type History interface {
	GetHistoryList() (book []model.HistoryOutput, err error)
}
type Repository struct {
	Authorization
	Book
	Genre
	Like
	Rating
	History
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Book:          NewBookRepo(db),
		Genre:         NewGenreRepo(db),
		Like:          NewLikeRepo(db),
		Rating:        NewRatingRepo(db),
		History:       NewHistoryRepo(db),
	}
}

package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/model"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
type Service struct {
	Authorization
	Book
	Genre
	Like
	Rating
	History
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Book:          NewBookService(repos.Book),
		Genre:         NewGenreService(repos.Genre),
		Like:          NewLikeService(repos.Like),
		Rating:        NewRatingService(repos.Rating),
		History:       NewHistoryService(repos.History),
	}
}

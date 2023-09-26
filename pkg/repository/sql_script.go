package repository

const (
	CreateBookQuery      = `insert into books ( name, number_of_sheets, author, year_of_book, type_of_cover, language, description, genre_id) values ($1, $2, $3, $4, $5, $6, $7, $8)`
	GetAllBooksQuery     = `SELECT * FROM books`
	GetBooksByGenreQuery = `select name, author, year_of_book, "language", type_of_cover, description, genres.genre, number_of_sheets from books inner join genres on books.genre_id = genres.id where genres.genre = $1`
	GetByIdBookQuery     = `select name, number_of_sheets, author, year_of_book, type_of_cover, "language", description, number_of_views, rating from books where id = $1`
	UpdateHistory        = `select history($1, $2);`
	UpdateBookQuery      = `update books set name = $1 , number_of_sheets = $2 , author = $3 , year_of_book = $4 , type_of_cover = $5 , "language" = $6 , description = $7 , genre_id = $8 where id = $9`
	DeleteBookQuery      = `delete from books where id = $1`
	CreateGenreQuery     = `insert into genres(id, genre) values ($1, $2)`
	GetAllGenresQuery    = `select * from genres limit $1 offset $2`
	GetByIdGenreQuery    = `select * from genres where id = $1`
	DeleteGenreQuery     = `delete from genres where id = $1`
	UpdateGenreQuery     = `update genres set genre = $1 where id = $2`
	AppendUsersQuery     = `INSERT INTO users (name, username, password_hash, email, telephone, address) values ($1, $2, $3, $4, $5, $6) returning id`
	GetUserQuery         = `SELECT id FROM users WHERE username=$1 AND password_hash=$2`
	CreateLikeQuery      = `select insert_or_select($1, $2)`
	GetLikedBooksQuery   = `select users.name, books.name, books.author, books.year_of_book, books.number_of_sheets from likedBook inner join users on users.id = likedBook.userId inner join books on books.id = likedBook.bookId where likedBook.userId = $1`
	CreateRating         = `select insert_rating($1, $2)`
	GetHistory           = `select * from history`
)

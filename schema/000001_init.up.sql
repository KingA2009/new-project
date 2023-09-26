CREATE TABLE if not exists users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    email         varchar(255) not null unique ,
    telephone     varchar(255) not null unique ,
    address       varchar(255) not null
);
CREATE TABLE likedBook(
    id     serial       not null,
    userId int not null,
    bookId int not null
);
CREATE TABLE books(
    id               serial       not null unique,
    name             varchar(255) not null unique,
    number_of_sheets varchar(255) not null ,
    author           varchar(255) not null ,
    year_of_book     varchar(255) not null ,
    type_of_cover    boolean      not null ,
    language         varchar(255) not null ,
    description      varchar(255) not null ,
    genre_id         int          not null,
    rating           int          default 0,
    number_of_views  int          default 0
);
create table genres
(
    id    serial       not null unique,
    genre varchar(255) not null
);

create table rating
(
    id     serial       not null unique,
    bookId int          not null,
    rating int          not null
);

create table history
(
    id       serial      not null unique ,
    user_id   int         not null ,
    book_id   int         not null
);


CREATE FUNCTION insert_or_select(uId int, bId int)
    RETURNS void
LANGUAGE plpgsql
AS
$$
BEGIN
  IF not exists (select userId, bookId from likedbook where bookId = $2 and userId  = $1 ) then
  		insert into likedbook (userId, bookId) values ($1,$2);
  ELSE
  delete from likedBook where bookId = $2 and userId  = $1;
  END IF;
END;
$$;



CREATE FUNCTION insert_rating(bId int, r int)
    RETURNS void
    LANGUAGE plpgsql
AS
$$
BEGIN
insert into rating(bookId, rating) values ($1, $2);
update books b set rating = (rating + $2)/2 where id = $1 and rating > 0;
update books b set rating = $2 where id = $1 and rating = 0;
END;
$$;

CREATE FUNCTION history(bId int, uId int)
    RETURNS void AS
$$
BEGIN
    update books  set number_of_views = (number_of_views + 1) where id = $1;
    insert into history(user_id, book_id) values ($2, $1);
END;
$$LANGUAGE plpgsql;

\c libraries

CREATE TABLE IF NOT EXISTS library
(
    id          SERIAL PRIMARY KEY,
    library_uid UUID UNIQUE NOT NULL,
    name        VARCHAR(80) NOT NULL,
    city        VARCHAR(255) NOT NULL,
    address     VARCHAR(255) NOT NULL
    );

CREATE TABLE IF NOT EXISTS books
(
    id        SERIAL PRIMARY KEY,
    book_uid  UUID UNIQUE NOT NULL,
    name      VARCHAR(255) NOT NULL,
    author    VARCHAR(255),
    genre     VARCHAR(255),
    condition VARCHAR(20) DEFAULT 'EXCELLENT'
    CHECK (condition IN ('EXCELLENT', 'GOOD', 'BAD'))
    );

CREATE TABLE IF NOT EXISTS library_books
(
    book_id         INT REFERENCES books(id),
    library_id      INT REFERENCES library(id),
    available_count INT NOT NULL
    );

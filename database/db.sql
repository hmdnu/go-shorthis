-- Active: 1709977213404@@127.0.0.1@3306@go_shorthis


create table URL(
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    original_url VARCHAR(225) NOT NULL,
    short_code VARCHAR(10) NOT NULL
);


SELECT * from URL;

DELETE FROM url;

DESCRIBE URL;
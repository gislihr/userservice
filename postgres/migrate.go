package postgres

import "github.com/jmoiron/sqlx"

func MigrateDB(db *sqlx.DB) error {
	_, err := db.Exec(`
		create schema if not exists userservice;
		create extension if not exists "uuid-ossp";

		create table userservice.user (
			id uuid DEFAULT uuid_generate_v4 (),
			name varchar(256) not null,
			username varchar(256) not null unique,
			email varchar(256) not null unique,
			hashed_password varchar(265) not null,
			primary key(id)
		)
	`)

	return err
}

func DropDB(db *sqlx.DB) error {
	_, err := db.Exec("drop schema if exists userservice cascade")
	return err
}

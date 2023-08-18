package storage

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	schema = `
	create table if not exists users(
		id serial primary key,
		name text not null
	);
	`
	insertUser  = `insert into users (name) values ($1);`
	selectUsers = `select name from users;`
)

type Storage struct {
	conn *pgx.Conn
}

func Open(ctx context.Context) (*Storage, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(ctx, schema)
	return &Storage{
		conn: conn,
	}, err
}

func (s *Storage) Close(ctx context.Context) error {
	return s.conn.Close(ctx)
}

func (s *Storage) StoreUser(ctx context.Context, name string) error {
	_, err := s.conn.Exec(ctx, insertUser, name)
	return err
}

func (s *Storage) AllUsers(ctx context.Context) ([]string, error) {
	rows, err := s.conn.Query(ctx, selectUsers)
	if err != nil {
		return nil, err
	}
	var names []string
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

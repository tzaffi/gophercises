package urlshort

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQLpgx struct {
	pool *pgxpool.Pool
}

func NewPostgreSQLpgx() (*PostgreSQLpgx, error) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLpgx{
		pool: pool,
	}, nil
}

func (p *PostgreSQLpgx) Close() {
	p.pool.Close()
}

type Shortener struct {
	id   int64
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p *PostgreSQLpgx) AllUrls() ([]Shortener, error) {
	rows, err := p.pool.Query(context.Background(), `SELECT id, name, url FROM "urlshort"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var res []Shortener
	for rows.Next() {
		var url Shortener
		err := rows.Scan(&url.id, &url.Name, &url.Url)
		if err != nil {
			return res, err
		}
		res = append(res, url)
	}
	if err := rows.Err(); err != nil {
		return res, err
	}

	return res, nil
}

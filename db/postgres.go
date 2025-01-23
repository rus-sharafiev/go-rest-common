package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rus-sharafiev/go-rest-common/exception"
)

type Postgres struct {
	pool *pgxpool.Pool
}

// -- Create instance -------------------------------------------------------------

var (
	connectOnce sync.Once
	Instance    Postgres
)

func Connect(connString string) {
	connectOnce.Do(func() {

		pool, err := pgxpool.New(context.Background(), connString)
		if err != nil {
			log.Fatalf("\x1b[2mPostgreSQL:\x1b[0m\x1b[31m unable to create database connection: %s \x1b[0m\n\n", err.Error())
		}

		fmt.Println("\x1b[32mСonnection to the database has been established\x1b[0m")
		Instance = Postgres{pool}
	})
}

// -- Methods ---------------------------------------------------------------------

func (p *Postgres) Query(ctx context.Context, query *string, args ...any) (pgx.Rows, error) {
	return p.pool.Query(ctx, *query, args...)
}

func (p *Postgres) QueryRow(ctx context.Context, query *string, args ...any) pgx.Row {
	return p.pool.QueryRow(ctx, *query, args...)
}

func (p *Postgres) PgxPoolClose() {
	p.pool.Close()
}

// Returns JSON string serialized by postgres
func (p *Postgres) JsonString(ctx context.Context, w http.ResponseWriter, query *string, args ...any) (string, error) {
	var result sql.NullString
	if err := p.pool.QueryRow(ctx, *query, args...).Scan(&result); err != nil {
		if err == pgx.ErrNoRows {
			return "null", nil
		} else {
			return "", err
		}
	}
	return result.String, nil
}

// Writes JSON string serialized by postgres to provided http.ResponseWriter
func (p *Postgres) WriteJsonString(ctx context.Context, w http.ResponseWriter, query *string, args ...any) {
	var result sql.NullString
	if err := p.pool.QueryRow(ctx, *query, args...).Scan(&result); err != nil {
		exception.PgxNoRows(w, err)
		return
	}
	fmt.Fprint(w, result.String)
}

// Message JSON string serialized by PostgreSQL via provided websocket connection
func (p *Postgres) MessageJsonString(ctx context.Context, conn *websocket.Conn, query *string, args ...any) error {
	if err := conn.WriteMessage(websocket.PingMessage, make([]byte, 0)); err == nil {
		var result sql.NullString
		if err = p.pool.QueryRow(ctx, *query, args...).Scan(&result); err == nil {

			if result.Valid {
				if err = conn.WriteMessage(websocket.TextMessage, []byte(result.String)); err != nil {
					return err
				}
			}

		} else if err != pgx.ErrNoRows {
			fmt.Println("error sending message")
			return err
		}
	} else {
		return err
	}

	return nil
}

func (p *Postgres) CopyFrom(ctx context.Context, tableName string, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	count, err := p.pool.CopyFrom(ctx, pgx.Identifier{tableName}, columnNames, rowSrc)
	if err != nil {
		return 0, err
	}
	return count, nil
}

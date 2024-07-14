package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/phonghaido/api-gateway/types"
)

type PostgreSQL struct {
	Conn string
}

func NewPostgreSQL(conn string) *PostgreSQL {
	return &PostgreSQL{
		Conn: conn,
	}
}

func (p *PostgreSQL) SearchUser(username, password string) (types.User, error) {
	db, err := sql.Open("postgres", p.Conn)
	if err != nil {
		return types.User{}, err
	}
	defer db.Close()

	row, err := db.Query(fmt.Sprintf("SELECT * FROM user WHERE username = %s AND password = %s", username, password))
	if err != nil {
		return types.User{}, err
	}
	defer row.Close()

	var user types.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Scope); err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (p *PostgreSQL) InsertUser(user types.User) error {
	db, err := sql.Open("postgres", p.Conn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user (username, password, email, scope) VALUES ($1, $2, $3, $4)", user.Username, user.Password, user.Email, user.Scope)
	if err != nil {
		return err
	}

	return nil
}

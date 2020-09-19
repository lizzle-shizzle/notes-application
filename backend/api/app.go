package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Client *Client
}

type Client struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	DB       *sql.DB
}

func (a *App) Run(addr string) {
	log.Printf(fmt.Sprintf("Notes API listening on %v...\n", addr[1:]))
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func (c *Client) connect() (*sql.DB, error) {
	params := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	db, err := sql.Open("postgres", params)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	return db, err
}

func (a *App) Initialize() error {
	var err error
	database, err := a.Client.connect()
	if err != nil {
		log.Println("Failed to connect to database")
		return err
	}
	a.Client.DB = database

	return nil
}

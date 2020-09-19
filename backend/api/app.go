package api

import (
	"database/sql"
	"encoding/json"
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

	// Initialize routes
	http.HandleFunc("/notes", a.NotesHandler)

	return nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal response body: %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) NotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {
	case "OPTIONS":
		return
	case "GET":
		notes, err := FetchNotes(a.Client)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				fmt.Println("No notes found")
				respondWithError(w, http.StatusNotFound, "No notes found")
			default:
				fmt.Println("Failed to fetch notes")
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if len(notes) <= 0 {
			fmt.Println("No notes found")
			respondWithError(w, http.StatusNotFound, "No notes found")
		} else {
			respondWithJSON(w, http.StatusOK, notes)
		}
	case "POST":
		var note Note

		decoder := json.NewDecoder(r.Body)
		if r.Body == nil {
			log.Println("Create note failed - Invalid payload request: empty body")
			respondWithError(w, http.StatusBadRequest, "Invalid payload request: empty body")
			return
		}
		if err := decoder.Decode(&note); err != nil {
			log.Printf("Failed to decode note: %s\n", err)
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		msg, err := CreateNote(a.Client, note)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusCreated, msg)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

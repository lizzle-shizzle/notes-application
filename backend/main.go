package main

import (
	"github.com/lizzle-shizzle/notes-application/backend/api"
)

func main() {
	a := api.App{}

	a.Run(":8080")
}

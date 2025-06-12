package main

import (
	"fmt"

	"github.com/thapakazi/go-hex-arch/internal/adapter/storage"

	"github.com/thapakazi/go-hex-arch/internal/adapter/config"
	"github.com/thapakazi/go-hex-arch/internal/adapter/http"
)

func main() {
	cfg := config.Environment
	defer storage.Database.Close()

	router := http.NewRouter()
	router.Serve()

	fmt.Println("Hello, World!", cfg)
}

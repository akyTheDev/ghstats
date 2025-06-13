package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akyTheDev/ghstats/internal/config"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Config couldn't be loaded: %v", err)
	}

	fmt.Println(cfg)
}

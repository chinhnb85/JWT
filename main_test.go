package main

import (
	"testing"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"./config"
)

func TestDr(t *testing.T){
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	configDir := flag.String("config_dir", currentDir, "find your configuration file location")

	if err := config.InitConfig(*configDir); err != nil {
		panic(err)
	}

	host := os.Getenv("HOST")
	tokenExp := os.Getenv("TOKEN_EXPIRED_MINUTES")
	
	fmt.Println("app running on " + host)
	fmt.Println("app running on " + tokenExp)
}

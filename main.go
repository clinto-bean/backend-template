package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	db "github.com/clinto-bean/backend-template/pkg/db"
	godotenv "github.com/joho/godotenv"
)

func main() {
	fmt.Println()
	log.Println("Initializing environment variables...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(fmt.Sprintf("could not load environment variables: %v", err))
	}

	dbName, dbUrl, dbUser := os.Getenv("DB_NAME"), os.Getenv("DB_URL"), os.Getenv("DB_USER")

	db, err := db.NewDB(dbName, dbUrl, dbUser)

	if err != nil {
		log.Fatalf("could not connect to database:\n\tname: %v\n\turl: %v\n\tuser: %v\n", dbName, dbUrl, dbUser)
	}

	APIPORT := os.Getenv("APIPORT")

	log.Println("Starting API server...")

	api := NewAPI(":"+APIPORT, db)
	go func() {
		api.Start()
	}()

	fmt.Println()
	fmt.Printf("Type '--help' or '-h' at any time to view commands.\n\n")

	cmdChan := make(chan string)

	go func() {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			cmdChan <- input
		}
	}()

	for {
		select {
		case cmd := <-cmdChan:
			switch cmd {
			case "exit":
				fmt.Printf("\nShutting down the server...\n\n")
				return // Exits the application
			case "status":
				fmt.Printf("\nServer is running on %v\n\n", api.Addr)
				fmt.Print("> ")
			case "help":
				fmt.Printf("\nCurrent commands:\n\thelp - displays current commands\n\texit - shut down the server\n\tstatus - display information about the server\n\trestart - restarts the server\n\n")
				fmt.Print("> ")
			case "":
				fmt.Printf("\nType '--help' or '-h' at any time to view commands.\n")
				fmt.Print("\n> ")
			case "restart":
				fmt.Printf("\nRestarting...\n\n")
				fmt.Print("> ")
			default:
				fmt.Printf("Unknown command: %s\n\n", cmd)
				fmt.Printf("\nType '--help' or '-h' at any time to view commands.\n\n")
				fmt.Print("> ")
			}
		}
	}
}

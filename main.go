package main

import (
    "fmt"
    "log"
    "net/http"
    "database" // Import the database package
    "./handlers" // Import the handlers package
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home HTTP")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello WebSocket")
}

func setupRoutes() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)

    // Game handlers
    http.HandleFunc("/start_game", handlers.StartGameHandler)
    http.HandleFunc("/shoot_ball", handlers.ShootBallHandler)

    // User handlers
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/authenticate", handlers.AuthenticateHandler)
}

func main() {
    fmt.Println("Hello World")
err := database.InitDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Set up HTTP routes
    setupRoutes()

    // Start the HTTP server
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


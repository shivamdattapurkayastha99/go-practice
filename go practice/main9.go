package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    // Start a new goroutine
    go sayHello()

    // Let the main goroutine sleep for a while to allow the goroutine to execute
    time.Sleep(100 * time.Millisecond)

    fmt.Println("Hello from main goroutine!")
}

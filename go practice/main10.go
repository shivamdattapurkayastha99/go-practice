package main

import (
    "fmt"
)

func sendData(ch chan string) {
    ch <- "Hello" // Send "Hello" to the channel
    ch <- "from"  // Send "from" to the channel
    ch <- "channel" // Send "channel" to the channel
    close(ch) // Close the channel after sending all values
}

func main() {
    // Create a channel of type string
    ch := make(chan string)

    // Start a goroutine to send data to the channel
    go sendData(ch)

    // Read values from the channel using a for loop
    for {
        // Attempt to receive a value from the channel
        value, ok := <-ch
        // Check if the channel is closed
        if !ok {
            fmt.Println("Channel closed")
            break
        }
        // Print the received value
        fmt.Println("Received:", value)
    }
}

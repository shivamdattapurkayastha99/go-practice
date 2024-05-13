package main

import "fmt"

func main() {
    // Declare a variable and initialize it with a value
    var num int = 10

    // Print the value and memory address of the variable
    fmt.Println("Value of num:", num)
    fmt.Println("Memory address of num:", &num)

    // Declare a pointer variable and assign the memory address of num to it
    var ptr *int = &num

    // Print the value of the pointer (memory address)
    fmt.Println("Value of ptr:", ptr)

    // Dereferencing the pointer to get the value it points to
    fmt.Println("Value at the memory address stored in ptr:", *ptr)

    // Update the value of num using the pointer
    *ptr = 20
    fmt.Println("Updated value of num:", num)
}

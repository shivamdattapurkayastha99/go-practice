package main

import "fmt"

// Define a struct type called Person
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Create a new instance of Person struct
    person1 := Person{
        Name: "John Doe",
        Age:  30,
        City: "New York",
    }

    // Access fields of the struct
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)
    fmt.Println("City:", person1.City)

    // Update a field
    person1.Age = 31
    fmt.Println("Updated Age:", person1.Age)

    // Initialize a struct without specifying field names
    person2 := Person{"Jane Smith", 25, "San Francisco"}
	fmt.Println("Updated Age:", person2.Age)

    // Access fields of the second struct instan
}
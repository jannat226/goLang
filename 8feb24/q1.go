package main

import "fmt"

type Product struct {
    Name     string
    Price    int
    Quantity int
}

func main() {
    var noOfProduct int
    fmt.Println("Enter the number of Products you want")
    fmt.Scanln(&noOfProduct)
	
    // Create a slice to store products
    products := make([]Product, noOfProduct)

    // taking input from user
    for i := 0; i < noOfProduct; i++ {
        var name string
        var price, quantity int
        fmt.Printf("Enter details for Product %d:\n", i+1)
        fmt.Print("Name: ")
        fmt.Scanln(&name)
        fmt.Print("Price: ")
        fmt.Scanln(&price)
        fmt.Print("Quantity: ")
        fmt.Scanln(&quantity)
        products[i] = Product{Name: name, Price: price, Quantity: quantity}
    }

    // Print out the products
    fmt.Println("\nProducts in inventory:")
    for i, product := range products {
        fmt.Printf("Product %d:\n", i+1)
        fmt.Printf("Name: %s\n", product.Name)
        fmt.Printf("Price: %d\n", product.Price)
        fmt.Printf("Quantity: %d\n", product.Quantity)
    }
}

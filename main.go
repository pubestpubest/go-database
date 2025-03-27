package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "pubest"
	password = "puebst"
	dbname   = "go101"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	sdb, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Fatal(err)
	}
	db = sdb
	err = sdb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Database successfully connected")
	addedProduct, err := createProduct(Product{0, "Go505", 500})
	if err != nil {
		log.Fatal("Error at adding stage")
		log.Fatal(err)
	}
	fmt.Println(addedProduct)
	p1, err := readProduct(1)
	if err != nil {
		log.Fatal("Error at reading stage")
		log.Fatal(err)
	}
	fmt.Print(p1)
}

func createProduct(p Product) (product Product, err error) {
	var newProduct Product
	errors := db.QueryRow(`
		INSERT INTO public.products(name, price) VALUES ($1, $2) RETURNING *;
		`, p.Name, p.Price).Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price)
	if errors != nil {
		return Product{}, errors
	}
	return newProduct, nil
}
func readProduct(id int) (product Product, err error) {
	errors := db.QueryRow(`
		SELECT * FROM public.products WHERE id=$1;
		`, id).Scan(&product.ID, &product.Name, &product.Price)
	if errors != nil {
		return Product{}, errors
	}
	return product, nil
}

func updateProduct(id int, p Product) (updated Product, err error) {

	return updated, nil
}

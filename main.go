package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
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

func setupDataabse() *sql.DB {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database successfully connected")
	return db
}

func main() {

	app := fiber.New()
	db = setupDataabse()
	defer db.Close()
	app.Listen(":3000")

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
	fmt.Println(p1)
	up, err := updateProduct(5, Product{0, "Go555", 555})
	if err != nil {
		log.Fatal("Error at updating stage")
		log.Fatal(err)
	}
	fmt.Println(up)
	del, err := deleteProduct(13)
	if err != nil {
		log.Fatal("Error at deleting stage")
		log.Fatal(err)
	}
	fmt.Println(del)
	arr, err := readProducts()
	if err != nil {
		log.Fatal("Error at read all stage")
		log.Fatal(err)
	}
	fmt.Println("Read all")
	fmt.Println(arr)
}

func createProduct(p Product) (product Product, err error) {
	var newProduct Product
	err = db.QueryRow(`
		INSERT INTO public.products(name, price) VALUES ($1, $2) RETURNING *;
		`, p.Name, p.Price).Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price)
	if err != nil {
		return Product{}, err
	}
	return newProduct, nil
}
func readProduct(id int) (product Product, err error) {
	err = db.QueryRow(`
		SELECT * FROM public.products WHERE id=$1;
		`, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func updateProduct(id int, p Product) (updated Product, err error) {
	err = db.QueryRow(`
		UPDATE public.products
		SET name=$1, price=$2
		WHERE id=$3 RETURNING *;
		`, p.Name, p.Price, id).Scan(&updated.ID, &updated.Name, &updated.Price)
	if err != nil {
		return Product{}, err
	}
	return updated, nil
}
func deleteProduct(id int) (deleted Product, err error) {
	err = db.QueryRow(`
		DELETE FROM public.products
		WHERE id=$1 RETURNING *;
		`, id).Scan(&deleted.ID, &deleted.Name, &deleted.Price)
	if err != nil {
		return Product{}, err
	}
	return deleted, nil
}
func readProducts() ([]Product, error) {
	rows, err := db.Query(`SELECT * FROM public.products;`)
	if err != nil {
		return nil, err
	}
	var products []Product
	defer rows.Close()
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, err
}

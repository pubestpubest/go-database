# Go PostgreSQL Database Connection and Raw SQL Queries

This project demonstrates how to connect to a PostgreSQL database using Go and perform basic CRUD (Create, Read, Update, Delete) operations using raw SQL queries.

## Prerequisites

- Go installed on your system
- PostgreSQL installed and running in Docker
- The `github.com/lib/pq` package

## Database Configuration

The project uses the following database configuration:
```go
const (
    host     = "localhost"
    port     = 5432
    username = "pubest"
    password = "puebst"
    dbname   = "go101"
)
```

## Database Schema

The project uses a simple `products` table with the following structure:
```sql
CREATE TABLE public.products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price INTEGER
);
```

## Features

The project implements the following database operations:

1. **Create Product**
   - Inserts a new product into the database
   - Returns the created product with its ID

2. **Read Product**
   - Retrieves a single product by ID
   - Returns product details or error if not found

3. **Update Product**
   - Updates an existing product's details
   - Returns the updated product

4. **Delete Product**
   - Removes a product from the database
   - Returns the deleted product details

5. **Read All Products**
   - Retrieves all products from the database
   - Returns a slice of products

## Code Structure

The main components of the project are:

- `main.go`: Contains all the database operations and connection logic
- `Product` struct: Represents the product data structure
- Database connection handling using `database/sql` package
- Raw SQL queries for CRUD operations

## Usage

1. Ensure your PostgreSQL server is running
2. Create the database and table using the schema provided above
3. Update the database configuration constants if needed
4. Run the program:
   ```bash
   go run main.go
   ```

## Example Operations

```go
// Create a new product
product := Product{Name: "Go505", Price: 500}
createdProduct, err := createProduct(product)

// Read a product
product, err := readProduct(1)

// Update a product
updatedProduct, err := updateProduct(5, Product{Name: "Go555", Price: 555})

// Delete a product
deletedProduct, err := deleteProduct(13)

// Read all products
products, err := readProducts()
```

## Error Handling

The project includes basic error handling for database operations. Each function returns both the result and an error, which should be checked before using the returned values.

## Dependencies

- `github.com/lib/pq`: PostgreSQL driver for Go's `database/sql` package

## Best Practices Demonstrated

1. Using parameterized queries to prevent SQL injection
2. Proper connection handling and closing
3. Error handling for database operations
4. Using transactions where appropriate
5. Proper resource cleanup with `defer`
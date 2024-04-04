package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	id       int
	name     string
	email    string
	password string
}

func (u *User) String() string {
	return fmt.Sprintf("id: %d; name: %s; email: %s; password: %s;",
		u.id, u.name, u.email, u.password)
}

type Product struct {
	id    int
	name  string
	price float64
}

func (p *Product) String() string {
	return fmt.Sprintf("id: %d; name: %s; price: %v;",
		p.id, p.name, p.price)
}

type Order struct {
	id           int
	user_id      int
	order_date   time.Time
	total_amount float64
}

func (o *Order) String() string {
	return fmt.Sprintf(
		"id: %d; user id: %d; order date: %v; total amount: %v;",
		o.id, o.user_id, o.order_date.Format(time.DateOnly), o.total_amount)
}

func main() {
	dsn := "postgres://postgres:admin@localhost/db?sslmode=disable"
	ctx := context.Background()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")

	// Insert users
	newUsers := make([]User, 5)
	newUsers[0] = User{id: 1, name: "John Doe", email: "john@example.com", password: "qwerty1"}
	newUsers[1] = User{id: 2, name: "Bob", email: "bob@example.com", password: "12345"}
	newUsers[2] = User{id: 3, name: "Alice", email: "alice@example.com", password: "qwerty"}
	newUsers[3] = User{id: 4, name: "Tom", email: "tom@example.com", password: "123456"}
	newUsers[4] = User{id: 5, name: "Jerry", email: "jerry@example.com", password: "654321"}

	for _, user := range newUsers {
		err = InsertUser(ctx, db, user)
		if err != nil {
			log.Printf("Error when adding a user: %v", err)
		} else {
			log.Printf("User %s has been successfully added", user.name)
		}
	}

	// Edit user Alice
	updateUser := User{id: 3, name: "Alice Smith", email: "alice@example.com", password: "gnd3rtw35"}
	err = EditUser(ctx, db, updateUser)
	if err != nil {
		log.Printf("Error when updating a user: %v", err)
	} else {
		log.Printf("User %s has been successfully updated", updateUser.name)
	}

	// Delete user Bob
	userName := "Bob"
	err = DeleteUser(ctx, db, userName)
	if err != nil {
		log.Printf("Error when deleting a user: %v", err)
	} else {
		log.Printf("User %s has been successfully deleted", userName)
	}

	// Insert products
	newPoducts := make([]Product, 7)
	newPoducts[0] = Product{id: 1, name: "milk", price: 5}
	newPoducts[1] = Product{id: 2, name: "bread", price: 2.13}
	newPoducts[2] = Product{id: 3, name: "tomato", price: 0.5}
	newPoducts[3] = Product{id: 4, name: "apple", price: 0.7}
	newPoducts[4] = Product{id: 5, name: "cake", price: 4.1}
	newPoducts[5] = Product{id: 6, name: "banana", price: 0.3}
	newPoducts[6] = Product{id: 7, name: "donut", price: 2}

	for _, product := range newPoducts {
		err = InsertProduct(ctx, db, product)
		if err != nil {
			log.Printf("Error when adding a product: %v", err)
		} else {
			log.Printf("Product %s has been successfully added", product.name)
		}
	}
	// Edit product
	updateProduct := Product{id: 1, name: "milk", price: 6.5}
	err = EditProduct(ctx, db, updateProduct)
	if err != nil {
		log.Printf("Error when updated a product: %v", err)
	} else {
		log.Printf("Product %s has been successfully updated", updateProduct.name)
	}

	// Delete product
	productName := "tomato"
	err = DeleteProduct(ctx, db, productName)
	if err != nil {
		log.Printf("Error when deleting a user: %v", err)
	} else {
		log.Printf("Product %s has been successfully deleted", productName)
	}

	// Add orders
	newOrders := make([]Order, 6)
	newOrders[0] = Order{id: 1, user_id: 5, order_date: time.Now(), total_amount: 10}
	newOrders[1] = Order{id: 2, user_id: 4, order_date: time.Now(), total_amount: 21}
	newOrders[2] = Order{id: 3, user_id: 5, order_date: time.Now(), total_amount: 2}
	newOrders[3] = Order{id: 4, user_id: 1, order_date: time.Now(), total_amount: 7}
	newOrders[4] = Order{id: 5, user_id: 3, order_date: time.Now(), total_amount: 3.5}
	newOrders[5] = Order{id: 6, user_id: 5, order_date: time.Now(), total_amount: 1}

	for _, order := range newOrders {
		err = InsertOrder(ctx, db, order)
		if err != nil {
			log.Printf("Error when adding a order: %v", err)
		} else {
			log.Printf("Order %d has been successfully added", order.id)
		}
	}

	// Delete order
	orderID := 3
	err = DeleteOrder(ctx, db, orderID)
	if err != nil {
		log.Printf("Error when deleting a order: %v", err)
	} else {
		log.Printf("Order %d has been successfully deleted", orderID)
	}

	// Select Users

	users, err := SelectUsers(ctx, db)
	for _, result := range users {
		if err != nil {
			log.Printf("Error when selecting a users: %v", err)
		} else {
			log.Println(result.String())
		}
	}

	// Select Products

	products, err := SelectProducts(ctx, db)
	for _, result := range products {
		if err != nil {
			log.Printf("Error when selecting a products: %v", err)
		} else {
			log.Println(result.String())
		}
	}

	// Select orders by user

	userID := 5
	orders, err := SelectOrders(ctx, db, userID)
	for _, order := range orders {
		if err != nil {
			log.Printf("Error when selecting a products: %v", err)
		} else {
			log.Println(order.String())
		}
	}
}

func InsertUser(ctx context.Context, db *sql.DB, user User) error {
	q := `insert into Users(id, name, email, password) values($1, $2, $3, $4)`
	_, err := db.ExecContext(ctx, q,
		user.id, user.name, user.email, user.password)
	if err != nil {
		return err
	}
	return nil
}

func InsertProduct(ctx context.Context, db *sql.DB, product Product) error {
	q := `insert into Products(id, name, price) values($1, $2, $3)`
	_, err := db.ExecContext(ctx, q,
		product.id, product.name, product.price)
	if err != nil {
		return err
	}
	return nil
}

func EditUser(ctx context.Context, db *sql.DB, user User) error {
	q := `update Users
	set name=$2, email=$3, password=$4
	where id=$1;`
	_, err := db.ExecContext(ctx, q,
		user.id, user.name, user.email, user.password)
	if err != nil {
		return err
	}
	return nil
}

func EditProduct(ctx context.Context, db *sql.DB, product Product) error {
	q := `update Products
	set name=$2, price=$3
	where id=$1;`
	_, err := db.ExecContext(ctx, q,
		product.id, product.name, product.price)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, db *sql.DB, userName string) error {
	q := `delete from Users
	where name=$1;`
	_, err := db.ExecContext(ctx, q, userName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(ctx context.Context, db *sql.DB, productName string) error {
	q := `delete from Products
	where name=$1;`
	_, err := db.ExecContext(ctx, q, productName)
	if err != nil {
		return err
	}
	return nil
}

func InsertOrder(ctx context.Context, db *sql.DB, order Order) error {
	q := `insert into Orders(id, user_id, order_date, total_amount) values($1, $2, $3, $4)`
	_, err := db.ExecContext(ctx, q,
		order.id, order.user_id, order.order_date, order.total_amount)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(ctx context.Context, db *sql.DB, orderID int) error {
	q := `delete from Orders
	where id=$1;`
	_, err := db.ExecContext(ctx, q, orderID)
	if err != nil {
		return err
	}
	return nil
}

func SelectUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	var users []User
	q := `select * from Users`

	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var user User
		if err := rows.Scan(&user.id, &user.name, &user.email, &user.password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func SelectProducts(ctx context.Context, db *sql.DB) ([]Product, error) {
	var products []Product
	q := `select * from Products`

	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.id, &product.name, &product.price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func SelectOrders(ctx context.Context, db *sql.DB, userID int) ([]Order, error) {
	var orders []Order
	q := `select * from Orders
	where user_id=$1`

	rows, err := db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.id, &order.user_id, &order.order_date, &order.total_amount); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

// rowsAffected, err := result.RowsAffected() // int64
// if err != nil {
// 	return err
// }

// log.Printf("The number of modified rows: %v", rowsAffected)

// func Edit(ctx context.Context, dsn string) error {

// }

// func Delete(ctx context.Context, dsn string) error {

// }

// func Select(ctx context.Context, dsn string) error {
// 	query := `
// 		select id, title, descr
// 		from events
// 		where owner = $1 and start_date = $2
// 	`
// 	rows, err := db.QueryContext(ctx, query, owner, date)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int64
// 		var title, descr string
// 		if err := rows.Scan(&id, &title, &descr); err != nil {
// 			// ошибка сканирования
// 		}
// 		// обрабатываем строку
// 		fmt.Printf("%d %s %s\n", id, title, descr)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Println(err)
// 	}
// }

// This function will make a connection to the database only once.
// func init() {
// 	var err error

// 	connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
// 	db, err = sql.Open("postgres", connStr)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	}
// 	// this will be printed in the terminal, confirming the connection to the database
// 	fmt.Println("The database is connected")
// }

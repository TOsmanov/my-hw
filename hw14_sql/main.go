package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

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
	id          int
	userID      int
	orderDate   time.Time
	totalAmount float64
}

type NewOrder struct {
	id         int
	userID     int
	orederDate time.Time
	productID  int
	amount     int
}

func (o *Order) String() string {
	return fmt.Sprintf(
		"id: %d; user id: %d; order date: %v; total amount: %v;",
		o.id, o.userID, o.orderDate.Format(time.DateOnly), o.totalAmount)
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
	dbPreparation(ctx, db)
	gettingData(ctx, db)
}

func dbPreparation(ctx context.Context, db *sql.DB) {
	var err error
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
	productID := 3
	err = DeleteProduct(ctx, db, productID)
	if err != nil {
		log.Printf("Error when deleting a user: %v", err)
	} else {
		log.Printf("Product with ID %d has been successfully deleted", productID)
	}

	// Add orders
	newOrders := make([]NewOrder, 9)
	newOrders[0] = NewOrder{id: 1, userID: 5, orederDate: time.Now(), productID: 7, amount: 2}
	newOrders[1] = NewOrder{id: 2, userID: 4, orederDate: time.Now(), productID: 7, amount: 5}
	newOrders[2] = NewOrder{id: 3, userID: 5, orederDate: time.Now(), productID: 4, amount: 1}
	newOrders[3] = NewOrder{id: 4, userID: 1, orederDate: time.Now(), productID: 2, amount: 6}
	newOrders[4] = NewOrder{id: 5, userID: 3, orederDate: time.Now(), productID: 4, amount: 3}
	newOrders[5] = NewOrder{id: 6, userID: 5, orederDate: time.Now(), productID: 6, amount: 1}
	newOrders[6] = NewOrder{id: 7, userID: 5, orederDate: time.Now(), productID: 1, amount: 6}
	newOrders[7] = NewOrder{id: 8, userID: 5, orederDate: time.Now(), productID: 4, amount: 8}
	newOrders[8] = NewOrder{id: 9, userID: 5, orederDate: time.Now(), productID: 6, amount: 10}

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
}

func gettingData(ctx context.Context, db *sql.DB) {
	var err error
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

	// Select user statistics
	userIDs := []int{5, 1, 4}

	for _, userID := range userIDs {
		var total, avg float64
		total, avg, err = SelectUserStat(ctx, db, userID)
		if err != nil {
			log.Printf("Error when getting user statistics: %v", err)
		} else {
			log.Printf("User statistics with ID: %d; Total cost of orders: %.2f; Average product price: %.2f",
				userID, total, avg)
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

func DeleteProduct(ctx context.Context, db *sql.DB, productID int) error {
	q := `delete from OrderProducts
	where product_id=$1;`
	_, err := db.ExecContext(ctx, q, productID)
	if err != nil {
		return err
	}

	q = `delete from Products
	where name=$1;`
	_, err = db.ExecContext(ctx, q, productID)
	if err != nil {
		return err
	}
	return nil
}

func InsertOrder(ctx context.Context, db *sql.DB, order NewOrder) error {
	price, err := GetProductPrice(ctx, db, order.productID)
	if err != nil {
		return err
	}

	q := `insert into Orders (id, user_id, order_date, total_amount)
	values($1, $2, $3, $4);`
	_, err = db.ExecContext(ctx, q,
		order.id, order.userID, order.orederDate, fmt.Sprintf("%.2f", (float64(order.amount)*price)))
	if err != nil {
		return err
	}
	q = `insert into OrderProducts (order_id, product_id, amount)
	values($1, $2, $3)`
	_, err = db.ExecContext(ctx, q,
		order.id, order.productID, order.amount)
	if err != nil {
		return err
	}
	return nil
}

func GetProductPrice(ctx context.Context, db *sql.DB, productID int) (float64, error) {
	var price float64
	q := `select price from Products
	where id=$1`

	rows, err := db.QueryContext(ctx, q, productID)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&price); err != nil {
			return -1, err
		}
	}
	return price, err
}

func DeleteOrder(ctx context.Context, db *sql.DB, orderID int) error {
	q := `delete from OrderProducts
	where order_id=$1;`
	_, err := db.ExecContext(ctx, q, orderID)
	if err != nil {
		return err
	}

	q = `delete from Orders
	where id=$1;`
	_, err = db.ExecContext(ctx, q, orderID)
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
		if err = rows.Scan(&user.id, &user.name, &user.email, &user.password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
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
		if err = rows.Scan(&product.id, &product.name, &product.price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
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
		if err = rows.Scan(&order.id, &order.userID, &order.orderDate, &order.totalAmount); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func SelectUserStat(ctx context.Context, db *sql.DB, userID int) (float64, float64, error) {
	q := `select sum(Orders.total_amount) as total_order_amount,
    avg(Products.price) as avg_product_price
	from Users
	join Orders on $1 = Orders.user_id
	join OrderProducts on Orders.id = OrderProducts.order_id
	join Products on OrderProducts.product_id = Products.id
	group by $1`

	rows, err := db.QueryContext(ctx, q, userID)
	if err != nil {
		return -1, -1, err
	}
	defer rows.Close()
	var sum, avg float64
	for rows.Next() {
		if err = rows.Scan(&sum, &avg); err != nil {
			return -1, -1, err
		}
	}
	if err = rows.Err(); err != nil {
		return -1, -1, err
	}
	return sum, avg, err
}

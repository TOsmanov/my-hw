CREATE TABLE Users (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    password VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX idx_email ON Users (email);

CREATE TABLE Orders (
    id serial PRIMARY KEY,
    user_id INT,
    order_date DATE,
    total_amount FLOAT NOT NULL,
    FOREIGN KEY (user_id)  REFERENCES Users (id)
);

CREATE TABLE Products (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL
);

CREATE TABLE OrderProducts (
    products_id INT,
    orders_id INT,
    CONSTRAINT fk_productsid FOREIGN KEY (products_id) REFERENCES Products (id),
    CONSTRAINT fk_ordersid FOREIGN KEY (orders_id) REFERENCES Orders (id)
);

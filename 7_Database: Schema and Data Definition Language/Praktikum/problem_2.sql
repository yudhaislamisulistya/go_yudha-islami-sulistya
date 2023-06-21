drop database alta_online_shop;

CREATE DATABASE alta_online_shop;

use alta_online_shop;

CREATE TABLE user (
    user_id INT PRIMARY KEY,
    name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);


CREATE TABLE product_type (
    product_type_id INT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE operators (
    operator_id INT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE product_description (
    product_description_id INT PRIMARY KEY,
    description VARCHAR(255)
);

CREATE TABLE payment_method (
    payment_method_id INT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE product (
    product_id INT PRIMARY KEY,
    name VARCHAR(255),
    product_type_id INT,
    operator_id INT,
    product_description_id INT,
    payment_method_id INT,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (product_type_id) REFERENCES product_type(product_type_id),
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id),
    FOREIGN KEY (product_description_id) REFERENCES product_description(product_description_id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(payment_method_id)
);

CREATE TABLE transaction (
    transaction_id INT PRIMARY KEY,
    user_id INT,
    transaction_date DATETIME,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE transaction_detail (
    transaction_detail_id INT PRIMARY KEY,
    transaction_id INT,
    product_id INT,
    quantity INT,
    total_amount DECIMAL(10, 2),
    FOREIGN KEY (transaction_id) REFERENCES transaction(transaction_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE kurir (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

ALTER TABLE kurir ADD ongkos_dasar DECIMAL(10, 2);

ALTER TABLE kurir RENAME TO shipping;

DROP TABLE shipping;

CREATE TABLE payment_method_description (
    description_id INT PRIMARY KEY,
    description VARCHAR(255)
);

ALTER TABLE payment_method
ADD COLUMN description_id INT,
ADD CONSTRAINT fk_payment_method_description
FOREIGN KEY (description_id) REFERENCES payment_method_description(description_id);

CREATE TABLE address (
    address_id INT PRIMARY KEY,
    user_id INT,
    address VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);


CREATE TABLE user_payment_method_detail (
    user_id INT,
    payment_method_id INT,
    PRIMARY KEY (user_id, payment_method_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(payment_method_id)
);


CREATE TABLE orders (
                        id VARCHAR(255) PRIMARY KEY,
                        customer_id VARCHAR(255) NOT NULL,
                        cashier_id VARCHAR(255) NOT NULL,
                        paid INTEGER NOT NULL,
                        change INTEGER,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        FOREIGN KEY (customer_id) REFERENCES users(id),
                        FOREIGN KEY (cashier_id) REFERENCES users(id)
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        customer_id int NOT NULL,
                        cashier_id int NOT NULL,
                        paid INTEGER NOT NULL,
                        change INTEGER,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE CASCADE,
                        FOREIGN KEY (cashier_id) REFERENCES users(id) ON DELETE CASCADE
);

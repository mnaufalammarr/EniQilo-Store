CREATE TABLE product_details (
                                 product_id int NOT NULL,
                                 checkout_id int NOT NULL,
                                 quantity INTEGER NOT NULL,
                                 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
                                 FOREIGN KEY (checkout_id) REFERENCES orders(id) ON DELETE CASCADE
);

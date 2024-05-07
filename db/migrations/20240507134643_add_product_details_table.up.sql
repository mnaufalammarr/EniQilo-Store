CREATE TABLE product_details (
                                 product_id VARCHAR(255) NOT NULL,
                                 checkout_id VARCHAR(255) NOT NULL,
                                 quantity INTEGER NOT NULL,
                                 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 FOREIGN KEY (product_id) REFERENCES products(id),
                                 FOREIGN KEY (checkout_id) REFERENCES orders(id)
);

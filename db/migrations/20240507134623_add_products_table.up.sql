CREATE TABLE products (
                          id VARCHAR(255) PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          sku VARCHAR(255) NOT NULL,
                          category VARCHAR(255),
                          image_url VARCHAR(255),
                          note VARCHAR(255),
                          price INTEGER NOT NULL,
                          stock INTEGER NOT NULL,
                          location VARCHAR(255),
                          is_available BOOLEAN NOT NULL DEFAULT FALSE,
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

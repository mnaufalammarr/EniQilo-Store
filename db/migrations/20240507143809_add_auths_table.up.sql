CREATE TABLE auths
(
    id         VARCHAR(255) PRIMARY KEY,
    password   VARCHAR(255) NOT NULL,
    user_id    VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
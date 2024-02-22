CREATE TABLE wallets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    currency VARCHAR(50) NOT NULL,
    balance DECIMAL(18, 2) DEFAULT 0.0,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

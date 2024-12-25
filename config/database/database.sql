CREATE DATABASE backend_started;
USE backend_started;

CREATE TABLE users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role ENUM('buyer', 'seller') NOT NULL,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE products (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          description TEXT,
                          price DECIMAL(10, 2) NOT NULL,
                          seller_id INT NOT NULL,
                          status ENUM('available', 'sold') DEFAULT 'available',
                          created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                          updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          FOREIGN KEY (seller_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE transactions (
                              id INT AUTO_INCREMENT PRIMARY KEY,
                              product_id INT NOT NULL,
                              buyer_id INT NOT NULL,
                              quantity INT NOT NULL,
                              total_price DECIMAL(10, 2) NOT NULL,
                              created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
                              FOREIGN KEY (buyer_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE payments (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          transaction_id INT NOT NULL,
                          payment_status ENUM('paid', 'pending') DEFAULT 'pending',
                          payment_date DATETIME,
                          FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE
);

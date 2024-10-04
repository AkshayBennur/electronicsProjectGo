--products
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(100) NOT NULL,
    brand VARCHAR(100) NOT NULL,
    rating INT DEFAULT 0,
    selected BOOLEAN DEFAULT FALSE,
    ordered BOOLEAN DEFAULT FALSE
);

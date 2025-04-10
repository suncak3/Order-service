-- Drop tables if they exist
DROP TABLE IF EXISTS orders;


-- Create Order table
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        product_id INTEGER NOT NULL,
                        quantities INTEGER NOT NULL CHECK (quantities > 0)
);

-- Insert sample orders
INSERT INTO orders (product_id, quantities) VALUES
                                                (1, 2),  -- 2 Smartphones
                                                (2, 5),  -- 5 Books
                                                (3, 1);  -- 1 T-shirt
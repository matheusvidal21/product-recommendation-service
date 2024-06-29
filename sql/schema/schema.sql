CREATE TABLE categories (
                            id UUID PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            description TEXT
);

CREATE TABLE products (
                          id UUID PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          price NUMERIC(10, 2) NOT NULL,
                          category_id UUID REFERENCES categories(id)
);

CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL
);

CREATE TABLE user_activities (
                                 user_id UUID REFERENCES users(id) NOT NULL,
                                 product_id UUID REFERENCES products(id) NOT NULL,
                                 action VARCHAR(50) NOT NULL,
                                 PRIMARY KEY (user_id, product_id)
);

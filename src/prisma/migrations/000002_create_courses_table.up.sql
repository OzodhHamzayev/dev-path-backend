CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    category TEXT NOT NULL,
    level TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    rating NUMERIC(2, 1) CHECK (rating >= 0 AND rating <= 5),
    duration INT,
    instructor VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS todos (
   id VARCHAR(40) PRIMARY KEY,
   title VARCHAR(50) NOT NULL,
   description VARCHAR(255),
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  provider VARCHAR(100),
  provider_id VARCHAR(255) UNIQUE,

  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,

  created_at DATETIME NOT NULL DEFAULT (NOW()),
  updated_at DATETIME NOT NULL DEFAULT (NOW())
);

INSERT INTO users (provider, name, email) VALUES
("google", "Super IT KB23", "itkampungbudaya@gmail.com")

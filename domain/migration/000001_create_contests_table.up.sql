CREATE TABLE IF NOT EXISTS contests (
  id INT PRIMARY KEY AUTO_INCREMENT,

  name VARCHAR(255) NOT NULL,

  created_at DATETIME NOT NULL DEFAULT (NOW()),
  updated_at DATETIME NOT NULL DEFAULT (NOW())
);
INSERT INTO contests (name) VALUES
('Musik'),
('Tari'),
('Busana Kreasi Umum'),
('Busana Kreasi Forda'),
('Stand Bazzar Forda');

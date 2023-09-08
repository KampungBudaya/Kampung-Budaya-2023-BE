CREATE TABLE IF NOT EXISTS participants(
  id INT PRIMARY KEY AUTO_INCREMENT,
  contests_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  birth VARCHAR(255) NOT NULL,
  status ENUM('PENDING', 'ACCEPTED', 'REJECTED') DEFAULT ('PENDING') NOT NULL,
  category ENUM('FORDA', 'UMUM') NOT NULL,
  institution VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  instagram VARCHAR(255) NOT NULL,
  line VARCHAR(255) NOT NULL,
  phone_number VARCHAR(13) NOT NULL,
  form VARCHAR(255) NOT NULL,
  video_url VARCHAR(255),
  payment_proof VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT (NOW()),
  updated_at DATETIME NOT NULL DEFAULT (NOW())
);

ALTER TABLE participants ADD CONSTRAINT participants_fk_contests_id FOREIGN KEY (contests_id) REFERENCES contests (id)

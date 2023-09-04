CREATE TABLE IF NOT EXISTS participants(
  id INT PRIMARY KEY AUTO_INCREMENT,
  contests_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  is_verified TINYINT(1) NOT NULL DEFAULT 0,
  origin VARCHAR(255) NOT NULL,
  form_url VARCHAR(255) NOT NULL,
  video_url VARCHAR(255),
  payment_proof VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT (NOW()),
  updated_at DATETIME NOT NULL DEFAULT (NOW())
);

ALTER TABLE participants ADD CONSTRAINT participants_fk_contests_id FOREIGN KEY (contests_id) REFERENCES contests (id)

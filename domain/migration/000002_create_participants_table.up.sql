CREATE TABLE IF NOT EXISTS participants(
  id INT PRIMARY KEY AUTO_INCREMENT,
  contest_id INT NOT NULL,

  name VARCHAR(255) NOT NULL,
  is_verified TINYINT(1) DEFAULT 0,
  origin VARCHAR(255) NOT NULL,
  phone_number VARCHAR(13) NOT NULL,
  form_url VARCHAR(255) NOT NULL,
  video_url VARCHAR(255),
  payment_proof VARCHAR(255) NOT NULL,

  created_at DATETIME NOT NULL DEFAULT (NOW()),
  updated_at DATETIME NOT NULL DEFAULT (NOW())
);

ALTER TABLE participants ADD CONSTRAINT participants_fk_contest_id FOREIGN KEY (contest_id) REFERENCES contests (id)

CREATE TABLE IF NOT EXISTS user_has_roles (
  user_id INT NOT NULL,
  role_id INT NOT NULL,

  created_at DATETIME NOT NULL DEFAULT (NOW()),
  updated_at DATETIME NOT NULL DEFAULT (NOW())
);

ALTER TABLE user_has_roles ADD CONSTRAINT user_has_roles_fk_user_id FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_has_roles ADD CONSTRAINT user_has_roles_fk_role_id FOREIGN KEY (role_id) REFERENCES roles (id);

INSERT INTO user_has_roles (user_id, role_id) VALUES
(
  (SELECT id FROM users WHERE email = "ITKampungBudaya@gmail.com" LIMIT 1),
  (SELECT id FROM roles WHERE name = "Super Admin" LIMIT 1)
);

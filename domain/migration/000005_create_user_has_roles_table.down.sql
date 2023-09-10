ALTER TABLE user_has_roles DROP CONSTRAINT user_has_roles_fk_user_id;
ALTER TABLE user_has_roles DROP CONSTRAINT user_has_roles_fk_role_id;

DROP TABLE IF EXISTS user_has_roles;

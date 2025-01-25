ALTER TABLE login_logs
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);

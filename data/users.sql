DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id            INT AUTO_INCREMENT NOT NULL,
  username      VARCHAR(255) NOT NULL,
  phone         VARCHAR(255) NOT NULL,
  dateOfBirth   DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users
  (username, phone, dateOfBirth)
VALUES
  ('Thien', '0365888722', '1998-07-04'),
  ('Tra', '0123456789', '1998-02-03'),
  ('Jeru', '1203564879', '1999-05-02'),
  ('Sarah Vaughan', '021569853', '2000-03-04');
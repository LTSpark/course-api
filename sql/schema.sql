CREATE TABLE courses
(
    id       VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;

CREATE TABLE users
(
  id        VARCHAR(255) NOT NULL,
  name      VARCHAR(255) NOT NULL,
  email     VARCHAR(255) NOT NULL,
  password  VARCHAR(255) NOT NULL
)
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR(64) NOT NULL UNIQUE,
  password VARCHAR(128) NOT NULL
);

INSERT INTO users (id, username, password) VALUES (1, 'admin', 'kCrYsDFKcy+HecmW+OeLOjiUCrjUW7+yg3exkuHdDa/1HZTvMUU57pAjGHk1vE7a0raPbi26+8muc2/BDLi93w==');
INSERT INTO users (id, username, password) VALUES (2, 'user', 'zYI4i7fymHZa1S3393Mes266aixu1rJ4K/xu2oUa95oek9X8S1AGdM/VYAvUlU2Pw0497lN42c9f2EvrbV6WWw==');
INSERT INTO users (id, username, password) VALUES (3, 'test', 'test');

CREATE TABLE stories (
  id BIGSERIAL PRIMARY KEY,
  type SMALLINT NOT NULL,
  link VARCHAR(256) NOT NULL,
  text VARCHAR(128) NOT NULL,
  text_pos_x INT NOT NULL,
  text_pos_y INT NOT NULL
);

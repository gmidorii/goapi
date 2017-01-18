# create table
CREATE TABLE t_user (
  id int PRIMARY KEY AUTO_INCREMENT,
  name varchar(20),
  color varchar(10),
  insert_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_time datetime
);

# update column
ALTER TABLE t_user CHANGE
insert_time
insert_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE t_user CHANGE
id
id int NOT NULL AUTO_INCREMENT PRIMARY KEY;

# insert
INSERT INTO t_user (
  name, color, update_time
) VALUES (
  'so', 'green', now()
);

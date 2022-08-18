CREATE TABLE IF NOT EXISTS orders (
  id bigint PRIMARY KEY,
  name varchar(255) NOT NULL,
  description varchar(255),
  total_price bigint NOT NULL,
  quantity int NOT NULL,
  discount int,
  shipping varchar(255),
  user_id int NOT NULL,
  date_purchased date,
  status varchar(50),
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS users (
  id bigint PRIMARY KEY,
  name varchar(255) NOT NULL,
  username varchar(50) UNIQUE NOT NULL,
  password varchar(50) NOT NULL,
  phone_number varchar(20),
  address varchar(255),
  age smallint,
  role varchar(10),
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS audit_trail (
  id bigint PRIMARY KEY,
  user_id int NOT NULL,
  order_id int NOT NULL,
  operation varchar(100),
  created_at timestamp,
  updated_at timestamp
);


ALTER TABLE orders ADD CONSTRAINT orders_user_id FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE audit_trail ADD CONSTRAINT audit_trail_user_id FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE audit_trail ADD CONSTRAINT audit_trail_order_id FOREIGN KEY (order_id) REFERENCES orders (id);
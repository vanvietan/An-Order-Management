
CREATE TABLE "order" (
  "id" int PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "description" varchar(255),
  "total_price" bigint NOT NULL,
  "quantity" int NOT NULL,
  "discount" smallint(11),
  "shipping" varchar(255),
  "user_id" int NOT NULL,
  "date_purchased" date,
  "status" varchar(50),
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "username" varchar(50) UNIQUE NOT NULL,
  "password" varchar(50) NOT NULL,
  "phone_number" varchar(20),
  "address" varchar(255),
  "age" smallint,
  "role" varchar(10),
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "history" (
  "id" int PRIMARY KEY,
  "user_id" int NOT NULL,
  "order_id" int NOT NULL,
  "operation" varchar(100),
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "order" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "history" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "history" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

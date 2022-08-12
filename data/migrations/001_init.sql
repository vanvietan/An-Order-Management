DROP TABLE IF EXISTS order;
CREATE TABLE "order" (
                         "id" SERIAL PRIMARY KEY,
                         "name" varchar(255) NOT NULL,
                         "description" varchar(255),
                         "total_price" float NOT NULL,
                         "quantity" int NOT NULL,
                         "discount" smallint(11),
                         "shipping" varchar(255),
                         "payment_method" varchar(255),
                         "order_status" int(5),
                         "user_id" int NOT NULL,
                         "date_purchased" datetime,
                         "created_at" timestamp,
                         "updated_at" timestamp
);

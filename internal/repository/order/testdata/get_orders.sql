TRUNCATE table users CASCADE;
INSERT INTO public.users
(id, "name", username, "password", phone_number, address, age, "role", created_at, updated_at)
VALUES
    (99, 'nghia', 'nghia', 'nghia', '123', 'nghia', 1, 'USER', '2022-03-14 14:00:00', '2022-03-14 14:00:00'),
    (100, 'abc', 'abc', 'abc', '123', 'abc', 1, 'ADMIN', '2022-03-15 15:00:00', '2022-03-15 15:00:00'),
    (101, 'abc', 'abc1', 'abc', '123', 'abc', 1, 'ADMIN', '2022-03-15 16:00:00', '2022-03-15 16:00:00');

TRUNCATE table orders CASCADE;
INSERT INTO public.orders
(id, "name", description, total_price, quantity, discount, shipping, status, user_id, date_purchased, created_at, updated_at)
VALUES
(99, 'an', 'an', 99, 99, 9, 'hcm', 'SHIPPED', 99, '2022-03-14', '2022-03-14 14:00:00', '2022-03-14 14:00:00'),
(100, 'abc', 'abc', 100, 100, 10, 'abc', 'APPROVED', 100, '2022-03-15', '2022-03-15 15:00:00', '2022-03-15 15:00:00'),
(101, 'abc1', 'abc1', 101, 101, 11, 'abc', 'APPROVED', 101, '2022-03-16', '2022-03-16 16:00:00', '2022-03-16 16:00:00');

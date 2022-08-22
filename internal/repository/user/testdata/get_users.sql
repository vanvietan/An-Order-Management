TRUNCATE table users CASCADE;
INSERT INTO public.users
(id, "name", username, "password", phone_number, address, age, "role", created_at, updated_at)
VALUES
(100, 'abc', 'abc', 'abc', '123', 'abc', 1, 'ADMIN', '2022-03-15 15:00:00', '2022-03-15 15:00:00'),
(101, 'abc', 'abc1', 'abc', '123', 'abc', 1, 'ADMIN', '2022-03-15 16:00:00', '2022-03-15 16:00:00');

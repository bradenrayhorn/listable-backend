CREATE TABLE IF NOT EXISTS users
(
    `id`         int unsigned auto_increment primary key,
    `name`       varchar(255) unique not null,
    `password`   char(60)            not null,
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp
)

CREATE TABLE IF NOT EXISTS `groups`
(
    `id`         int unsigned auto_increment primary key,
    `name`       varchar(255) unique not null,
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp
)

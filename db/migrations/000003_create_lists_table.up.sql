CREATE TABLE IF NOT EXISTS lists
(
    `id`         int unsigned auto_increment primary key,
    `group_id`   int unsigned,
    `name`       varchar(255) not null,
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp
)

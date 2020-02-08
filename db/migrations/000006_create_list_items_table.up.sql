CREATE TABLE IF NOT EXISTS list_items
(
    `id`         int unsigned auto_increment primary key,
    `list_id`    int unsigned            not null,
    `content`    varchar(255)            not null,
    `checked`    boolean   default false not null,
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp
)

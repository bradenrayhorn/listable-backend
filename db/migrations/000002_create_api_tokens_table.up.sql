CREATE TABLE IF NOT EXISTS api_tokens
(
    `user_id`    int unsigned,
    `token`      char(64) not null primary key,
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp
)

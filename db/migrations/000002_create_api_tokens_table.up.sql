CREATE TABLE IF NOT EXISTS api_tokens
(
    `user_id`    int unsigned,
    `token`      char(64) not null,
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp,
    PRIMARY KEY (user_id, token)
)

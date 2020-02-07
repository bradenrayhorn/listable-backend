CREATE TABLE IF NOT EXISTS groups_users
(
    `user_id`    int unsigned,
    `group_id`   int unsigned,
    `created_at` timestamp default current_timestamp,
    `updated_at` timestamp default current_timestamp,
    PRIMARY KEY (`user_id`, `group_id`)
)

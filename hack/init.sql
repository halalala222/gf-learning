CREATE TABLE IF NOT EXISTS user
(
    id         int auto_increment primary key,
    address    char(42) null,
    created_at datetime null,
    updated_at datetime null,
    deleted_at datetime null
)
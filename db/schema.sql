CREATE TABLE user (
    `id` BIGINT UNSIGNED NOT NULL,

    `name` VARBINARY(32) NOT NULL,
    `password_hash` VARBINARY(254) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY (name),

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE user_session (
    `user_id` BIGINT UNSIGNED NOT NULL,
    `token` VARBINARY(512) NOT NULL,

    `expires_at` DATETIME(6) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE diary (
    `id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,

    `name` VARBINARY(32) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    KEY (user_id),
    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE article (
    `id` BIGINT UNSIGNED NOT NULL,
    `diary_id` BIGINT UNSIGNED NOT NULL,

    `title` VARBINARY(100) NOT NULL,
    `body` MEDIUMBLOB NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    KEY (diary_id),
    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

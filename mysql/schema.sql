CREATE TABLE `user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name` varchar(20) NOT NULL COMMENT 'ユーザー名',
    `password` VARCHAR(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role` VARCHAR(80) NOT NULL COMMENT 'ロール',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_bin COMMENT = 'ユーザー';
CREATE TABLE `post` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '記事の識別子',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '記事を作成したユーザーの識別子',
    `title` VARCHAR(128) NOT NULL COMMENT '記事のタイトル',
    `body` VARCHAR(128) NOT NULL COMMENT '記事の本文',
    `status` VARCHAR(20) NOT NULL COMMENT '記事の状態',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_bin COMMENT = '記事';

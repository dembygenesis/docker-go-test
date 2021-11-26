CREATE DATABASE abc;
CREATE DATABASE abc1;
CREATE TABLE `users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`)
VALUES ('Solomon'),
       ('Menelik');
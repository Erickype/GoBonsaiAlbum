create database if not exists test;
create table test (id int primary key, name text);

CREATE SCHEMA `users`;

CREATE TABLE `users`.`user` (
                                `id` int PRIMARY KEY AUTO_INCREMENT,
                                `userName` varchar(30),
                                `userLastname` varchar(30),
                                `userNickname` varchar(15),
                                `createdAt` datetime
);

CREATE TABLE `users`.`friendship` (
                                      `firstUser` int,
                                      `secondUser` int,
                                      `createdAt` datetime,
                                      PRIMARY KEY (`firstUser`, `secondUser`)
);

ALTER TABLE `users`.`friendship` ADD FOREIGN KEY (`firstUser`) REFERENCES `users`.`user` (`id`);

ALTER TABLE `users`.`friendship` ADD FOREIGN KEY (`secondUser`) REFERENCES `users`.`user` (`id`);
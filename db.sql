START TRANSACTION;
BEGIN;
CREATE DATABASE fidibo;
CREATE TABLE fidibo.books (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(191),
    cover VARCHAR(191),
    type TINYINT UNSIGNED
);
CREATE USER 'book_service_user'@'%' IDENTIFIED BY 'book_service_password';
GRANT ALL PRIVILEGES ON fidibo.* TO 'book_service_user'@'%' WITH GRANT OPTION;
COMMIT;

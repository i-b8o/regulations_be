CREATE USER 'doc_user'@'localhost' IDENTIFIED BY '031501';
GRANT ALL PRIVILEGES ON * . * TO 'doc_user'@'localhost';
FLUSH PRIVILEGES;

SET GLOBAL sql_mode='';

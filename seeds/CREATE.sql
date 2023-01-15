SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

CREATE DATABASE golang;

CREATE TABLE golang.users (
  `id` int(11) NOT NULL,
  `nome` varchar(255) DEFAULT NULL,
  `senha` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO golang.users (`id`, `nome`, `senha`) VALUES
(1, 'vitor', '123'),
(2, 'jeff', '321');

ALTER TABLE golang.users
  ADD PRIMARY KEY (`id`);

ALTER TABLE golang.users
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;
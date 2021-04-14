CREATE TABLE `pelicula`
(
    `id`          INT(11) NOT NULL AUTO_INCREMENT,
    `nombre`      VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `director`    VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `escritor`    VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `pais`        VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `idioma`      VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
    `lanzamiento` INT(11) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB;

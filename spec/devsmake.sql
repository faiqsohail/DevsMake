--
-- Database: `devsmake`
-- 

-- --------------------------------------------------------

CREATE TABLE `accounts` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`provider` VARCHAR(32) NOT NULL,
	`provider_id` BIGINT(64) NOT NULL,
	`username` VARCHAR(64) NOT NULL,
	`email` VARCHAR(255),
	`modified` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB;

ALTER TABLE accounts
ADD COLUMN `points` INT NOT NULL DEFAULT 0 AFTER `username`,
ALGORITHM=INPLACE, LOCK=NONE;
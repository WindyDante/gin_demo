CREATE TABLE `userinfo` (
                            `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
                            `username` VARCHAR(64) NULL,
                            `department` VARCHAR(64) NULL,
                            `created` DATE NULL
);

CREATE TABLE `userdetail` (
                              `uid` INT(10) NULL,
                              `intro` TEXT NULL,
                              `profile` TEXT NULL,
                              PRIMARY KEY (`uid`)
);

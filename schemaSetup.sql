-- run this file to get started

create database accountancyApp;

use accountancyApp;


-- create accountHolders table and seed data

CREATE TABLE `accountHolders` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `jobTitle` varchar(200) DEFAULT NULL,
  `updated` int(11) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

INSERT into accountHolders (name, jobTitle, updated, created) VALUES ("Jim Davies", "Sales North East", UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT into accountHolders (name, jobTitle, updated, created) VALUES ("Darrel Mathes", "Sales South East", UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT into accountHolders (name, jobTitle, updated, created) VALUES ("Michael Rupert", "Sales North West", UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT into accountHolders (name, jobTitle, updated, created) VALUES ("Jim Davison", "Sales South West", UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


-- create currencies table and seed data

CREATE TABLE `currencies` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `exchangeRate` float DEFAULT '1',
  `updated` int(11) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT into currencies (name, exchangeRate, updated, created) VALUES ("US Dollar", 1.0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT into currencies (name, exchangeRate, updated, created) VALUES ("Canadian Dollar", 0.82, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT into currencies (name, exchangeRate, updated, created) VALUES ("UK POUND Sterling", 1.51, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT into currencies (name, exchangeRate, updated, created) VALUES ("Euro", 1.12, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


-- create accounts table and seed data
CREATE TABLE `accounts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `accountHolderId` int(11) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `address` varchar(400) DEFAULT NULL,
  `currentCredit` int(11) DEFAULT '0',
  `currentDebit` int(11) DEFAULT '0',
  `currencyId` int(11) DEFAULT NULL,
  `updated` int(11) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

INSERT into accounts (accountHolderId, name, address, currencyId, updated, created) VALUES ( (SELECT id AS accountHolderId from accountHolders WHERE `name` = "Jim Davies") , "Miami Parts", "20 Main Street", (SELECT id AS currencyId from currencies WHERE `name` = "US DOLLAR"), UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT into accounts (accountHolderId, name, address, currencyId, updated, created) VALUES ( (SELECT id AS accountHolderId from accountHolders WHERE `name` = "Jim Davies") , "Atlanta WholeSale", "200 Centre Avenue", (SELECT id AS currencyId from currencies WHERE `name` = "US DOLLAR"), UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- create transactions table and seed data
CREATE TABLE `transactions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `accountId` int(11) DEFAULT NULL,
  `details` varchar(500) DEFAULT NULL,
  `paymentOrProduct` enum('payment','product') DEFAULT 'product',
  `amount` decimal(8,2) DEFAULT NULL,
  `date` DATE DEFAULT NULL,
  `updated` int(11) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
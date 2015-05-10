ALTER TABLE accountHolders ADD `currencyId` int(11) DEFAULT 0;

ALTER TABLE accountHolders ADD `taxRateId` int(11) DEFAULT 0;


UPDATE accountHolders AS ah
JOIN accounts AS a ON a.accountHolderId = ah.id 
SET ah.currencyId = a.currencyId, ah.taxRateId = a.taxRateId;


ALTER TABLE accounts ADD `type` enum('tax','revenue','payment','commission') DEFAULT 'payment';

ALTER TABLE accounts DROP COLUMN currencyId;
ALTER TABLE accounts DROP COLUMN taxRateId;
ALTER TABLE accounts DROP COLUMN name;
ALTER TABLE accounts DROP COLUMN address;


INSERT into accountHolders (name, jobTitle, currencyId, taxRateId,updated, created) VALUES ( "Company", "", (SELECT id AS currencyId from currencies WHERE `name` = "US DOLLAR"), (SELECT id AS taxRateId from taxRates WHERE `name` = "US Tax"), UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

ALTER TABLE transactions DROP COLUMN paymentOrProduct;


DELETE FROM transactions;
DELETE  FROM accounts;

INSERT into accounts (id, updated, created, type)  VALUES (null, UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), "payment") ;
INSERT into accounts (id, updated, created, type)  VALUES (null, UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), "revenue") ;
INSERT into accounts (id, updated, created, type)  VALUES (null, UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), "tax") ;
INSERT into accounts (id, updated, created, type)  VALUES (null, UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), "commission") ;

ALTER TABLE accounts DROP COLUMN accountHolderId;

rename Table accounts TO accountTypes;


ALTER TABLE transactions ADD `accountTypeId` int not null;

ALTER TABLE transactions ADD `accountHolderId` int not null;

ALTER TABLE transactions DROP COLUMN accountId;
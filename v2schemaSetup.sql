


ALTER TABLE accountHolders ADD `currencyId` int(11) DEFAULT 0;

ALTER TABLE accountHolders ADD `taxRateId` int(11) DEFAULT 0;


UPDATE accountHolders AS ah
JOIN accounts AS a ON a.accountHolderId = ah.id 
SET ah.currencyId = a.currencyId, ah.taxRateId = a.taxRateId;







ALTER TABLE accounts ADD `type` enum('tax','revenue','product') DEFAULT 'product';

ALTER TABLE accounts DROP COLUMN currencyId;
ALTER TABLE accounts DROP COLUMN taxRateId;
ALTER TABLE accounts DROP COLUMN name;
ALTER TABLE accounts DROP COLUMN address;


INSERT into accountHolders (name, jobTitle, currencyId, taxRateId,updated, created) VALUES ( "Company", "", (SELECT id AS currencyId from currencies WHERE `name` = "US DOLLAR"), (SELECT id AS taxRateId from taxRates WHERE `name` = "US Tax"), UNIX_TIMESTAMP(), UNIX_TIMESTAMP());



-- ALTER TABLE transactions CHANGE accountId accountHolderId INT;

ALTER TABLE transactions DROP COLUMN paymentOrProduct;


DELETE FROM transactions;
DELETE  FROM accounts;

INSERT into accounts select null AS id, id AS accountHolderId, UNIX_TIMESTAMP() AS updated, UNIX_TIMESTAMP() AS created, "revenue" AS type
from accountHolders;

INSERT into accounts select null AS id, id AS accountHolderId, UNIX_TIMESTAMP() AS updated, UNIX_TIMESTAMP() AS created, "tax" AS type
from accountHolders;

INSERT into accounts select null AS id, id AS accountHolderId, UNIX_TIMESTAMP() AS updated, UNIX_TIMESTAMP() AS created, "product" AS type
from accountHolders;
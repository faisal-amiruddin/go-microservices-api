CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS customers;

CREATE TABLE customers (
  customer_id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  date_of_birth DATE NOT NULL,
  city VARCHAR(100) NOT NULL,
  zipcode VARCHAR(10) NOT NULL,
  status TINYINT(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (customer_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO customers VALUES
(2000, 'Steve', '1978-12-15', 'Delhi', '110075',1),
(2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550',1),
(2002, 'Hadley', '1988-04-30', 'Englewood, NJ','07631',1),
(2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', 0),
(2004, 'Nina', '1988-05-14','Clarkston, MI','48348',1),
(2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782',0);

CREATE TABLE accounts (
  account_id INT(11) NOT NULL AUTO_INCREMENT,
  customer_id INT(11) NOT NULL,
  opening_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  account_type VARCHAR(10) NOT NULL,
  pin VARCHAR(10) NOT NULL,
  status TINYINT(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (account_id),
  KEY accounts_fk (customer_id),
  CONSTRAINT accounts_fk FOREIGN KEY (customer_id)
    REFERENCES customers(customer_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO accounts VALUES
(95470, 2000, '2020-08-22 10:20:06', 'Saving', '1075',1),
(95471, 2001, '2020-06-15 10:27:22', 'Saving', '1255',1),
(95472, 2002, '2020-08-09 10:27:22', 'Checking', '0763',1),
(95473, 2000, '2020-06-03 10:27:22', 'Saving', '0310',1),
(95474, 2004, '2020-02-27 10:27:22', 'Checking', '4834',1),
(95475, 2005, '2020-03-20 10:27:22', 'Saving', '2078', 0);

CREATE TABLE transactions (
  transaction_id INT(11) NOT NULL AUTO_INCREMENT,
  account_id INT(11) NOT NULL,
  amount INT(11) NOT NULL,
  transaction_type VARCHAR(10) NOT NULL,
  transaction_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (transaction_id),
  KEY transactions_fk (account_id),
  CONSTRAINT transactions_fk FOREIGN KEY (account_id)
    REFERENCES accounts(account_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

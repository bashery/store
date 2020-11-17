CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    username varchar(100),
    password varchar(100),
    email varchar(100),
    phon varchar(100),
    linkavatar varchar(255),
    primary key (id)
);

create table tutorials_tbl( 
	usreid INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(100) NOT NULL,
	password VARCHAR(40) NOT NULL,
	phon VARCHAR(14) NOT NULL,
	email VARCHAR(42)NOT NULL,
	PRIMARY KEY ( userid )
);

CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    title varchar(150), photos varchar(1000),
    ts TIMESTAMP,primary key (id)
);

// --------------- general info -----------------

// add new column in spicial position:
ALTER TABLE table_name ADD new_column_name column_definition  [FIRST | AFTER column_name];
example:
ALTER TABLE products ADD description varchar(500) AFTER title;

// remove table:
DROP TABLE table_name;

// remove column:
ALTER TABLE products DROP COLUMN owner;


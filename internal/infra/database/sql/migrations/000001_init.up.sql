CREATE TABLE orders(
    id varchar(36) NOT NULL PRIMARY KEY,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL
);

INSERT INTO orders (id, price, tax, final_price) VALUES ("AAAA",10.0, 1.0, 100.0);
INSERT INTO orders (id, price, tax, final_price) VALUES ("BBBB",20.0, 2.0, 200.0);
INSERT INTO orders (id, price, tax, final_price) VALUES ("CCCC",30.0, 3.0, 300.0);
INSERT INTO orders (id, price, tax, final_price) VALUES ("DDDD",40.0, 4.0, 400.0);
INSERT INTO orders (id, price, tax, final_price) VALUES ("EEEE",50.0, 5.0, 500.0);
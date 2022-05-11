-- CREATE DATABASE IF NOT EXISTS mortgage_project;
-- USE mortgage_project;
-- CREATE USER IF NOT EXISTS xxxxx IDENTIFIED BY 'xxxxx';
-- GRANT ALL ON mortgage_project.* TO xxxxx;

CREATE TABLE IF NOT EXISTS houses (
  id INT AUTO_INCREMENT PRIMARY KEY,
  price INT NOT NULL,
  min_downpayment INT NOT NULL,
  property_tax SMALLINT NOT NULL,
  mainteance_fee SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS customers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  credit_score SMALLINT NOT NULL,
  salary INT NOT NULL,
  downpayment INT NOT NULL,
  house_id INT NOT NULL,
  CONSTRAINT fk_house_id 
    FOREIGN KEY (house_id) 
      REFERENCES houses(id)
);

INSERT INTO houses (price, min_downpayment, property_tax, mainteance_fee) VALUES
  (2000000, 400000, 9200, 1800),
  (1750000, 350000, 8050, 1800),
  (1500000, 300000, 6900, 1800),
  (1250000, 250000, 5750, 1800);

INSERT INTO customers (first_name, last_name, credit_score, salary, downpayment, house_id) VALUES
  ('Nick', 'White', 670, 253000, 320000, 3),
  ('Tom', 'Jones', 770, 264000, 800000, 1),
  ('Nicole', 'Freedman', 750, 203000, 850000, 2),
  ('June', 'Green', 700, 148000, 600000, 4);

CREATE TABLE IF NOT EXISTS houses (
  id SERIAL PRIMARY KEY,
  price INTEGER NOT NULL,
  min_downpayment INTEGER NOT NULL,
  property_tax SMALLINT NOT NULL,
  mainteance_fee SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS customers (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  credit_score SMALLINT NOT NULL,
  salary INTEGER NOT NULL,
  downpayment INTEGER NOT NULL,
  house_id INTEGER NOT NULL,
  CONSTRAINT fk_house_id 
    FOREIGN KEY (house_id) 
      REFERENCES houses(id) 
        ON DELETE SET NULL
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

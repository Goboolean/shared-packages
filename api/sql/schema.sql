CREATE TABLE stock_meta (
  stock_id     varchar(32) NOT NULL,
  stock_code   varchar(32) NOT NULL,
  fetch_origin varchar(4)  NOT NULL,
  stock_name   varchar(32) NOT NULL,
  created_at   timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (stock_id),
  UNIQUE (stock_code, fetch_origin)
);

CREATE TABLE store_log (
  stock_id  varchar(32) NOT NULL,
  stored_at timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  status    string      NOT NULL,

  PRIMARY KEY (stored_at, stock_id),
  FOREIGN KEY (stock_id) REFERENCES stock_meta (stock_id)
);
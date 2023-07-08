-- metadata table for product data
CREATE TABLE stock_meta (
  hash         varchar(32) NOT NULL, /* an identifier */

  stock_name   varchar(32) NOT NULL, /* human readable stock name */
  symbol       varchar(4)  NOT NULL, /* stock symbol */
  description  varchar(128)        , /* stock description */

  product_type varchar(8)  NOT NULL, /* examples are stock, encrypt */
  exchange     varchar(32) NOT NULL, /* examples are kospi, nasdaq. */
  location     varchar(32) NOT NULL, /* examples are korea, usa.    */

  PRIMARY KEY (hash)
);

CREATE TABLE stock_platform (
  platform   varchar(16) NOT NULL, /* available platform is buycycle, polygon, kis */
  identifier varchar(32) NOT NULL, /* a string that is used to specific stock on such platform query */
  stock_hash varchar(32) NOT NULL,

  PRIMARY KEY (stock_hash, platform),
  FOREIGN KEY (stock_hash) REFERENCES stock_meta (hash)
);

CREATE TABLE store_log (
  stock_hash  varchar(32) NOT NULL,
  stored_at timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  status    string      NOT NULL,

  PRIMARY KEY (stored_at, stock_id),
  FOREIGN KEY (stock_hash) REFERENCES stock_meta (hash)
);
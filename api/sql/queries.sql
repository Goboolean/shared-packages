-- name: CreateAccessInfo :exec
INSERT INTO store_log (stock_hash, status) VALUES (?, ?);

-- name: InsertNewStockMeta :exec
INSERT INTO stock_meta (hash, stock_name, symbol, description, product_type, exchange, location) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: InsertNewStockPlatformMeta :exec
INSERT INTO stock_platform (platform, identifier, stock_hash) VALUES (?, ?, ?);

-- name: CheckStockExist :one
SELECT EXISTS(SELECT 1 FROM stock_meta WHERE hash = (?));

-- name: GetStockMeta :one
SELECT hash, stock_name, symbol, description, product_type, exchange, location FROM stock_meta WHERE hash = (?);

-- name: GetAllStockMetaList :many
SELECT hash, stock_name, symbol, description, product_type, exchange, location FROM stock_meta;

-- name: GetStockMetaWithPlatform :one
SELECT hash, stock_name, symbol, description, product_type, exchange, location, platform, identifier FROM stock_meta JOIN stock_platform ON stock_meta.hash = stock_platform.stock_hash WHERE stock_hash = (?);

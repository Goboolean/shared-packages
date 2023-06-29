-- name: CreateAccessInfo :exec
INSERT INTO store_log (stock_id, status) VALUES (stock_id, status);

-- name: CheckStockExist :one
SELECT EXISTS(SELECT 1 FROM stock_meta WHERE stock_id = (?));

-- name: GetStockList :many
SELECT stock_id FROM stock_meta;

-- name: GetStockListByOrigin :many
SELECT stock_id FROM stock_meta WHERE fetch_origin = (?);

-- name: GetStockListWithDetail :many
SELECT stock_id, stock_code, fetch_origin, stock_name, created_at FROM stock_meta;

-- name: CreateStockMeta :exec
INSERT INTO stock_meta (stock_id, stock_code, fetch_origin, stock_name)
VALUES (stock_code || '&' || fetch_origin, stock_code, fetch_origin, stock_name);


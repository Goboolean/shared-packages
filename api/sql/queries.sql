-- name: CreateAccessInfo :exec
INSERT INTO access_log (connected_at) VALUES (NOW());

-- name: GetStockOrigin :one
SELECT fetch_origin FROM stock WHERE stock_id = (?);

-- name: CheckStockExist :one
SELECT EXISTS(SELECT 1 FROM stock WHERE stock_id = (?));
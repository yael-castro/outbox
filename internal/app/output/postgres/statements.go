package postgres

// SQL INSERT statements
const (
	insertPurchase       = `INSERT INTO purchases(order_id) VALUES ($1) RETURNING id`
	insertPurchaseOutbox = `INSERT INTO purchases_outbox(purchase_id, order_id) VALUES ($1, $2)`
)

// SQL SELECT statements
const (
	// TODO: replace this query to read from WAL
	selectPurchaseMessages = `
		SELECT
			id,
			purchase_id,
			order_id
		FROM purchases_outbox p
		WHERE delivered_at IS NULL
		ORDER BY order_id DESC
		LIMIT $1
	`

	updatePurchaseMessage = `UPDATE purchases_outbox SET updated_at = now(), delivered_at = now() WHERE id = $1`
)

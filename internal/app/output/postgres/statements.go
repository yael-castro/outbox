package postgres

// SQL INSERT statements
const (
	insertPurchase       = `INSERT INTO purchases(order_id) VALUES ($1) RETURNING id`
	insertPurchaseOutbox = `INSERT INTO purchases_outbox(purchase_id, order_id) VALUES ($1, $2)`
)

// SQL SELECT statements
const (
	// TODO: replace this query to read from WAL
	selectPurchaseMessage = `
		SELECT 
			purchase_id,
			order_id
		FROM purchases p
		WHERE NOT delivered
		ORDER BY order_id DESC
		LIMIT $1
	`
)

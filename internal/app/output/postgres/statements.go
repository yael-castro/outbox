package postgres

// SQL INSERT statements
const (
	insertPurchase      = `INSERT INTO purchases(order_id) VALUES ($1) RETURNING id`
	insertOutboxMessage = `INSERT INTO outbox_messages(topic, partition_key, headers, value) VALUES ($1, $2, $3, $4)`
)

// SQL SELECT statements
const (
	selectPurchaseMessages = `
		SELECT
			id,
			topic,
			partition_key,
			headers,
			"value"
		FROM outbox_messages
		WHERE delivered_at IS NULL
		ORDER BY created_at ASC
		LIMIT $1
	`

	updatePurchaseMessage = `UPDATE outbox_messages SET updated_at = now(), delivered_at = now() WHERE id = $1`
)

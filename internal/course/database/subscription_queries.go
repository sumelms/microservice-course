package database

const (
	createSubscription  = "create subscription"
	deleteSubscription  = "delete subscription by uuid"
	getSubscription     = "get subscription by uuid"
	listSubscription    = "list subscriptions"
	updateSubscription  = "update subscription by uuid"
	courseSubscriptions = "list subscriptions by course"
	userSubscriptions   = "list subscriptions by user"
)

func queriesSubscription() map[string]string {
	return map[string]string{
		createSubscription: `INSERT INTO subscriptions (course_id, matrix_id, user_id, expires_at) 
								VALUES ($1, $2, $3, $4) RETURNING *`,
		deleteSubscription: `UPDATE subscriptions SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`,
		getSubscription:    `SELECT * FROM subscriptions WHERE id = $1 AND deleted_at IS NULL`,
		listSubscription:   `SELECT * FROM subscriptions WHERE deleted_at IS NULL`,
		updateSubscription: `UPDATE subscriptions SET user_id = $1, course_id = $2, matrix_id = $3, expires_at = $4 
								WHERE id = $5 AND deleted_at IS NULL RETURNING *`,
		courseSubscriptions: `SELECT * FROM subscriptions WHERE course_id = $1 AND deleted_at IS NULL`,
		userSubscriptions:   `SELECT * FROM subscriptions WHERE user_id = $1 and deleted_at IS NULL`,
	}
}

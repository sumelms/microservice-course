package database

const (
	createSubscription = "create subscription"
	deleteSubscription = "delete subscription by uuid"
	getSubscription    = "get subscription by uuid"
	listSubscription   = "list subscriptions"
	updateSubscription = "update subscription by uuid"
)

type Query string

func queries() map[string]Query {
	return map[string]Query{
		createSubscription: Query("INSERT INTO subscriptions (course_id, matrix_id, user_id, valid_until) VALUES ($1, $2, $3, $4) RETURNING *"),
		deleteSubscription: Query("UPDATE subscriptions SET deleted_at = NOW() WHERE uuid = $1"),
		getSubscription:    Query("SELECT * FROM subscriptions WHERE uuid = $1"),
		listSubscription:   Query("SELECT * FROM subscriptions"),
		updateSubscription: Query("UPDATE subscriptions SET user_id = $1, course_id = $2, matrix_id = $3, valid_until = $4 WHERE uuid = $5 RETURNING *"),
	}
}

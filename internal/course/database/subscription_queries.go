package database

const (
	createSubscription = "create subscription"
	deleteSubscription = "delete subscription by uuid"
	getSubscription    = "get subscription by uuid"
	listSubscription   = "list subscriptions"
	updateSubscription = "update subscription by uuid"
)

func subscriptionQueries() map[string]string {
	return map[string]string{
		createSubscription: "INSERT INTO subscriptions (course_id, matrix_id, user_id, valid_until) VALUES ($1, $2, $3, $4) RETURNING *",
		deleteSubscription: "UPDATE subscriptions SET deleted_at = NOW() WHERE id = $1",
		getSubscription:    "SELECT * FROM subscriptions WHERE id = $1",
		listSubscription:   "SELECT * FROM subscriptions",
		updateSubscription: "UPDATE subscriptions SET user_id = $1, course_id = $2, matrix_id = $3, valid_until = $4 WHERE id = $5 RETURNING *",
	}
}

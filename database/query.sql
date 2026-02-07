-- name: GetUserPreference :one
SELECT value FROM user_preferences
WHERE key = ? LIMIT 1;

-- name: SetUserPreference :exec
INSERT INTO user_preferences (key, value)
VALUES (?, ?)
ON CONFLICT(key) DO UPDATE SET value = excluded.value;

package storage

type Blog struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	IsCompleted bool   `db:"is_completed"`
}

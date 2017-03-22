package app

import (
	"database/sql"
	"time"
)

// Task represents a cron task
type task struct {
	ID        uint          `json:"id"`
	HostID    uint          `json:"host_id"`
	URLID     string        `json:"url_uuid"`
	Name      string        `json:"name"`
	Frequency time.Duration `json:"frequency"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at"`
}

func migrateTask(db *sql.DB) error {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			host_id INTEGER NOT NULL,
			url_uuid VARCHAR NOT NULL,
			frequency VARCHAR NOT NULL,
			name VARCHAR NOT NULL,
			ip_address VARCHAR NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			FOREIGN KEY(host_id) REFERENCES hosts(id)
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		return err
	}

	return nil
}

package utils

import (
	"database/sql"
	"time"
)

func ResolveNullString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func ResolveNullTime(ns sql.NullTime) *time.Time {
	if ns.Valid {
		return &ns.Time
	}
	return nil
}

package resources

import (
	"2018_2_vi_studio_server/resources/sessions"
	"2018_2_vi_studio_server/resources/users"
	"database/sql"
)

type StorageBundle struct {
	Sessions *sessions.SessionStorage
	Users    *users.UserStorage
}

func NewStorageBundle(db *sql.DB) *StorageBundle {
	return &StorageBundle{
		Sessions: sessions.NewSessionStorage(db),
		Users:    users.NewUserStorage(db),
	}
}

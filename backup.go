package govultr

import (
	"context"
	"net/http"
)

// BackupService is the interface to interact with the backup endpoint on the Vultr API
// Link: https://www.vultr.com/api/#backup
type BackupService interface {
	GetList(ctx context.Context) ([]Backup, error)
}

// BackupServiceHandler handles interaction with the backup methods for the Vultr API
type BackupServiceHandler struct {
	client *Client
}

// Backup represents a Vultr backup
type Backup struct {
	BackupID    string `json:"BACKUPID"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Status      string `json:"status"`
}

// GetList retrieves a list of all backups on the current account
func (b *BackupServiceHandler) GetList(ctx context.Context) ([]Backup, error) {
	uri := "/v1/backup/list"
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	backupsMap := make(map[string]Backup)

	err = b.client.DoWithContext(ctx, req, &backupsMap)
	if err != nil {
		return nil, err
	}

	var backups []Backup
	for _, backup := range backupsMap {
		backups = append(backups, backup)
	}

	return backups, nil
}

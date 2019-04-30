package govultr

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// ServerService is the interface to interact with the server endpoints on the Vultr API
// Link: https://www.vultr.com/api/#server
type ServerService interface {
	ChangeApp(ctx context.Context, vpsID, appID string) error
	ListApps(ctx context.Context, vpsID string) ([]Application, error)
	AppInfo(ctx context.Context, vpsID string) (*ServerAppInfo, error)
	EnableBackup(ctx context.Context, vpsID string) error
	DisableBackup(ctx context.Context, vpsID string) error
	GetBackupSchedule(ctx context.Context, vpsID string) (*BackupSchedule, error)
	SetBackupSchedule(ctx context.Context, vpsID string, backup *BackupSchedule) error
	RestoreBackup(ctx context.Context, vpsID, backupID string) error
	RestoreSnapshot(ctx context.Context, vpsID, snapshotID string) error
}

// ServerServiceHandler handles interaction with the server methods for the Vultr API
type ServerServiceHandler struct {
	client *Client
}

// ServerAppInfo represents information about the application on your VPS
type ServerAppInfo struct {
	AppInfo string `json:"app_info"`
}

// BackupSchedule represents a schedule of a backup that runs on a VPS
type BackupSchedule struct {
	Enabled  bool   `json:"enabled"`
	CronType string `json:"cron_type"`
	NextRun  string `json:"next_scheduled_time_utc"`
	Hour     int    `json:"hour"`
	Dow      int    `json:"dow"`
	Dom      int    `json:"dom"`
}

// ChangeApp changes the VPS to a different application.
func (s *ServerServiceHandler) ChangeApp(ctx context.Context, vpsID, appID string) error {

	uri := "/v1/server/app_change"

	values := url.Values{
		"SUBID": {vpsID},
		"APPID": {appID},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = s.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// ListApps retrieves a list of applications to which a virtual machine can be changed.
func (s *ServerServiceHandler) ListApps(ctx context.Context, vpsID string) ([]Application, error) {

	uri := "/v1/server/app_change_list"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", vpsID)
	req.URL.RawQuery = q.Encode()

	var appMap map[string]Application
	err = s.client.DoWithContext(ctx, req, &appMap)

	if err != nil {
		return nil, err
	}

	var appList []Application
	for _, a := range appMap {
		appList = append(appList, a)
	}

	return appList, nil
}

// AppInfo retrieves the application information for a given VPS ID
func (s *ServerServiceHandler) AppInfo(ctx context.Context, vpsID string) (*ServerAppInfo, error) {

	uri := "/v1/server/get_app_info"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("SUBID", vpsID)
	req.URL.RawQuery = q.Encode()

	appInfo := new(ServerAppInfo)

	err = s.client.DoWithContext(ctx, req, appInfo)

	if err != nil {
		return nil, err
	}

	return appInfo, nil
}

// EnableBackup enables automatic backups on a given VPS
func (s *ServerServiceHandler) EnableBackup(ctx context.Context, vpsID string) error {

	uri := "/v1/server/backup_enable"

	values := url.Values{
		"SUBID": {vpsID},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = s.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// DisableBackup disable automatic backups on a given VPS
func (s *ServerServiceHandler) DisableBackup(ctx context.Context, vpsID string) error {

	uri := "/v1/server/backup_disable"

	values := url.Values{
		"SUBID": {vpsID},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = s.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetBackupSchedule retrieves the backup schedule for a given vps - all time values are in UTC
func (s *ServerServiceHandler) GetBackupSchedule(ctx context.Context, vpsID string) (*BackupSchedule, error) {

	uri := "/v1/server/backup_get_schedule"

	values := url.Values{
		"SUBID": {vpsID},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return nil, err
	}

	backup := new(BackupSchedule)
	err = s.client.DoWithContext(ctx, req, backup)

	if err != nil {
		return nil, err
	}

	return backup, nil
}

// SetBackupSchedule sets the backup schedule for a given vps - all time values are in UTC
func (s *ServerServiceHandler) SetBackupSchedule(ctx context.Context, vpsID string, backup *BackupSchedule) error {

	uri := "/v1/server/backup_set_schedule"

	values := url.Values{
		"SUBID":     {vpsID},
		"cron_type": {backup.CronType},
		"hour":      {strconv.Itoa(backup.Hour)},
		"dow":       {strconv.Itoa(backup.Dow)},
		"dom":       {strconv.Itoa(backup.Dom)},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = s.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// RestoreBackup will restore the specified backup to the given VPS
func (s *ServerServiceHandler) RestoreBackup(ctx context.Context, vpsID, backupID string) error {

	uri := "/v1/server/restore_backup"

	values := url.Values{
		"SUBID":    {vpsID},
		"BACKUPID": {backupID},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = s.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// RestoreSnapshot will restore the specified snapshot to the given VPS
func (s *ServerServiceHandler) RestoreSnapshot(ctx context.Context, vpsID, snapshotID string) error {

	uri := "/v1/server/restore_snapshot"

	values := url.Values{
		"SUBID":      {vpsID},
		"SNAPSHOTID": {snapshotID},
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, values)

	if err != nil {
		return err
	}

	err = s.client.DoWithContext(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

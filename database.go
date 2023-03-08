package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const databasePath = "/v2/databases"

// DatabaseService is the interface to interact with the Database endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/managed-databases
type DatabaseService interface {
	ListPlans(ctx context.Context, options *DBPlanListOptions) ([]DatabasePlan, *Meta, error)

	List(ctx context.Context, options *DBListOptions) ([]Database, *Meta, error)
	Create(ctx context.Context, databaseReq *DatabaseCreateReq) (*Database, error)
	Get(ctx context.Context, databaseID string) (*Database, error)
	Update(ctx context.Context, databaseID string, databaseReq *DatabaseUpdateReq) (*Database, error)
	Delete(ctx context.Context, databaseID string) error

	ListUsers(ctx context.Context, databaseID string) ([]DatabaseUser, *Meta, error)
	CreateUser(ctx context.Context, databaseID string, databaseUserReq *DatabaseUserCreateReq) (*DatabaseUser, error)
	GetUser(ctx context.Context, databaseID string, username string) (*DatabaseUser, error)
	UpdateUser(ctx context.Context, databaseID string, username string, databaseUserReq *DatabaseUserUpdateReq) (*DatabaseUser, error)
	DeleteUser(ctx context.Context, databaseID string, username string) error

	ListDBs(ctx context.Context, databaseID string) ([]DatabaseDB, *Meta, error)
	CreateDB(ctx context.Context, databaseID string, databaseDBReq *DatabaseDBCreateReq) (*DatabaseDB, error)
	GetDB(ctx context.Context, databaseID string, dbname string) (*DatabaseDB, error)
	DeleteDB(ctx context.Context, databaseID string, dbname string) error

	ListMaintenanceUpdates(ctx context.Context, databaseID string) ([]string, error)
	StartMaintenance(ctx context.Context, databaseID string) (string, error)

	ListServiceAlerts(ctx context.Context, databaseID string, databaseAlertsReq *DatabaseListAlertsReq) ([]DatabaseAlert, error)
}

// DatabaseServiceHandler handles interaction with the server methods for the Vultr API
type DatabaseServiceHandler struct {
	client *Client
}

type DBPlanListOptions struct {
	Engine string `url:"engine,omitempty"`
	Nodes  int    `url:"nodes,omitempty"`
	Region string `url:"region,omitempty"`
}

// DatabasePlan represents a Managed Database plan
type DatabasePlan struct {
	ID               string           `json:"id"`
	NumberOfNodes    int              `json:"number_of_nodes"`
	Type             string           `json:"type"`
	VCPUCount        int              `json:"vcpu_count"`
	RAM              int              `json:"ram"`
	Disk             int              `json:"disk"`
	MonthlyCost      int              `json:"monthly_cost"`
	SupportedEngines SupportedEngines `json:"supported_engines"`
	MaxConnections   *MaxConnections  `json:"max_connections,omitempty"`
	Locations        []string         `json:"locations"`
}

// SupportedEngines represents an object containing supported database engine types for Managed Database plans
type SupportedEngines struct {
	MySQL *bool `json:"mysql"`
	PG    *bool `json:"pg"`
	Redis *bool `json:"redis"`
}

// MaxConnections represents an object containing the maximum number of connections by engine type for Managed Database plans
type MaxConnections struct {
	MySQL int `json:"mysql,omitempty"`
	PG    int `json:"pg,omitempty"`
}

type databasePlansBase struct {
	DatabasePlans []DatabasePlan `json:"plans"`
	Meta          *Meta          `json:"meta"`
}

type DBListOptions struct {
	Label  string `url:"label,omitempty"`
	Tag    string `url:"tag,omitempty"`
	Region string `url:"region,omitempty"`
}

// Database represents a Managed Database subscription
type Database struct {
	ID                     string        `json:"id"`
	DateCreated            string        `json:"date_created"`
	Plan                   string        `json:"plan"`
	PlanDisk               int           `json:"plan_disk"`
	PlanRAM                int           `json:"plan_ram"`
	PlanVCPUs              int           `json:"plan_vcpus"`
	PlanReplicas           int           `json:"plan_replicas"`
	Region                 string        `json:"region"`
	Status                 string        `json:"status"`
	Label                  string        `json:"label"`
	Tag                    string        `json:"tag"`
	DatabaseEngine         string        `json:"database_engine"`
	DatabaseEngineVersion  string        `json:"database_engine_version"`
	DBName                 string        `json:"dbname,omitempty"`
	Host                   string        `json:"host"`
	User                   string        `json:"user"`
	Password               string        `json:"password"`
	Port                   string        `json:"port"`
	MaintenanceDOW         string        `json:"maintenance_dow"`
	MaintenanceTime        string        `json:"maintenance_time"`
	LatestBackup           string        `json:"latest_backup"`
	TrustedIPs             []string      `json:"trusted_ips"`
	MySQLSQLModes          []string      `json:"mysql_sql_modes,omitempty"`
	MySQLRequirePrimaryKey *bool         `json:"mysql_require_primary_key,omitempty"`
	MySQLSlowQueryLog      *bool         `json:"mysql_slow_query_log,omitempty"`
	MySQLLongQueryTime     int           `json:"mysql_long_query_time,omitempty"`
	PGAvailableExtensions  []PGExtension `json:"pg_available_extensions,omitempty"`
	RedisEvictionPolicy    string        `json:"redis_eviction_policy,omitempty"`
	ClusterTimeZone        string        `json:"cluster_time_zone,omitempty"`
	ReadReplicas           []Database    `json:"read_replicas,omitempty"`
}

// PGExtension represents an object containing extension name and version information
type PGExtension struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions,omitempty"`
}

type databasesBase struct {
	Databases []Database `json:"databases"`
	Meta      *Meta      `json:"meta"`
}

type databaseBase struct {
	Database *Database `json:"database"`
}

// DatabaseCreateReq struct used to create a database.
type DatabaseCreateReq struct {
	DatabaseEngine         string   `json:"database_engine,omitempty"`
	DatabaseEngineVersion  string   `json:"database_engine_version,omitempty"`
	Region                 string   `json:"region,omitempty"`
	Plan                   string   `json:"plan,omitempty"`
	Label                  string   `json:"label,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
	MaintenanceDOW         string   `json:"maintenance_dow,omitempty"`
	MaintenanceTime        string   `json:"maintenance_time,omitempty"`
	TrustedIPs             []string `json:"trusted_ips,omitempty"`
	MySQLSQLModes          []string `json:"mysql_sql_modes,omitempty"`
	MySQLRequirePrimaryKey *bool    `json:"mysql_require_primary_key,omitempty"`
	MySQLSlowQueryLog      *bool    `json:"mysql_slow_query_log,omitempty"`
	MySQLLongQueryTime     int      `json:"mysql_long_query_time,omitempty"`
	RedisEvictionPolicy    string   `json:"redis_eviction_policy,omitempty"`
}

// DatabaseUpdateReq struct used to update a dataase.
type DatabaseUpdateReq struct {
	DatabaseEngine         string   `json:"database_engine,omitempty"`
	DatabaseEngineVersion  string   `json:"database_engine_version,omitempty"`
	Region                 string   `json:"region,omitempty"`
	Plan                   string   `json:"plan,omitempty"`
	Label                  string   `json:"label,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
	MaintenanceDOW         string   `json:"maintenance_dow,omitempty"`
	MaintenanceTime        string   `json:"maintenance_time,omitempty"`
	ClusterTimeZone        string   `json:"cluster_time_zone,omitempty"`
	TrustedIPs             []string `json:"trusted_ips,omitempty"`
	MySQLSQLModes          []string `json:"mysql_sql_modes,omitempty"`
	MySQLRequirePrimaryKey *bool    `json:"mysql_require_primary_key,omitempty"`
	MySQLSlowQueryLog      *bool    `json:"mysql_slow_query_log,omitempty"`
	MySQLLongQueryTime     int      `json:"mysql_long_query_time,omitempty"`
	RedisEvictionPolicy    string   `json:"redis_eviction_policy,omitempty"`
}

// DatabaseUser represents a user within a Managed Database cluster
type DatabaseUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Encryption string `json:"encryption,omitempty"`
}

type databaseUserBase struct {
	DatabaseUser *DatabaseUser `json:"user"`
}

type databaseUsersBase struct {
	DatabaseUsers []DatabaseUser `json:"users"`
	Meta          *Meta          `json:"meta"`
}

// DatabaseUserCreateReq struct used to create a user within a Managed Database.
type DatabaseUserCreateReq struct {
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	Encryption string `json:"encryption,omitempty"`
}

// DatabaseUserUpdateReq struct used to update a user within a Managed Database.
type DatabaseUserUpdateReq struct {
	Password string `json:"password,omitempty"`
}

// DatabaseDB represents a logical database within a Managed Database cluster
type DatabaseDB struct {
	Name string `json:"name"`
}

type databaseDBBase struct {
	DatabaseDB *DatabaseDB `json:"db"`
}

type databaseDBsBase struct {
	DatabaseDBs []DatabaseDB `json:"dbs"`
	Meta        *Meta        `json:"meta"`
}

// DatabaseDBCreateReq struct used to create a logical database within a Managed Database.
type DatabaseDBCreateReq struct {
	Name string `json:"name"`
}

type databaseUpdatesBase struct {
	AvailableUpdates []string `json:"available_updates"`
	Message          string   `json:"message"`
}

// DatabaseAlert represents a service alert for a Managed Database cluster
type DatabaseAlert struct {
	Timestamp            string `json:"timestamp"`
	MessageType          string `json:"message_type"`
	Description          string `json:"description"`
	Recommendation       string `json:"recommendation,omitempty"`
	MaintenanceScheduled string `json:"maintenance_scheduled,omitempty"`
	ResourceType         string `json:"resource_type,omitempty"`
	TableCount           int    `json:"table_count,omitempty"`
}

type databaseAlertsBase struct {
	DatabaseAlerts []DatabaseAlert `json:"alerts"`
}

// DatabaseListAlertsReq struct used to query service alerts for a Managed Database.
type DatabaseListAlertsReq struct {
	Period string `json:"period"`
}

// List all database plans
func (i *DatabaseServiceHandler) ListPlans(ctx context.Context, options *DBPlanListOptions) ([]DatabasePlan, *Meta, error) {
	uri := fmt.Sprintf("%s/plans", databasePath)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	databasePlans := new(databasePlansBase)
	if err = i.client.DoWithContext(ctx, req, databasePlans); err != nil {
		return nil, nil, err
	}

	return databasePlans.DatabasePlans, databasePlans.Meta, nil
}

// List all databases on your account.
func (i *DatabaseServiceHandler) List(ctx context.Context, options *DBListOptions) ([]Database, *Meta, error) {
	req, err := i.client.NewRequest(ctx, http.MethodGet, databasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	databases := new(databasesBase)
	if err = i.client.DoWithContext(ctx, req, databases); err != nil {
		return nil, nil, err
	}

	return databases.Databases, databases.Meta, nil
}

// Create will create the Managed Database with the given parameters
func (i *DatabaseServiceHandler) Create(ctx context.Context, databaseReq *DatabaseCreateReq) (*Database, error) {
	req, err := i.client.NewRequest(ctx, http.MethodPost, databasePath, databaseReq)
	if err != nil {
		return nil, err
	}

	database := new(databaseBase)
	if err = i.client.DoWithContext(ctx, req, database); err != nil {
		return nil, err
	}

	return database.Database, nil
}

// Get will get the server with the given databaseID
func (i *DatabaseServiceHandler) Get(ctx context.Context, databaseID string) (*Database, error) {
	uri := fmt.Sprintf("%s/%s", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	database := new(databaseBase)
	if err = i.client.DoWithContext(ctx, req, database); err != nil {
		return nil, err
	}

	return database.Database, nil
}

// Update will update the Managed Database with the given parameters
func (i *DatabaseServiceHandler) Update(ctx context.Context, databaseID string, databaseReq *DatabaseUpdateReq) (*Database, error) {
	uri := fmt.Sprintf("%s/%s", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodPut, uri, databaseReq)
	if err != nil {
		return nil, err
	}

	database := new(databaseBase)
	if err := i.client.DoWithContext(ctx, req, database); err != nil {
		return nil, err
	}

	return database.Database, nil
}

// Delete a Managed database. All data will be permanently lost.
func (i *DatabaseServiceHandler) Delete(ctx context.Context, databaseID string) error {
	uri := fmt.Sprintf("%s/%s", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// List all database users on your Managed Database.
func (i *DatabaseServiceHandler) ListUsers(ctx context.Context, databaseID string) ([]DatabaseUser, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/users", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseUsers := new(databaseUsersBase)
	if err = i.client.DoWithContext(ctx, req, databaseUsers); err != nil {
		return nil, nil, err
	}

	return databaseUsers.DatabaseUsers, databaseUsers.Meta, nil
}

// Create a user within the Managed Database with the given parameters
func (i *DatabaseServiceHandler) CreateUser(ctx context.Context, databaseID string, databaseUserReq *DatabaseUserCreateReq) (*DatabaseUser, error) {
	uri := fmt.Sprintf("%s/%s/users", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, databaseUserReq)
	if err != nil {
		return nil, err
	}

	databaseUser := new(databaseUserBase)
	if err = i.client.DoWithContext(ctx, req, databaseUser); err != nil {
		return nil, err
	}

	return databaseUser.DatabaseUser, nil
}

// Get information on an individual user within a Managed Database based on a username and databaseID
func (i *DatabaseServiceHandler) GetUser(ctx context.Context, databaseID string, username string) (*DatabaseUser, error) {
	uri := fmt.Sprintf("%s/%s/users/%s", databasePath, databaseID, username)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	databaseUser := new(databaseUserBase)
	if err = i.client.DoWithContext(ctx, req, databaseUser); err != nil {
		return nil, err
	}

	return databaseUser.DatabaseUser, nil
}

// Update a user within the Managed Database with the given parameters
func (i *DatabaseServiceHandler) UpdateUser(ctx context.Context, databaseID string, username string, databaseUserReq *DatabaseUserUpdateReq) (*DatabaseUser, error) {
	uri := fmt.Sprintf("%s/%s/users/%s", databasePath, databaseID, username)

	req, err := i.client.NewRequest(ctx, http.MethodPut, uri, databaseUserReq)
	if err != nil {
		return nil, err
	}

	databaseUser := new(databaseUserBase)
	if err := i.client.DoWithContext(ctx, req, databaseUser); err != nil {
		return nil, err
	}

	return databaseUser.DatabaseUser, nil
}

// Delete a user within the Managed database. All data will be permanently lost.
func (i *DatabaseServiceHandler) DeleteUser(ctx context.Context, databaseID string, username string) error {
	uri := fmt.Sprintf("%s/%s/users/%s", databasePath, databaseID, username)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// List all logical databases on your Managed Database.
func (i *DatabaseServiceHandler) ListDBs(ctx context.Context, databaseID string) ([]DatabaseDB, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/dbs", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseDBs := new(databaseDBsBase)
	if err = i.client.DoWithContext(ctx, req, databaseDBs); err != nil {
		return nil, nil, err
	}

	return databaseDBs.DatabaseDBs, databaseDBs.Meta, nil
}

// Create a logical database within the Managed Database with the given parameters
func (i *DatabaseServiceHandler) CreateDB(ctx context.Context, databaseID string, databaseDBReq *DatabaseDBCreateReq) (*DatabaseDB, error) {
	uri := fmt.Sprintf("%s/%s/dbs", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, databaseDBReq)
	if err != nil {
		return nil, err
	}

	databaseDB := new(databaseDBBase)
	if err = i.client.DoWithContext(ctx, req, databaseDB); err != nil {
		return nil, err
	}

	return databaseDB.DatabaseDB, nil
}

// Get information on an individual logical database within a Managed Database based on a dbname and databaseID
func (i *DatabaseServiceHandler) GetDB(ctx context.Context, databaseID string, dbname string) (*DatabaseDB, error) {
	uri := fmt.Sprintf("%s/%s/dbs/%s", databasePath, databaseID, dbname)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	databaseDB := new(databaseDBBase)
	if err = i.client.DoWithContext(ctx, req, databaseDB); err != nil {
		return nil, err
	}

	return databaseDB.DatabaseDB, nil
}

// Delete a user within the Managed database. All data will be permanently lost.
func (i *DatabaseServiceHandler) DeleteDB(ctx context.Context, databaseID string, dbname string) error {
	uri := fmt.Sprintf("%s/%s/dbs/%s", databasePath, databaseID, dbname)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return i.client.DoWithContext(ctx, req, nil)
}

// List all available maintenance updates for your Managed Database.
func (i *DatabaseServiceHandler) ListMaintenanceUpdates(ctx context.Context, databaseID string) ([]string, error) {
	uri := fmt.Sprintf("%s/%s/maintenance", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	databaseUpdates := new(databaseUpdatesBase)
	if err = i.client.DoWithContext(ctx, req, databaseUpdates); err != nil {
		return nil, err
	}

	if databaseUpdates.AvailableUpdates == nil {
		return []string{databaseUpdates.Message}, nil
	}

	return databaseUpdates.AvailableUpdates, nil
}

// Start the maintenance update process for your Managed Database
func (i *DatabaseServiceHandler) StartMaintenance(ctx context.Context, databaseID string) (string, error) {
	uri := fmt.Sprintf("%s/%s/maintenance", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return "", err
	}

	databaseUpdates := new(databaseUpdatesBase)
	if err = i.client.DoWithContext(ctx, req, databaseUpdates); err != nil {
		return "", err
	}

	return databaseUpdates.Message, nil
}

// Query for service alerts for the Managed Database using the given parameters
func (i *DatabaseServiceHandler) ListServiceAlerts(ctx context.Context, databaseID string, databaseAlertsReq *DatabaseListAlertsReq) ([]DatabaseAlert, error) {
	uri := fmt.Sprintf("%s/%s/alerts", databasePath, databaseID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, databaseAlertsReq)
	if err != nil {
		return nil, err
	}

	databaseAlerts := new(databaseAlertsBase)
	if err = i.client.DoWithContext(ctx, req, databaseAlerts); err != nil {
		return nil, err
	}

	return databaseAlerts.DatabaseAlerts, nil
}

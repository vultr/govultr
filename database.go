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
	Get(ctx context.Context, databaseID string) (*Database, error)
}

// DatabaseServiceHandler handles interaction with the server methods for the Vultr API
type DatabaseServiceHandler struct {
	client *Client
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

type DBPlanListOptions struct {
	Engine string `url:"engine,omitempty"`
	Nodes  int    `url:"nodes,omitempty"`
	Region string `url:"region,omitempty"`
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

type databasesBase struct {
	Databases []Database `json:"databases"`
	Meta      *Meta      `json:"meta"`
}

type databaseBase struct {
	Database *Database `json:"database"`
}

// List all database plans
func (i *DatabaseServiceHandler) ListPlans(ctx context.Context, options *DBPlanListOptions) ([]DatabasePlan, *Meta, error) {
	req, err := i.client.NewRequest(ctx, http.MethodGet, databasePath+"/plans", nil)
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

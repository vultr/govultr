package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDatabaseServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/databases", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"databases": [
				{
					"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
					"date_created": "2022-05-09 10:13:31",
					"plan": "vultr-dbaas-business-cc-2-80-4",
					"plan_disk": 80,
					"plan_ram": 4096,
					"plan_vcpus": 2,
					"plan_replicas": 1,
					"region": "EWR",
					"status": "Running",
					"label": "testy-mc-testerton-the-8th",
					"tag": "bing bong",
					"database_engine": "mysql",
					"database_engine_version": "8",
					"dbname": "defaultdb",
					"host": "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
					"user": "vultradmin",
					"password": "AVNS_8E9hjx1LdsiA5EZ",
					"port": "16751",
					"maintenance_dow": "sunday",
					"maintenance_time": "02:00:00",
					"latest_backup": "2023-03-13 00:59:07",
					"trusted_ips": [],
					"mysql_sql_modes": [
						"ANSI",
						"ERROR_FOR_DIVISION_BY_ZERO",
						"NO_ENGINE_SUBSTITUTION",
						"NO_ZERO_DATE",
						"NO_ZERO_IN_DATE",
						"STRICT_ALL_TABLES"
					],
					"mysql_require_primary_key": true,
					"mysql_slow_query_log": false,
					"cluster_time_zone": "America/New_York",
					"read_replicas": [
						{
							"id": "daeb6d62-a6a2-458c-9f74-e053735d7f50",
							"date_created": "2022-05-09 10:12:43",
							"plan": "vultr-dbaas-startup-cc-2-80-4",
							"plan_disk": 80,
							"plan_ram": 4096,
							"plan_vcpus": 2,
							"plan_replicas": 0,
							"region": "EWR",
							"status": "Running",
							"label": "testy-mc-testerton-the-7th",
							"tag": "",
							"database_engine": "mysql",
							"database_engine_version": "8",
							"dbname": "defaultdb",
							"host": "vultr-prod-87086a7d-4bc8-47ca-aa88-f88138d82772-vultr-prod-85e0.vultrdb.com",
							"user": "vultradmin",
							"password": "AVNS_UBen_MjqAwDd2BWFc-Y",
							"port": "16751",
							"maintenance_dow": "sunday",
							"maintenance_time": "02:00:00",
							"latest_backup": "2023-03-12 22:07:06",
							"trusted_ips": [],
							"mysql_sql_modes": [
								"ANSI",
								"ERROR_FOR_DIVISION_BY_ZERO",
								"NO_ENGINE_SUBSTITUTION",
								"NO_ZERO_DATE",
								"NO_ZERO_IN_DATE",
								"STRICT_ALL_TABLES"
							],
							"mysql_require_primary_key": true,
							"mysql_slow_query_log": false,
							"cluster_time_zone": "America/New_York"
						}
					]
				}
			],
			"meta": {
				"total": 1
			}
		}`
		fmt.Fprint(writer, response)
	})

	database, meta, err := client.Database.List(ctx, nil)
	if err != nil {
		t.Errorf("Database.List returned %+v", err)
	}

	mysql_sql_modes := []string{
		"ANSI",
		"ERROR_FOR_DIVISION_BY_ZERO",
		"NO_ENGINE_SUBSTITUTION",
		"NO_ZERO_DATE",
		"NO_ZERO_IN_DATE",
		"STRICT_ALL_TABLES",
	}

	replicas := []Database{
		{
			ID:                     "daeb6d62-a6a2-458c-9f74-e053735d7f50",
			DateCreated:            "2022-05-09 10:12:43",
			Plan:                   "vultr-dbaas-startup-cc-2-80-4",
			PlanDisk:               80,
			PlanRAM:                4096,
			PlanVCPUs:              2,
			PlanReplicas:           0,
			Region:                 "EWR",
			Status:                 "Running",
			Label:                  "testy-mc-testerton-the-7th",
			Tag:                    "",
			DatabaseEngine:         "mysql",
			DatabaseEngineVersion:  "8",
			DBName:                 "defaultdb",
			Host:                   "vultr-prod-87086a7d-4bc8-47ca-aa88-f88138d82772-vultr-prod-85e0.vultrdb.com",
			User:                   "vultradmin",
			Password:               "AVNS_UBen_MjqAwDd2BWFc-Y",
			Port:                   "16751",
			MaintenanceDOW:         "sunday",
			MaintenanceTime:        "02:00:00",
			LatestBackup:           "2023-03-12 22:07:06",
			TrustedIPs:             []string{},
			MySQLSQLModes:          mysql_sql_modes,
			MySQLRequirePrimaryKey: BoolToBoolPtr(true),
			MySQLSlowQueryLog:      BoolToBoolPtr(false),
			ClusterTimeZone:        "America/New_York",
		},
	}

	expected := []Database{
		{
			ID:                     "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
			DateCreated:            "2022-05-09 10:13:31",
			Plan:                   "vultr-dbaas-business-cc-2-80-4",
			PlanDisk:               80,
			PlanRAM:                4096,
			PlanVCPUs:              2,
			PlanReplicas:           1,
			Region:                 "EWR",
			Status:                 "Running",
			Label:                  "testy-mc-testerton-the-8th",
			Tag:                    "bing bong",
			DatabaseEngine:         "mysql",
			DatabaseEngineVersion:  "8",
			DBName:                 "defaultdb",
			Host:                   "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
			User:                   "vultradmin",
			Password:               "AVNS_8E9hjx1LdsiA5EZ",
			Port:                   "16751",
			MaintenanceDOW:         "sunday",
			MaintenanceTime:        "02:00:00",
			LatestBackup:           "2023-03-13 00:59:07",
			TrustedIPs:             []string{},
			MySQLSQLModes:          mysql_sql_modes,
			MySQLRequirePrimaryKey: BoolToBoolPtr(true),
			MySQLSlowQueryLog:      BoolToBoolPtr(false),
			ClusterTimeZone:        "America/New_York",
			ReadReplicas:           replicas,
		},
	}

	if !reflect.DeepEqual(database, expected) {
		t.Errorf("Database.List returned %+v, expected %+v", database, expected)
	}

	expectedMeta := &Meta{
		Total: 1,
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Errorf("Database.List meta returned %+v, expected %+v", meta, expectedMeta)
	}
}

func TestDatabaseServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/databases", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"database": {
				"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
				"date_created": "2022-05-09 10:13:31",
				"plan": "vultr-dbaas-business-cc-2-80-4",
				"plan_disk": 80,
				"plan_ram": 4096,
				"plan_vcpus": 2,
				"plan_replicas": 1,
				"region": "EWR",
				"status": "Running",
				"label": "testy-mc-testerton-the-8th",
				"tag": "",
				"database_engine": "mysql",
				"database_engine_version": "8",
				"dbname": "defaultdb",
				"host": "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
				"user": "vultradmin",
				"password": "AVNS_8E9hjx1LdsiA5EZ",
				"port": "16751",
				"maintenance_dow": "sunday",
				"maintenance_time": "02:00:00",
				"latest_backup": "2023-03-13 00:59:07",
				"trusted_ips": [],
				"mysql_sql_modes": [
					"ANSI",
					"ERROR_FOR_DIVISION_BY_ZERO",
					"NO_ENGINE_SUBSTITUTION",
					"NO_ZERO_DATE",
					"NO_ZERO_IN_DATE",
					"STRICT_ALL_TABLES"
				],
				"mysql_require_primary_key": true,
				"mysql_slow_query_log": false,
				"cluster_time_zone": "America/New_York",
				"read_replicas": []
			}
		}`
		fmt.Fprint(writer, response)
	})

	options := &DatabaseCreateReq{
		DatabaseEngine:        "mysql",
		DatabaseEngineVersion: "8",
		Region:                "ewr",
		Plan:                  "vultr-dbaas-business-cc-2-80-4",
		Label:                 "testy-mc-testerton-the-8th",
	}

	database, err := client.Database.Create(ctx, options)
	if err != nil {
		t.Errorf("Database.Create returned %+v", err)
	}

	mysql_sql_modes := []string{
		"ANSI",
		"ERROR_FOR_DIVISION_BY_ZERO",
		"NO_ENGINE_SUBSTITUTION",
		"NO_ZERO_DATE",
		"NO_ZERO_IN_DATE",
		"STRICT_ALL_TABLES",
	}

	expected := &Database{
		ID:                     "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
		DateCreated:            "2022-05-09 10:13:31",
		Plan:                   "vultr-dbaas-business-cc-2-80-4",
		PlanDisk:               80,
		PlanRAM:                4096,
		PlanVCPUs:              2,
		PlanReplicas:           1,
		Region:                 "EWR",
		Status:                 "Running",
		Label:                  "testy-mc-testerton-the-8th",
		Tag:                    "",
		DatabaseEngine:         "mysql",
		DatabaseEngineVersion:  "8",
		DBName:                 "defaultdb",
		Host:                   "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
		User:                   "vultradmin",
		Password:               "AVNS_8E9hjx1LdsiA5EZ",
		Port:                   "16751",
		MaintenanceDOW:         "sunday",
		MaintenanceTime:        "02:00:00",
		LatestBackup:           "2023-03-13 00:59:07",
		TrustedIPs:             []string{},
		MySQLSQLModes:          mysql_sql_modes,
		MySQLRequirePrimaryKey: BoolToBoolPtr(true),
		MySQLSlowQueryLog:      BoolToBoolPtr(false),
		ClusterTimeZone:        "America/New_York",
		ReadReplicas:           []Database{},
	}

	if !reflect.DeepEqual(database, expected) {
		t.Errorf("Database.Create returned %+v, expected %+v", database, expected)
	}
}

func TestDatabaseServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/databases/999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"database": {
				"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
				"date_created": "2022-05-09 10:13:31",
				"plan": "vultr-dbaas-business-cc-2-80-4",
				"plan_disk": 80,
				"plan_ram": 4096,
				"plan_vcpus": 2,
				"plan_replicas": 1,
				"region": "EWR",
				"status": "Running",
				"label": "testy-mc-testerton-the-8th",
				"tag": "bing bong",
				"database_engine": "mysql",
				"database_engine_version": "8",
				"dbname": "defaultdb",
				"host": "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
				"user": "vultradmin",
				"password": "AVNS_8E9hjx1LdsiA5EZ",
				"port": "16751",
				"maintenance_dow": "sunday",
				"maintenance_time": "02:00:00",
				"latest_backup": "2023-03-13 00:59:07",
				"trusted_ips": [],
				"mysql_sql_modes": [
					"ANSI",
					"ERROR_FOR_DIVISION_BY_ZERO",
					"NO_ENGINE_SUBSTITUTION",
					"NO_ZERO_DATE",
					"NO_ZERO_IN_DATE",
					"STRICT_ALL_TABLES"
				],
				"mysql_require_primary_key": true,
				"mysql_slow_query_log": false,
				"cluster_time_zone": "America/New_York",
				"read_replicas": [
					{
						"id": "daeb6d62-a6a2-458c-9f74-e053735d7f50",
						"date_created": "2022-05-09 10:12:43",
						"plan": "vultr-dbaas-startup-cc-2-80-4",
						"plan_disk": 80,
						"plan_ram": 4096,
						"plan_vcpus": 2,
						"plan_replicas": 0,
						"region": "EWR",
						"status": "Running",
						"label": "testy-mc-testerton-the-7th",
						"tag": "",
						"database_engine": "mysql",
						"database_engine_version": "8",
						"dbname": "defaultdb",
						"host": "vultr-prod-87086a7d-4bc8-47ca-aa88-f88138d82772-vultr-prod-85e0.vultrdb.com",
						"user": "vultradmin",
						"password": "AVNS_UBen_MjqAwDd2BWFc-Y",
						"port": "16751",
						"maintenance_dow": "sunday",
						"maintenance_time": "02:00:00",
						"latest_backup": "2023-03-12 22:07:06",
						"trusted_ips": [],
						"mysql_sql_modes": [
							"ANSI",
							"ERROR_FOR_DIVISION_BY_ZERO",
							"NO_ENGINE_SUBSTITUTION",
							"NO_ZERO_DATE",
							"NO_ZERO_IN_DATE",
							"STRICT_ALL_TABLES"
						],
						"mysql_require_primary_key": true,
						"mysql_slow_query_log": false,
						"cluster_time_zone": "America/New_York"
					}
				]
			}
		}`
		fmt.Fprint(writer, response)
	})

	database, err := client.Database.Get(ctx, "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5")
	if err != nil {
		t.Errorf("Database.Get returned %+v", err)
	}

	mysql_sql_modes := []string{
		"ANSI",
		"ERROR_FOR_DIVISION_BY_ZERO",
		"NO_ENGINE_SUBSTITUTION",
		"NO_ZERO_DATE",
		"NO_ZERO_IN_DATE",
		"STRICT_ALL_TABLES",
	}

	replicas := []Database{
		{
			ID:                     "daeb6d62-a6a2-458c-9f74-e053735d7f50",
			DateCreated:            "2022-05-09 10:12:43",
			Plan:                   "vultr-dbaas-startup-cc-2-80-4",
			PlanDisk:               80,
			PlanRAM:                4096,
			PlanVCPUs:              2,
			PlanReplicas:           0,
			Region:                 "EWR",
			Status:                 "Running",
			Label:                  "testy-mc-testerton-the-7th",
			Tag:                    "",
			DatabaseEngine:         "mysql",
			DatabaseEngineVersion:  "8",
			DBName:                 "defaultdb",
			Host:                   "vultr-prod-87086a7d-4bc8-47ca-aa88-f88138d82772-vultr-prod-85e0.vultrdb.com",
			User:                   "vultradmin",
			Password:               "AVNS_UBen_MjqAwDd2BWFc-Y",
			Port:                   "16751",
			MaintenanceDOW:         "sunday",
			MaintenanceTime:        "02:00:00",
			LatestBackup:           "2023-03-12 22:07:06",
			TrustedIPs:             []string{},
			MySQLSQLModes:          mysql_sql_modes,
			MySQLRequirePrimaryKey: BoolToBoolPtr(true),
			MySQLSlowQueryLog:      BoolToBoolPtr(false),
			ClusterTimeZone:        "America/New_York",
		},
	}

	expected := &Database{
		ID:                     "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
		DateCreated:            "2022-05-09 10:13:31",
		Plan:                   "vultr-dbaas-business-cc-2-80-4",
		PlanDisk:               80,
		PlanRAM:                4096,
		PlanVCPUs:              2,
		PlanReplicas:           1,
		Region:                 "EWR",
		Status:                 "Running",
		Label:                  "testy-mc-testerton-the-8th",
		Tag:                    "bing bong",
		DatabaseEngine:         "mysql",
		DatabaseEngineVersion:  "8",
		DBName:                 "defaultdb",
		Host:                   "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
		User:                   "vultradmin",
		Password:               "AVNS_8E9hjx1LdsiA5EZ",
		Port:                   "16751",
		MaintenanceDOW:         "sunday",
		MaintenanceTime:        "02:00:00",
		LatestBackup:           "2023-03-13 00:59:07",
		TrustedIPs:             []string{},
		MySQLSQLModes:          mysql_sql_modes,
		MySQLRequirePrimaryKey: BoolToBoolPtr(true),
		MySQLSlowQueryLog:      BoolToBoolPtr(false),
		ClusterTimeZone:        "America/New_York",
		ReadReplicas:           replicas,
	}

	if !reflect.DeepEqual(database, expected) {
		t.Errorf("Database.Get returned %+v, expected %+v", database, expected)
	}
}

func TestDatabaseServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/databases/999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"database": {
				"id": "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
				"date_created": "2022-05-09 10:13:31",
				"plan": "vultr-dbaas-business-cc-2-80-4",
				"plan_disk": 80,
				"plan_ram": 4096,
				"plan_vcpus": 2,
				"plan_replicas": 1,
				"region": "EWR",
				"status": "Running",
				"label": "testy-mc-testerton-the-8th-part-2",
				"tag": "bing bong updated",
				"database_engine": "mysql",
				"database_engine_version": "8",
				"dbname": "defaultdb",
				"host": "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
				"user": "vultradmin",
				"password": "AVNS_8E9hjx1LdsiA5EZ",
				"port": "16751",
				"maintenance_dow": "sunday",
				"maintenance_time": "02:00:00",
				"latest_backup": "2023-03-13 00:59:07",
				"trusted_ips": [],
				"mysql_sql_modes": [
					"ANSI",
					"ERROR_FOR_DIVISION_BY_ZERO",
					"NO_ENGINE_SUBSTITUTION",
					"NO_ZERO_DATE",
					"NO_ZERO_IN_DATE",
					"STRICT_ALL_TABLES"
				],
				"mysql_require_primary_key": true,
				"mysql_slow_query_log": true,
				"mysql_long_query_time": 2,
				"cluster_time_zone": "America/New_York",
				"read_replicas": [
					{
						"id": "daeb6d62-a6a2-458c-9f74-e053735d7f50",
						"date_created": "2022-05-09 10:12:43",
						"plan": "vultr-dbaas-startup-cc-2-80-4",
						"plan_disk": 80,
						"plan_ram": 4096,
						"plan_vcpus": 2,
						"plan_replicas": 0,
						"region": "EWR",
						"status": "Running",
						"label": "testy-mc-testerton-the-7th",
						"tag": "",
						"database_engine": "mysql",
						"database_engine_version": "8",
						"dbname": "defaultdb",
						"host": "vultr-prod-87086a7d-4bc8-47ca-aa88-f88138d82772-vultr-prod-85e0.vultrdb.com",
						"user": "vultradmin",
						"password": "AVNS_UBen_MjqAwDd2BWFc-Y",
						"port": "16751",
						"maintenance_dow": "sunday",
						"maintenance_time": "02:00:00",
						"latest_backup": "2023-03-12 22:07:06",
						"trusted_ips": [],
						"mysql_sql_modes": [
							"ANSI",
							"ERROR_FOR_DIVISION_BY_ZERO",
							"NO_ENGINE_SUBSTITUTION",
							"NO_ZERO_DATE",
							"NO_ZERO_IN_DATE",
							"STRICT_ALL_TABLES"
						],
						"mysql_require_primary_key": true,
						"mysql_slow_query_log": true,
						"mysql_long_query_time": 2,
						"cluster_time_zone": "America/New_York"
					}
				]
			}
		}`
		fmt.Fprint(writer, response)
	})

	options := &DatabaseUpdateReq{
		Label:              "testy-mc-testerton-the-8th-part-2",
		Tag:                "bing bong updated",
		MySQLSlowQueryLog:  BoolToBoolPtr(true),
		MySQLLongQueryTime: 2,
	}

	database, err := client.Database.Update(ctx, "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", options)
	if err != nil {
		t.Errorf("Database.Update returned %+v", err)
	}

	mysql_sql_modes := []string{
		"ANSI",
		"ERROR_FOR_DIVISION_BY_ZERO",
		"NO_ENGINE_SUBSTITUTION",
		"NO_ZERO_DATE",
		"NO_ZERO_IN_DATE",
		"STRICT_ALL_TABLES",
	}

	replicas := []Database{
		{
			ID:                     "daeb6d62-a6a2-458c-9f74-e053735d7f50",
			DateCreated:            "2022-05-09 10:12:43",
			Plan:                   "vultr-dbaas-startup-cc-2-80-4",
			PlanDisk:               80,
			PlanRAM:                4096,
			PlanVCPUs:              2,
			PlanReplicas:           0,
			Region:                 "EWR",
			Status:                 "Running",
			Label:                  "testy-mc-testerton-the-7th",
			Tag:                    "",
			DatabaseEngine:         "mysql",
			DatabaseEngineVersion:  "8",
			DBName:                 "defaultdb",
			Host:                   "vultr-prod-87086a7d-4bc8-47ca-aa88-f88138d82772-vultr-prod-85e0.vultrdb.com",
			User:                   "vultradmin",
			Password:               "AVNS_UBen_MjqAwDd2BWFc-Y",
			Port:                   "16751",
			MaintenanceDOW:         "sunday",
			MaintenanceTime:        "02:00:00",
			LatestBackup:           "2023-03-12 22:07:06",
			TrustedIPs:             []string{},
			MySQLSQLModes:          mysql_sql_modes,
			MySQLRequirePrimaryKey: BoolToBoolPtr(true),
			MySQLSlowQueryLog:      BoolToBoolPtr(true),
			MySQLLongQueryTime:     2,
			ClusterTimeZone:        "America/New_York",
		},
	}

	expected := &Database{
		ID:                     "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5",
		DateCreated:            "2022-05-09 10:13:31",
		Plan:                   "vultr-dbaas-business-cc-2-80-4",
		PlanDisk:               80,
		PlanRAM:                4096,
		PlanVCPUs:              2,
		PlanReplicas:           1,
		Region:                 "EWR",
		Status:                 "Running",
		Label:                  "testy-mc-testerton-the-8th-part-2",
		Tag:                    "bing bong updated",
		DatabaseEngine:         "mysql",
		DatabaseEngineVersion:  "8",
		DBName:                 "defaultdb",
		Host:                   "vultr-prod-2db1ff4d-9d78-4baa-b52e-ec2b166107bb-vultr-prod-85e0.vultrdb.com",
		User:                   "vultradmin",
		Password:               "AVNS_8E9hjx1LdsiA5EZ",
		Port:                   "16751",
		MaintenanceDOW:         "sunday",
		MaintenanceTime:        "02:00:00",
		LatestBackup:           "2023-03-13 00:59:07",
		TrustedIPs:             []string{},
		MySQLSQLModes:          mysql_sql_modes,
		MySQLRequirePrimaryKey: BoolToBoolPtr(true),
		MySQLSlowQueryLog:      BoolToBoolPtr(true),
		MySQLLongQueryTime:     2,
		ClusterTimeZone:        "America/New_York",
		ReadReplicas:           replicas,
	}

	if !reflect.DeepEqual(database, expected) {
		t.Errorf("Database.Update returned %+v, expected %+v", database, expected)
	}
}

func TestDatabaseServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/databases/999c4ed0-f2e4-4f2a-a951-de358ceb9ab5", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Database.Delete(ctx, "999c4ed0-f2e4-4f2a-a951-de358ceb9ab5")

	if err != nil {
		t.Errorf("Database.Delete returned %+v", err)
	}
}

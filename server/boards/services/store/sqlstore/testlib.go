// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/mattermost/mattermost-server/v6/server/boards/model"
	"github.com/mattermost/mattermost-server/v6/server/boards/services/store"
	"github.com/mattermost/mattermost-server/v6/server/channels/store/storetest"
	"github.com/mattermost/mattermost-server/v6/server/platform/shared/mlog"
	"github.com/mgdelacroix/foundation"
	"github.com/stretchr/testify/require"
)

type storeType struct {
	Name       string
	ConnString string
	Store      store.Store
	Logger     *mlog.Logger
}

var mainStoreTypes []*storeType

func NewStoreType(name string, driver string, skipMigrations bool, overrideParams *Params) *storeType {
	settings := storetest.MakeSqlSettings(driver, false)
	connectionString := *settings.DataSource

	logger := mlog.CreateConsoleTestLogger(false, mlog.LvlDebug)

	sqlDB, err := sql.Open(driver, connectionString)
	if err != nil {
		panic(fmt.Sprintf("cannot open database: %s", err))
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("cannot ping database: %s", err))
	}

	storeParams := Params{
		DBType:           driver,
		ConnectionString: connectionString,
		SkipMigrations:   skipMigrations,
		TablePrefix:      "focalboard_",
		Logger:           logger,
		DB:               sqlDB,
		IsPlugin:         false, // ToDo: to be removed
	}

	if overrideParams != nil {
		if overrideParams.DBType != "" {
			storeParams.DBType = overrideParams.DBType
		}

		if overrideParams.ConnectionString != "" {
			storeParams.ConnectionString = overrideParams.ConnectionString
		}

		if overrideParams.SkipMigrations {
			storeParams.SkipMigrations = overrideParams.SkipMigrations
		}

		if overrideParams.TablePrefix != "" {
			storeParams.TablePrefix = overrideParams.TablePrefix
		}

		if overrideParams.IsPlugin {
			storeParams.IsPlugin = overrideParams.IsPlugin
		}
	}

	store, err := New(storeParams)
	if err != nil {
		panic(fmt.Sprintf("cannot create store: %s", err))
	}

	return &storeType{name, connectionString, store, logger}
}

func initStores(skipMigrations bool, overrideParams *Params) []*storeType {
	var storeTypes []*storeType

	if os.Getenv("IS_CI") == "true" {
		switch os.Getenv("MM_SQLSETTINGS_DRIVERNAME") {
		case "mysql":
			storeTypes = append(storeTypes, NewStoreType("MySQL", model.MysqlDBType, skipMigrations, overrideParams))
		case "postgres":
			storeTypes = append(storeTypes, NewStoreType("PostgreSQL", model.PostgresDBType, skipMigrations, overrideParams))
		}
	} else {
		storeTypes = append(storeTypes,
			NewStoreType("PostgreSQL", model.PostgresDBType, skipMigrations, overrideParams),
			NewStoreType("MySQL", model.MysqlDBType, skipMigrations, overrideParams),
		)
	}

	return storeTypes
}

func RunStoreTests(t *testing.T, f func(*testing.T, store.Store)) {
	for _, st := range mainStoreTypes {
		st := st
		require.NoError(t, st.Store.DropAllTables())
		t.Run(st.Name, func(t *testing.T) {
			f(t, st.Store)
		})
	}
}

func RunStoreTestsWithSqlStore(t *testing.T, f func(*testing.T, *SQLStore)) {
	RunStoreTestsWithCustomSqlStore(t, f, mainStoreTypes)
}

func RunStoreTestsWithCustomSqlStore(t *testing.T, f func(*testing.T, *SQLStore), stores []*storeType) {
	for _, st := range stores {
		st := st
		sqlstore := st.Store.(*SQLStore)
		require.NoError(t, sqlstore.DropAllTables())
		t.Run(st.Name, func(t *testing.T) {
			f(t, sqlstore)
		})
	}
}

// RunStoreTestsWithFoundation executes a test for all store types. It
// requires a new instance of each store type as migration tests
// cannot reuse old stores with already run migrations
func RunStoreTestsWithFoundation(t *testing.T, f func(*testing.T, *foundation.Foundation)) {
	storeTypes := initStores(true, nil)

	for _, st := range storeTypes {
		st := st
		t.Run(st.Name, func(t *testing.T) {
			sqlstore := st.Store.(*SQLStore)
			f(t, foundation.New(t, NewBoardsMigrator(sqlstore)))
		})
		require.NoError(t, st.Store.Shutdown())
		require.NoError(t, st.Logger.Shutdown())
	}
}

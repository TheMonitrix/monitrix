package signoz

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ezeslucky/monitrix/pkg/sqlstore"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/sqlstore/sqlstoretest"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/telemetrystore"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/telemetrystore/telemetrystoretest"
	"github.com/stretchr/testify/assert"
)

// This is a test to ensure that provider factories can be created without panicking since
// we are using the factory.MustNewNamedMap function to initialize the provider factories.
// It also helps us catch these errors during testing instead of runtime.
func TestNewProviderFactories(t *testing.T) {
	assert.NotPanics(t, func() {
		NewCacheProviderFactories()
	})

	assert.NotPanics(t, func() {
		NewWebProviderFactories()
	})

	assert.NotPanics(t, func() {
		NewSQLStoreProviderFactories()
	})

	assert.NotPanics(t, func() {
		NewTelemetryStoreProviderFactories()
	})

	assert.NotPanics(t, func() {
		NewSQLMigrationProviderFactories(sqlstoretest.New(sqlstore.Config{Provider: "sqlite"}, sqlmock.QueryMatcherEqual))
	})

	assert.NotPanics(t, func() {
		NewPrometheusProviderFactories(telemetrystoretest.New(telemetrystore.Config{Provider: "clickhouse"}, sqlmock.QueryMatcherEqual))
	})

	assert.NotPanics(t, func() {
		NewAlertmanagerProviderFactories(sqlstoretest.New(sqlstore.Config{Provider: "sqlite"}, sqlmock.QueryMatcherEqual))
	})
}

package db

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"

	"github.com/ezeslucky/monitrix/pkg/cache"
	"github.com/ezeslucky/monitrix/pkg/prometheus"
	basechr "github.com/ezeslucky/monitrix/pkg/query-service/app/clickhouseReader"
	"github.com/ezeslucky/monitrix/pkg/sqlstore"
	"github.com/ezeslucky/monitrix/pkg/telemetrystore"
)

type ClickhouseReader struct {
	conn  clickhouse.Conn
	appdb sqlstore.SQLStore
	*basechr.ClickHouseReader
}

func NewDataConnector(
	sqlDB sqlstore.SQLStore,
	telemetryStore telemetrystore.TelemetryStore,
	prometheus prometheus.Prometheus,
	cluster string,
	fluxIntervalForTraceDetail time.Duration,
	cache cache.Cache,
) *ClickhouseReader {
	chReader := basechr.NewReader(sqlDB, telemetryStore, prometheus, cluster, fluxIntervalForTraceDetail, cache)
	return &ClickhouseReader{
		conn:             telemetryStore.ClickhouseDB(),
		appdb:            sqlDB,
		ClickHouseReader: chReader,
	}
}

func (r *ClickhouseReader) GetSQLStore() sqlstore.SQLStore {
	return r.appdb
}

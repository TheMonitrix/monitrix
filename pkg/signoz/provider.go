package signoz

import (
	"github.com/ezeslucky/monitrix/pkg/alertmanager"
	"github.com/ezeslucky/monitrixitrix/pkg/alertmanager/legacyalertmanager"
	"github.com/ezeslucky/monitrixitrix/pkg/alertmanager/signozalertmanager"
	"github.com/ezeslucky/monitrixitrix/pkg/cache"
	"github.com/ezeslucky/monitrixitrix/pkg/cache/memorycache"
	"github.com/ezeslucky/monitrixitrix/pkg/cache/rediscache"
	"github.com/ezeslucky/monitrixitrix/pkg/factory"
	"github.com/ezeslucky/monitrixitrix/pkg/prometheus"
	"github.com/ezeslucky/monitrixitrix/pkg/prometheus/clickhouseprometheus"
	"github.com/ezeslucky/monitrixitrix/pkg/sqlmigration"
	"github.com/ezeslucky/monitrixitrix/pkg/sqlstore"
	"github.com/ezeslucky/monitrixitrix/pkg/sqlstore/sqlitesqlstore"
	"github.com/ezeslucky/monitrixitrix/pkg/sqlstore/sqlstorehook"
	"github.com/ezeslucky/monitrixitrix/pkg/telemetrystore"
	"github.com/ezeslucky/monitrixitrix/pkg/telemetrystore/clickhousetelemetrystore"
	"github.com/ezeslucky/monitrixitrix/pkg/telemetrystore/telemetrystorehook"
	"github.com/ezeslucky/monitrixitrix/pkg/web"
	"github.com/ezeslucky/monitrixitrix/pkg/web/noopweb"
	"github.com/ezeslucky/monitrixitrix/pkg/web/routerweb"
)

func NewCacheProviderFactories() factory.NamedMap[factory.ProviderFactory[cache.Cache, cache.Config]] {
	return factory.MustNewNamedMap(
		memorycache.NewFactory(),
		rediscache.NewFactory(),
	)
}

func NewWebProviderFactories() factory.NamedMap[factory.ProviderFactory[web.Web, web.Config]] {
	return factory.MustNewNamedMap(
		routerweb.NewFactory(),
		noopweb.NewFactory(),
	)
}

func NewSQLStoreProviderFactories() factory.NamedMap[factory.ProviderFactory[sqlstore.SQLStore, sqlstore.Config]] {
	hook := sqlstorehook.NewLoggingFactory()
	return factory.MustNewNamedMap(
		sqlitesqlstore.NewFactory(hook),
	)
}

func NewSQLMigrationProviderFactories(sqlstore sqlstore.SQLStore) factory.NamedMap[factory.ProviderFactory[sqlmigration.SQLMigration, sqlmigration.Config]] {
	return factory.MustNewNamedMap(
		sqlmigration.NewAddDataMigrationsFactory(),
		sqlmigration.NewAddOrganizationFactory(),
		sqlmigration.NewAddPreferencesFactory(),
		sqlmigration.NewAddDashboardsFactory(),
		sqlmigration.NewAddSavedViewsFactory(),
		sqlmigration.NewAddAgentsFactory(),
		sqlmigration.NewAddPipelinesFactory(),
		sqlmigration.NewAddIntegrationsFactory(),
		sqlmigration.NewAddLicensesFactory(),
		sqlmigration.NewAddPatsFactory(),
		sqlmigration.NewModifyDatetimeFactory(),
		sqlmigration.NewModifyOrgDomainFactory(),
		sqlmigration.NewUpdateOrganizationFactory(sqlstore),
		sqlmigration.NewAddAlertmanagerFactory(sqlstore),
		sqlmigration.NewUpdateDashboardAndSavedViewsFactory(sqlstore),
		sqlmigration.NewUpdatePatAndOrgDomainsFactory(sqlstore),
		sqlmigration.NewUpdatePipelines(sqlstore),
		sqlmigration.NewDropLicensesSitesFactory(sqlstore),
		sqlmigration.NewUpdateInvitesFactory(sqlstore),
		sqlmigration.NewUpdatePatFactory(sqlstore),
		sqlmigration.NewUpdateAlertmanagerFactory(sqlstore),
		sqlmigration.NewUpdatePreferencesFactory(sqlstore),
		sqlmigration.NewUpdateApdexTtlFactory(sqlstore),
		sqlmigration.NewUpdateResetPasswordFactory(sqlstore),
		sqlmigration.NewUpdateRulesFactory(sqlstore),
		sqlmigration.NewAddVirtualFieldsFactory(),
		sqlmigration.NewUpdateIntegrationsFactory(sqlstore),
		sqlmigration.NewUpdateOrganizationsFactory(sqlstore),
		sqlmigration.NewDropGroupsFactory(sqlstore),
		sqlmigration.NewCreateQuickFiltersFactory(sqlstore),
		sqlmigration.NewUpdateQuickFiltersFactory(sqlstore),
		sqlmigration.NewAuthRefactorFactory(sqlstore),
	)
}

func NewTelemetryStoreProviderFactories() factory.NamedMap[factory.ProviderFactory[telemetrystore.TelemetryStore, telemetrystore.Config]] {
	return factory.MustNewNamedMap(
		clickhousetelemetrystore.NewFactory(telemetrystorehook.NewSettingsFactory(), telemetrystorehook.NewLoggingFactory()),
	)
}

func NewPrometheusProviderFactories(telemetryStore telemetrystore.TelemetryStore) factory.NamedMap[factory.ProviderFactory[prometheus.Prometheus, prometheus.Config]] {
	return factory.MustNewNamedMap(
		clickhouseprometheus.NewFactory(telemetryStore),
	)
}

func NewAlertmanagerProviderFactories(sqlstore sqlstore.SQLStore) factory.NamedMap[factory.ProviderFactory[alertmanager.Alertmanager, alertmanager.Config]] {
	return factory.MustNewNamedMap(
		legacyalertmanager.NewFactory(sqlstore),
		signozalertmanager.NewFactory(sqlstore),
	)
}

package signoz

import (
	"github.com/ezeslucky/monitrix/pkg/alertmanager"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/alertmanager/legacyalertmanager"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/alertmanager/signozalertmanager"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/cache"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/cache/memorycache"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/cache/rediscache"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/factory"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/prometheus"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/prometheus/clickhouseprometheus"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/sqlmigration"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/sqlstore"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/sqlstore/sqlitesqlstore"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/sqlstore/sqlstorehook"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/telemetrystore"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/telemetrystore/clickhousetelemetrystore"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/telemetrystore/telemetrystorehook"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/web"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/web/noopweb"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/web/routerweb"
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

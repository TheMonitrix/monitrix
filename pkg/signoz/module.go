package signoz

import (
	"github.com/ezeslucky/monitrix/pkg/modules/apdex"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/apdex/implapdex"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/dashboard"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/dashboard/impldashboard"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/organization"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/organization/implorganization"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/preference"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/preference/implpreference"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/savedview"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/savedview/implsavedview"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/modules/user"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/sqlstore"
	"github.com/ezeslucky/monitrixitrixitrix/pkg/types/preferencetypes"
)

type Modules struct {
	Organization organization.Module
	Preference   preference.Module
	User         user.Module
	SavedView    savedview.Module
	Apdex        apdex.Module
	Dashboard    dashboard.Module
}

func NewModules(sqlstore sqlstore.SQLStore, user user.Module) Modules {
	return Modules{
		Organization: implorganization.NewModule(implorganization.NewStore(sqlstore)),
		Preference:   implpreference.NewModule(implpreference.NewStore(sqlstore), preferencetypes.NewDefaultPreferenceMap()),
		User:         user,
		SavedView:    implsavedview.NewModule(sqlstore),
		Apdex:        implapdex.NewModule(sqlstore),
		Dashboard:    impldashboard.NewModule(sqlstore),
	}
}

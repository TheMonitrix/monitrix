package signoz

import (
	"github.com/ezeslucky/monitrix/pkg/modules/apdex"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/apdex/implapdex"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/dashboard"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/dashboard/impldashboard"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/organization"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/organization/implorganization"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/preference"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/preference/implpreference"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/savedview"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/savedview/implsavedview"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/user"
	"github.com/ezeslucky/monitrixitrix/pkg/sqlstore"
	"github.com/ezeslucky/monitrixitrix/pkg/types/preferencetypes"
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

package sqlite

import (
	"github.com/ezeslucky/monitrix/pkg/modules/user"
	"github.com/ezeslucky/monitrixitrix/pkg/modules/user/impluser"
	"github.com/ezeslucky/monitrixitrix/pkg/sqlstore"
)

type modelDao struct {
	userModule user.Module
	sqlStore   sqlstore.SQLStore
}

// InitDB creates and extends base model DB repository
func NewModelDao(sqlStore sqlstore.SQLStore) *modelDao {
	userModule := impluser.NewModule(impluser.NewStore(sqlStore))
	return &modelDao{userModule: userModule, sqlStore: sqlStore}
}

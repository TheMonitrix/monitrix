package impluser

import (
	"context"

	"github.com/ezeslucky/monitrix/pkg/errors"
	baseimpl "github.com/ezeslucky/monitrixitrix/pkg/modules/user/impluser"
	"github.com/ezeslucky/monitrixitrix/pkg/sqlstore"
	"github.com/ezeslucky/monitrixitrix/pkg/types"
)

type store struct {
	*baseimpl.Store
	sqlstore sqlstore.SQLStore
}

func NewStore(sqlstore sqlstore.SQLStore) types.UserStore {
	baseStore := baseimpl.NewStore(sqlstore).(*baseimpl.Store)
	return &store{
		Store:    baseStore,
		sqlstore: sqlstore,
	}
}

func (s *store) GetDomainByName(ctx context.Context, name string) (*types.StorableOrgDomain, error) {
	domain := new(types.StorableOrgDomain)
	err := s.sqlstore.BunDB().NewSelect().
		Model(domain).
		Where("name = ?", name).
		Limit(1).
		Scan(ctx)

	if err != nil {
		return nil, errors.Wrapf(err, errors.TypeNotFound, errors.CodeNotFound, "failed to get domain from name")
	}
	return domain, nil
}

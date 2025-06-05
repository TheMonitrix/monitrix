package querybuildertypesv5

import "github.com/ezeslucky/monitrix/pkg/valuer"

type QueryType struct {
	valuer.String
}

var (
	QueryTypeUnknown       = QueryType{valuer.NewString("unknown")}
	QueryTypeBuilder       = QueryType{valuer.NewString("builder_query")}
	QueryTypeFormula       = QueryType{valuer.NewString("builder_formula")}
	QueryTypeSubQuery      = QueryType{valuer.NewString("builder_sub_query")}
	QueryTypeJoin          = QueryType{valuer.NewString("builder_join")}
	QueryTypeClickHouseSQL = QueryType{valuer.NewString("clickhouse_sql")}
	QueryTypePromQL        = QueryType{valuer.NewString("promql")}
)

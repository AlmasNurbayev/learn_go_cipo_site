package productFilter

import (
	"cipo_cite_server/internal/utils"
	"cipo_cite_server/internal/utils/filter"
	"net/url"
	"strconv"
)

func Filters(queries url.Values) (*map[string]interface{}, error) {
	f := filter.New(queries)
	// if queries.Has("news") {
	// 	f.Search = true
	// }
	f.DisablePaging = true

	var idInt int64
	idString := queries.Get("id")
	if idString != "" {
		idConv, err := strconv.Atoi(idString)
		if err != nil {
			return nil, err
		}
		idInt = int64(idConv)
	}

	name1cString := queries.Get("name_1c")

	filterRes := map[string]interface{}{
		"Base":             *f,
		"products.id":      utils.Ternary(idInt != 0, idInt, nil), // так надо для запроса
		"products.name_1c": utils.Ternary(name1cString != "", name1cString, nil),
	}

	return &filterRes, nil
}

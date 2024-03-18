package productFilter

import (
	"cipo_cite_server/internal/utils"
	"cipo_cite_server/internal/utils/filter"
	"net/url"
	"strconv"
	"strings"
)

type FilterListT struct {
	Filters map[string]interface{}
	Base    *filter.Filter `json:"base"`
}

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

func FiltersList(queries url.Values) (*FilterListT, error) {
	f := filter.New(queries)
	filters := make(map[string]interface{})
	res := FilterListT{
		Filters: filters,
		Base:    f,
	}

	vidString := queries.Get("vid_modeli")
	if vidString != "" {
		vidRes, err := utils.StringToArrInt64(vidString)
		if err != nil {
			return nil, err
		}
		if len(vidRes) > 0 {
			res.Filters["Vid_modeli"] = &vidRes
		}
	}

	sizeString := queries.Get("size")
	if sizeString != "" {
		sizeRes, err := utils.StringToArrInt64(sizeString)
		if err != nil {
			return nil, err
		}
		if len(sizeRes) > 0 {
			res.Filters["size"] = &sizeRes
		}
	}

	pgString := queries.Get("product_group")
	if pgString != "" {
		pgRes, err := utils.StringToArrInt64(pgString)
		if err != nil {
			return nil, err
		}
		if len(pgRes) > 0 {
			res.Filters["product_group"] = &pgRes
		}
	}

	minString := queries.Get("minPrice")
	if minString != "" {
		min, err := strconv.Atoi(strings.TrimSpace(minString))
		if err != nil {
			return nil, err
		}
		res.Filters["minPrice"] = &min
	}

	maxString := queries.Get("maxPrice")
	if maxString != "" {
		max, err := strconv.Atoi(strings.TrimSpace(maxString))
		if err != nil {
			return nil, err
		}
		res.Filters["maxPrice"] = &max
	}

	searchNameString := queries.Get("search_name")
	if searchNameString != "" {
		trim := "%" + strings.TrimSpace(searchNameString) + "%"
		res.Filters["search_name"] = &trim
	}

	//utils.PrintAsJSON(res)

	return &res, nil
}

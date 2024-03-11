package filter

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	paginationDefaultPage = 1
	paginationDefaultSize = 30

	queryParamPage          = "page"
	queryParamLimit         = "limit"
	queryParamSkip          = "skip"
	queryParamDisablePaging = "disable_paging"
	queryParamSort          = "sort"
)

type Filter struct {
	Page          int
	Skip          int
	Limit         int
	DisablePaging bool

	Sort   map[string]string
	Search bool
}

func New(queries url.Values) *Filter {
	var page, limit, skip int
	page, err := strconv.Atoi(queries.Get(queryParamPage))
	if err != nil {
		page = paginationDefaultPage
	}
	limit, err = strconv.Atoi(queries.Get(queryParamLimit))
	if err != nil {
		limit = paginationDefaultSize
	}

	skip, err = strconv.Atoi(queries.Get(queryParamSkip))
	if err != nil {
		skip = limit * (page - 1) // calculates offset
	}

	disablePaging, _ := strconv.ParseBool(queries.Get(queryParamDisablePaging))

	sortKey := make(map[string]string)
	if queries.Has(queryParamSort) {
		s := queries[queryParamSort]
		for _, val := range s {
			key, order, found := strings.Cut(val, ",")
			if found {
				sortKey[key] = strings.ToUpper(order)
			} else {
				sortKey[key] = "ASC"
			}
		}
	}

	return &Filter{
		Page:          page,
		Skip:          skip,
		Limit:         limit,
		DisablePaging: disablePaging,
		Sort:          sortKey,
	}
}

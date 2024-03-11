package newsFilter

import (
	"cipo_cite_server/internal/utils/filter"
	"net/url"
)

type Filter struct {
	Base filter.Filter
	News string `json:"news" query:"news" validate:"required, dive"`
}

func Filters(queries url.Values) *Filter {
	f := filter.New(queries)
	if queries.Has("news") {
		f.Search = true
	}
	return &Filter{
		Base: *f,

		News: queries.Get("news"),
	}
}

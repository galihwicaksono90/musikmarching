package utils

import (
	"net/http"
	"strconv"
)

var DEFAULT_PAGINATON_LIMIT = 10
var DEFAULT_PAGINATON_OFFSET = 0

func ParsePagination(r *http.Request) (int32, int32) {
	limit, err := strconv.Atoi(r.FormValue("page_limit"))
	if err != nil {
		limit = DEFAULT_PAGINATON_LIMIT
	}

	offset, err := strconv.Atoi(r.FormValue("page_offset"))
	if err != nil {
		offset = DEFAULT_PAGINATON_OFFSET
	}

	return int32(limit), int32(offset)
}

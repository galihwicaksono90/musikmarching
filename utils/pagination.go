package utils

import (
	"net/http"
	"strconv"
)

var DEFAULT_PAGINATON_LIMIT = 10
var DEFAULT_PAGINATON_OFFSET = 0

func ParsePagination(r *http.Request) (int32, int32) {
	pageLimit, err := strconv.Atoi(r.FormValue("page_limit"))
	if err != nil {
		pageLimit = DEFAULT_PAGINATON_LIMIT
	}

	pageOffset, err := strconv.Atoi(r.FormValue("page_offset"))
	if err != nil {
		pageOffset = DEFAULT_PAGINATON_OFFSET
	}

	return int32(pageLimit), int32(pageOffset)
}

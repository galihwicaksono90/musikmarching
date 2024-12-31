package utils

import (
	"net/url"
	"strconv"
)

var DEFAULT_PAGINATON_LIMIT = 10
var DEFAULT_PAGINATON_OFFSET = 0

func ParsePagination(values url.Values) (int32, int32) {
	limit, err := strconv.Atoi(values.Get("limit"))
	if err != nil {
		limit = DEFAULT_PAGINATON_LIMIT
	}

	offset, err := strconv.Atoi(values.Get("page_offset"))
	if err != nil {
		offset = DEFAULT_PAGINATON_OFFSET
	}

	return int32(limit), int32(offset)
}

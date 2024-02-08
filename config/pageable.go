package config

import (
	"net/url"
	"strconv"
)

type Pageable struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func NewPageable(url *url.URL) Pageable {
	result := Pageable{
		Page: 0,
		Size: 10,
	}

	if url.Query().Get("page") != "" {
		page, err := strconv.Atoi(url.Query().Get("page"))
		if err == nil {
			result.Page = page
		}
	}

	if url.Query().Get("size") != "" {
		size, err := strconv.Atoi(url.Query().Get("size"))
		if err == nil {
			result.Size = size
		}
	}

	return result
}

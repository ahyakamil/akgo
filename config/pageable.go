package config

type Pageable struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func NewPageable() Pageable {
	return Pageable{
		Page: 0,
		Size: 10,
	}
}

package goshiki

type Studio struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	FilteredName string `json:"filtered_name"`
	IsReal       bool   `json:"real"`
}

package models

//easyjson:json
type Updated struct {
	UpdatedInfo bool `json:"updated"`
}

//easyjson:json
type Deleted struct {
	DeletedInfo bool `json:"deleted"`
}

//easyjson:json
type ImgBoard struct {
	ImgPath string `json:"img_board"`
}

//easyjson:json
type Appended struct {
	AppendedInfo bool `json:"appended"`
}

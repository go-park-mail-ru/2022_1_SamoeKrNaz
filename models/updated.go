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

//easyjson:json
type Is_okayIn struct {
	Is_okayInfo bool `json:"Is_okay"`
}

//easyjson:json
type Avatar struct {
	AvatarPath string `json:"avatar_path"`
}

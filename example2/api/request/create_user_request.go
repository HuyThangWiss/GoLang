package request


type Books struct {
	Id string `json:"Id" validate:"required"`
	Title string `json:"Title" validate:"required"`
	Description string `json:"Description" validate:"required"`
}


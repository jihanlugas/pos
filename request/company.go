package request

type CreateCompany struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Address     string `json:"address" validate:""`
}

type UpdateCompany struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Address     string `json:"address" validate:""`
}

type PageCompany struct {
	Paging
	Name string `json:"name" form:"name" query:"name"`
}

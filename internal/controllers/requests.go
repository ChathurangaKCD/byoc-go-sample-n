package controllers

type AddCategoryRequest struct {
	Name string `json:"name,omitempty"`
}

type AddPetRequest struct {
	Name       string `json:"name,omitempty"`
	CategoryId string `json:"category_id,omitempty"`
	Available  bool   `json:"available"`
}

type UpdatePetRequest struct {
	Name       *string `json:"name,omitempty"`
	CategoryId *string `json:"category_id,omitempty"`
	Available  *bool   `json:"available"`
}

package dto

type GetCurrentUserRequest struct {
}

type GetCurrentUserResponse struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Preference string `json:"preference"`
	WeightUnit string `json:"weightUnit"`
	HeightUnit string `json:"heightUnit"`
	Weight     int    `json:"weight"`
	Height     int    `json:"height"`
	ImageURI   string `json:"imageUri"`
}
type UpdateUserRequest struct {
	Email           *string `json:"email" validate:"email"`
	Name            *string `json:"name" validate:"min=4,max=52,ascii"`
	UserImageUri    *string `json:"userImageUri" validate:"url"`
	CompanyName     *string `json:"companyName" validate:"min=4,max=52,ascii"`
	CompanyImageUri *string `json:"companyImageUri" validate:"url"`
}

type UpdateUserResponse struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

type UserProfile struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

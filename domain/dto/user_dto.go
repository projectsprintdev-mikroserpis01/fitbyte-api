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
	Name       *string `json:"name" validate:"min=2,max=60,ascii"`
	Preference string  `json:"preference" validate:"required,ascii,oneof=CARDIO WEIGHT"`
	WeightUnit string  `json:"weightUnit" validate:"required,ascii,oneof=KG LBS"`
	HeightUnit string  `json:"heightUnit" validate:"required,ascii,oneof=CM INCH"`
	Weight     int     `json:"weight" validate:"required,min=10,max=1000"`
	Height     int     `json:"height" validate:"required,min=3,max=250"`
	ImageURI   *string `json:"imageUri" validate:"url"`
}

type UpdateUserResponse struct {
	Name       *string `json:"name"`
	Preference string  `json:"preference"`
	WeightUnit string  `json:"weightUnit"`
	HeightUnit string  `json:"heightUnit"`
	Weight     int     `json:"weight"`
	Height     int     `json:"height"`
	ImageURI   *string `json:"imageUri"`
}

type UserProfile struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

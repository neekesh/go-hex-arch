package adapter_entities

type QueryParams struct {
	FullName string `form:"full_name"`
}

type User struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

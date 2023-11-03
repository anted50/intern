package models

type User struct {
	// ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegistrationJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationStatus struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Username string `json:"User"`
}

type LoginJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginStatus struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Username string `json:"User"`
}

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

type Trip struct {
	ID          int    `json:"id"`
	Title       string `json:"Title"`
	Desc        string `json:"Description"`
	Url         string `json:"URL"`
	Price       string `json:"Price"`
	Capacity    int    `json:"Capacity"`
	Type        string `json:"Type"`
	HasCompFood bool   `json:"hasCompFood"`
	Additional  string `json:"additional"`
	ContactNo   string `json:"ContactNo"`
}

type Get struct {
	ID int `json:"id"`
}

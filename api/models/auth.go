package models

type SignUpModel struct {
	FirstName     string `json:"first_name"`
	PhoneNumber   string `json:"phone_number"`
	OTPCode       string `json:"code"`
	PassCodeToken string `json:"passcode_token"`
}

type User struct {
	ID          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	PhoneNumber string  `json:"phone_number"`
	UserImg     *string `json:"user_img"`
}

type SignInModel struct {
	PhoneNumber   string `json:"phone_number"`
	PassCodeToken string `json:"passcode_token"`
	OTPCode       string `json:"code"`
}

type UpdateUser struct {
	FirstName   string  `json:"first_name"`
	PhoneNumber string  `json:"phone_number"`
	UserImg     *string `json:"user_img"`
}

// Username- Samandar
// PhoneNumber

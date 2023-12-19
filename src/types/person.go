package types

type Person struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phoneNumber"`
}

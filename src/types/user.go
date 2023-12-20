package types

type User struct {
	Username    *string `json:"username" required:"true"`
	Password    *string `json:"password" required:"true"`
	Email       *string `json:"email" required:"true"`
	FirstName   *string `json:"firstName" required:"true"`
	LastName    *string `json:"lastName" required:"true"`
	PhoneNumber *string `json:"phoneNumber" required:"true"`
	Pet         *string `json:"pet"`
}

package dtos

type UserType struct {
}

type CreateUserRequest struct {
	Name  string `valid:"Required;" json:"name"`
	Phone string `valid:"Required;MinSize(11);MaxSize(11)" json:"phone"`
	Email string `json:"email"`
	Type  int    `valid:"Required;" json:"type"`
}

type CreateUserResponse struct {
}

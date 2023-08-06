// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type CreateUserInput struct {
	PhoneNumber string `db:"phone_number"`
	FullName    string `db:"full_name"`
	Password    []byte `db:"password"`
}

type CreateUserOutput struct {
	Id int64 `db:"id"`
}

type UpdateSuccessfulLoginCountInput struct {
	Id int64 `db:"id"`
}

type GetUserByPhoneNumberInput struct {
	PhoneNumber string `db:"phone_number"`
}

type GetUserByIdInput struct {
	Id int64 `db:"id"`
}

type Users struct {
	Id          int64  `pq:"id"`
	PhoneNumber string `pq:"phone_number"`
	FullName    string `pq:"full_name"`
	Password    []byte `db:"password"`
}

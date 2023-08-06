package repository

import (
	"context"
	"database/sql"

	errorlib "github.com/fazarrahman/user-profile/errorLib"
)

func (r *Repository) CreateUser(ctx context.Context, input CreateUserInput) (output *CreateUserOutput, errl *errorlib.Error) {
	var lastInsertId int64
	err := r.Db.QueryRowContext(ctx,
		"INSERT INTO users (phone_number, full_name, passwords) VALUES ($1, $2, $3)  RETURNING id",
		input.PhoneNumber,
		input.FullName, input.Password).Scan(&lastInsertId)
	if err != nil {
		return nil, errorlib.InternalServerError("Error when inserting user : " + err.Error())
	}
	return &CreateUserOutput{Id: lastInsertId}, nil
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, input GetUserByPhoneNumberInput) (*GetUserByPhoneNumberOutput, *errorlib.Error) {
	var output GetUserByPhoneNumberOutput
	err := r.Db.QueryRowContext(ctx, "SELECT id, phone_number, full_name FROM users WHERE phone_number = $1", input.PhoneNumber).Scan(&output.Id, &output.PhoneNumber, &output.FullName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errorlib.InternalServerError("Error when getting user by phone : " + err.Error())
	}
	return &output, nil
}

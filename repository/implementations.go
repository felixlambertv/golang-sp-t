package repository

import (
	"context"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateUser(ctx context.Context, input CreateUserInput) (output *CreateUserOutput, err error) {
	output = &CreateUserOutput{}
	query := `INSERT INTO users (phone_number, full_name, password) VALUES ($1, $2, $3) RETURNING id`
	err = r.Db.QueryRowContext(ctx, query, input.PhoneNumber, input.FullName, input.Password).Scan(&output.Id)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (r *Repository) IsPhoneNumberExists(ctx context.Context, phoneNumber string) (isExists bool, err error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = $1)"

	err = r.Db.QueryRowContext(ctx, query, phoneNumber).Scan(&isExists)
	if err != nil {
		return false, err
	}
	return isExists, nil
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (output *UserOutput, err error) {
	output = &UserOutput{}
	query := "SELECT id, phone_number, full_name, password FROM users WHERE phone_number = $1"
	err = r.Db.QueryRowContext(ctx, query, phoneNumber).Scan(&output.Id, &output.PhoneNumber, &output.FullName, &output.Password)
	if err != nil {
		return nil, err
	}
	return output, err
}

func (r *Repository) GetUserById(ctx context.Context, id uint) (output *UserOutput, err error) {
	output = &UserOutput{}
	query := "SELECT id, phone_number, full_name, password FROM users WHERE id = $1"
	err = r.Db.QueryRowContext(ctx, query, id).Scan(&output.Id, &output.PhoneNumber, &output.FullName, &output.Password)
	if err != nil {
		return nil, err
	}
	return output, err
}

func (r *Repository) IncrementUserLoginCount(ctx context.Context, id uint) error {
	query := "UPDATE users SET login_count = login_count + 1 WHERE id = $1"
	_, err := r.Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

package models

import "time"

type User struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
}

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// CalculateAge returns the age in full years based on the given date of birth.
func CalculateAge(dob time.Time, now time.Time) int {
	year, month, day := dob.Date()
	ny, nm, nd := now.Date()

	age := ny - year

	// If birthday has not occurred yet this year, subtract 1.
	if nm < month || (nm == month && nd < day) {
		age--
	}
	if age < 0 {
		return 0
	}
	return age
}



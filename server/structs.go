package main

// User is Interface for user details.
type User struct {
	ID      int
	Name    string
	Lname   string
	Country string
}

// Employee is Interface for employee details.
type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// ErrorResponse is interface for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}

package models

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

type CreateEmployeeRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Division string `json:"division" binding:"required"`
}

type UpdateEmployeeRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
	Email    string `json:"email"`
}

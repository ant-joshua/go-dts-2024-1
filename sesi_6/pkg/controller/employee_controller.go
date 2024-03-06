package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

type EmployeeController struct {
	DB *sql.DB
}

func (c *EmployeeController) Routes(r *gin.RouterGroup) {
	routeGroup := r.Group("/employees")

	routeGroup.GET("", c.GetEmployeeList)
	routeGroup.POST("", c.CreateEmployee)
	routeGroup.GET("/:id", c.GetEmployee)
	routeGroup.PUT("/:id", c.UpdateEmployee)
	routeGroup.DELETE("/:id", c.DeleteEmployee)
}

// GetEmployee godoc
// @Summary      Get All Employe List
// @Description  Get All Employe List
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Success      200  {object}  Employee
// @Response      404  {object}  HttpError false "error 404" example({"code": 404, "message": "Employee not found"})
// @Response      500  {object}  HttpError false "error 500" example({"code": 500, "message": "Internal Server Error"})
// @Router       /api/employees [get]
func (c *EmployeeController) GetEmployeeList(ctx *gin.Context) {
	sqlStatement := `SELECT * FROM employee`

	rows, err := c.DB.Query(sqlStatement)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var employees []Employee

	for rows.Next() {
		var employee Employee

		err = rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		employees = append(employees, employee)
	}

	ctx.JSON(http.StatusOK, employees)
}

// GetEmployee godoc
// @Summary      Get Employe By ID
// @Description  Get Employe By ID
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Employee ID"
// @Success      200  {object}  Employee
// @Router       /api/employees/{id} [get]
func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	sqlStatement := `SELECT * FROM employee where employee_id = $1`

	var employee Employee

	err := c.DB.QueryRow(sqlStatement, id).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Employee with id %s not found", id)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// CreateEmployee godoc
// @Summary      Crete Employe
// @Description  Crete Employe
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param 		 request  body Employee  true  "Employee Data"
// @Success      200  {object}  Employee
// @Router       /api/employees [post]
func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var request CreateEmployeeRequest

	var employee Employee

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO employee (full_name, email, age, division) VALUES ($1, $2, $3, $4) returning *`
	//
	err = c.DB.QueryRow(sqlStatement, request.Name, request.Email, request.Age, request.Division).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	var request UpdateEmployeeRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE employee SET full_name = $1, age = $2, division = $3, email = $4 WHERE employee_id = $5 returning *`

	var employee Employee

	err = c.DB.QueryRow(sqlStatement, request.Name, request.Age, request.Division, request.Email, id).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	// for i, emp := range employees {
	// 	if emp.ID == id {
	// 		// employees = append(employees[:i], employees[i+1:]...)

	// 		employees = slices.Delete(employees, i, 1)

	// 		ctx.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
	// 		return
	// 	}
	// }

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

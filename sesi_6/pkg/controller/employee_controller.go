package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"sesi_6/pkg/models"

	"github.com/gin-gonic/gin"
)

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

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee

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

	var employee models.Employee

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
	var request models.CreateEmployeeRequest

	var employee models.Employee

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

	var request models.UpdateEmployeeRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE employee SET full_name = $1, age = $2, division = $3, email = $4 WHERE employee_id = $5 returning *`

	var employee models.Employee

	err = c.DB.QueryRow(sqlStatement, request.Name, request.Age, request.Division, request.Email, id).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	sqlStatement := `DELETE FROM employee WHERE employee_id = $1`

	result, err := c.DB.Exec(sqlStatement, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	count, err := result.RowsAffected()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Employee with id %s not found", id)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Employee with id %s has been deleted", id)})

}

package controller

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

var employees = []Employee{
	{ID: "123e4567-e89b-12d3-a456-426614174000", Name: "John Doe", Age: 25, Division: "Engineering"},
	{ID: "123e4567-e89b-12d3-a456-426614174001", Name: "Jane Doe", Age: 30, Division: "Marketing"},
}

type EmployeeController struct {
}

func (c *EmployeeController) Routes(r *gin.RouterGroup) {
	routeGroup := r.Group("/employees")

	routeGroup.GET("", c.GetEmployeeList)
	routeGroup.POST("", c.CreateEmployee)
	routeGroup.GET("/:id", c.GetEmployee)
	routeGroup.PUT("/:id", c.UpdateEmployee)
	routeGroup.DELETE("/:id", c.DeleteEmployee)
}

func (c *EmployeeController) GetEmployeeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, employees)
}

func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, employee := range employees {
		if employee.ID == id {
			ctx.JSON(http.StatusOK, employee)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var employee Employee

	err := ctx.ShouldBindJSON(&employee)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUUID := uuid.New()

	employee.ID = newUUID.String()

	employees = append(employees, employee)

	ctx.JSON(http.StatusCreated, employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var employee Employee

	employee.ID = id

	err = ctx.ShouldBindJSON(&employee)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			employees[i] = employee
			ctx.JSON(http.StatusOK, employee)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			// employees = append(employees[:i], employees[i+1:]...)

			employees = slices.Delete(employees, i, 1)

			ctx.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

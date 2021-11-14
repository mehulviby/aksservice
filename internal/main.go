package main

import (
	"aksservice/internal/api"
	"os"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := SetupEcho()

	controller := &api.Controller{}

	// Route => handler
	e.POST("/aksvalidator", controller.AksValidator)
	e.GET("/akssearch", controller.AksSearch)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func SetupEcho() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Logger.SetOutput(os.Stderr)
	e.Use(middleware.Recover())
	validate := validator.New()
	_ = validate.RegisterValidation("regularExpression", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(fl.Param())
		return re.MatchString(fl.Field().String())
	})
	e.Validator = &Validator{validator: validate}
	return e
}

// Validator is implementation of validation of rquest values.
type Validator struct {
	validator *validator.Validate
}

// Validate do validation for request value.
func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}
	return err.(validator.ValidationErrors)
}

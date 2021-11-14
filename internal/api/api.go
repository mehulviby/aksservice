package api

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/yaml.v2"
)

type Maintainers struct {
	Name  string `yaml:"name" validate:"required"`
	Email string `yaml:"email" validate:"required,email"`
}

type AksConfig struct {
	Title       string        `yaml:"title" validate:"required"`
	Version     string        `yaml:"version" validate:"required,regularExpression=^[0-9]+.[0-9]+.[0-9]+$"`
	Maintainers []Maintainers `yaml:"maintainers" validate:"dive,required"`
	Company     string        `yaml:"company" validate:"required"`
	Website     string        `yaml:"website" validate:"required,url"`
	Source      string        `yaml:"source" validate:"required,url"`
	License     string        `yaml:"license" validate:"required"`
	Description string        `yaml:"description" validate:"required"`
}

// Controller is a controller for this application.
type Controller struct {
	AksConfigList []AksConfig
}

// AksValidator handles incoming validator requests
func (controller *Controller) AksValidator(ctx echo.Context) error {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return echo.ErrBadRequest
	}
	var aksConfig AksConfig
	if err := yaml.Unmarshal(body, &aksConfig); err != nil {
		return err
	}
	if err := ctx.Validate(aksConfig); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	controller.AksConfigList = append(controller.AksConfigList, aksConfig)
	return ctx.NoContent(http.StatusCreated)
}

// AksValidator handles incoming validator requests
func (controller *Controller) AksSearch(ctx echo.Context) error {
	title := ctx.FormValue("title")
	if len(title) == 0 {
		return ctx.JSON(http.StatusBadRequest, "Title cannot be empty")
	}

	var aksConfigResult []AksConfig
	for _, val := range controller.AksConfigList {
		if strings.Contains(val.Title, title) {
			aksConfigResult = append(aksConfigResult, val)
		}
	}
	if len(aksConfigResult) == 0 {
		return ctx.JSON(http.StatusOK, "No AKS config with given title")
	}

	return ctx.JSON(http.StatusOK, aksConfigResult)
}

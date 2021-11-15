package api

import (
	"fmt"
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

type Config struct {
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
	ConfigList []Config
}

// AddAksConfig handles incoming validator requests
func (controller *Controller) AddAksConfig(ctx echo.Context) error {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return echo.ErrBadRequest
	}
	var Config Config
	if err := yaml.Unmarshal(body, &Config); err != nil {
		return err
	}
	if err := ctx.Validate(Config); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	controller.ConfigList = append(controller.ConfigList, Config)
	return ctx.NoContent(http.StatusCreated)
}

// GetAllAksConfig handles incoming validator requests
func (controller *Controller) GetAllAksConfig(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, controller.ConfigList)
}

// GetAksConfigByTitle handles incoming validator requests
func (controller *Controller) GetAksConfigByTitle(ctx echo.Context) error {
	title := ctx.Param("title")
	if len(title) == 0 {
		return ctx.JSON(http.StatusBadRequest, "Title cannot be empty")
	}

	var ConfigResult []Config
	for _, val := range controller.ConfigList {
		if strings.Contains(val.Title, title) {
			ConfigResult = append(ConfigResult, val)
		}
	}
	if len(ConfigResult) == 0 {
		return ctx.JSON(http.StatusOK, "No AKS config with given title")
	}

	return ctx.JSON(http.StatusOK, ConfigResult)
}

// GetAksConfigByTitle handles incoming validator requests
func (controller *Controller) DeleteAllAksConfig(ctx echo.Context) error {
	controller.ConfigList = []Config{}
	fmt.Println(("testing"))
	return ctx.NoContent(http.StatusOK)
}

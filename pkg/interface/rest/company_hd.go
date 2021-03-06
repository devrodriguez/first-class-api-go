package rest

import (
	"net/http"

	"github.com/devrodriguez/first-class-api-go/pkg/application"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	srv *application.CompanyService
}

func NewCompanyHandler(srv *application.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		srv,
	}
}

// Implementation
func (ch *CompanyHandler) GetAll(c *gin.Context) {

	companies, err := ch.srv.GetAll()

	if err != nil {
		c.JSON(http.StatusNoContent, APIResponse{
			Message: "not data found",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusNoContent),
					Status: http.StatusNoContent,
				},
			},
		})
	}

	c.JSON(http.StatusOK, APIResponse{
		Data: companies,
	})
}

// Implementation
func (ch *CompanyHandler) Create(c *gin.Context) {
	var company entity.Company

	// Get data from request
	if err := c.BindJSON(&company); err != nil {

		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "error binding data",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusBadRequest),
					Status: http.StatusBadRequest,
				},
			},
		})
		return
	}

	if err := ch.srv.Create(c, company); err != nil {
		c.JSON(http.StatusNotModified, APIResponse{
			Message: "error saving data",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusNotModified),
					Status: http.StatusNotModified,
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: `success`,
	})
}

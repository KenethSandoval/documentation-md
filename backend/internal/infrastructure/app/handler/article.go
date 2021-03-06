package handler

import (
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) CreateArticle(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.ArticleCreatePayload{}

	// biding payload to struct
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate payload struct
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// call usecase of business logic
	payload.ID = primitive.NewObjectID()
	// add timestamp
	if err := h.aru.CreateArticle(ctx, payload); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "CREATED")
}

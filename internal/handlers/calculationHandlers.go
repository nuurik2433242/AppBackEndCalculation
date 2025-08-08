package handlers

import (
	"go/adv-demo/internal/calcuationService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandlers struct {
	service calcuationService.CalculationService
}

func NewCalculationHandler (s calcuationService.CalculationService) *CalculationHandlers{
	return &CalculationHandlers{service: s}
}



// основные методы ORM - Create, Find, Update, Delete

func (h *CalculationHandlers) GetCalculation(c echo.Context) error {
	calculations, err := h.service.GetAllCalculations()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cound not get calculation"})
	}
	
	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandlers) PostCalculation(c echo.Context) error {
	var req calcuationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(req.Expression) 

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create calculation"})
	}

	return c.JSON(http.StatusCreated, calc)
}

func (h *CalculationHandlers) PatchCalculation(c echo.Context) error {
	id := c.Param("id")

	var req calcuationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updateCalc,err := h.service.UpdateCalculation(id,req.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
	}

	return c.JSON(http.StatusOK, updateCalc)
}

func (h *CalculationHandlers) DeleteCalculation(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent)
}
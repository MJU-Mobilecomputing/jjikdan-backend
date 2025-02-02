package interfaces

import (
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
)

type IGPTService interface {
	GetMenuNutrient(string) (*utils.MenuNeutrient, error)
	GetWeeklySolution(repository.FindWeeklySummaryRow) (*string, error)
}

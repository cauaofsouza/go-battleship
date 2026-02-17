package medal

import (
	"github.com/allanjose001/go-battleship/internal/entity"
)

type Medal struct {
	Name         string                              `json:"name"`
	Description  string                              `json:"description"`
	IconPath     string                              `json:"icon"`
	Verification func(stats entity.PlayerStats) bool `json:"-"` //cada medalha tem seus criterios de verificação
}

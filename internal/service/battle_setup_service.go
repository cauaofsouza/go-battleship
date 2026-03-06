package service

import (
	"fmt"
	"reflect"

	"github.com/allanjose001/go-battleship/game/shared/board"
	"github.com/allanjose001/go-battleship/game/shared/placement"
	"github.com/allanjose001/go-battleship/internal/ai"
	"github.com/allanjose001/go-battleship/internal/entity"
)

type BattleSetupService struct{}

func NewBattleSetupService() *BattleSetupService {
	return &BattleSetupService{}
}

// InitBattleAI inicializa a inteligência artificial com base na dificuldade selecionada.
// Parâmetros:
// - difficulty: string que define o nível ("easy", "medium", "hard")
// - playerShips: os navios que o jogador posicionou (usados para espelhar a frota da IA)
func (s *BattleSetupService) InitBattleAI(difficulty string, playerShips []*placement.ShipPlacement) (*ai.AIPlayer, *entity.Board, *entity.Fleet) {
	fleet := entity.NewFleet()
	entityBoard := &entity.Board{}

	usedShips := make(map[int]bool)

	// Mapeamento dos navios posicionados para a estrutura lógica da IA
	for _, ps := range playerShips {
		if !ps.Placed {
			continue
		}

		var entShip *entity.Ship
		for i, ship := range fleet.Ships {
			if !usedShips[i] && ship.Size == ps.Size {
				entShip = ship
				usedShips[i] = true
				break
			}
		}

		if entShip != nil {
			entShip.Horizontal = ps.Orientation == board.Horizontal
			entityBoard.PlaceShip(entShip, ps.Y, ps.X)
		}
	}

	var aiPlayer *ai.AIPlayer

	fmt.Printf("Iniciando batalha com dificuldade: %s\n", difficulty)
	
	switch difficulty {
	case "easy":
		aiPlayer = ai.NewEasyAIPlayer()
	case "medium":
		aiPlayer = ai.NewMediumAIPlayer(fleet)
	case "hard":
		aiPlayer = ai.NewHardAIPlayer(fleet)
	default:
		aiPlayer = ai.NewEasyAIPlayer()
	}
	fmt.Printf("AI Player Instanciado: %v\n", reflect.TypeOf(aiPlayer))

	return aiPlayer, entityBoard, fleet
}
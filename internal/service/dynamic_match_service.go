package service

import (
	"fmt"
	"time"

	"github.com/allanjose001/go-battleship/internal/entity"
)

type DynamicMatchService struct {
	*MatchService
}

// Corrigido: recebe attack e aiDelay (mesma semântica de NewMatchService)
func NewDynamicMatchService(attack *AttackService, aiDelay time.Duration) *DynamicMatchService {
	return &DynamicMatchService{
		MatchService: NewMatchService(attack, aiDelay),
	}
}

// MovePlayerShip tenta mover `ship` do jogador para (newRow,newCol).
// now: tempo atual usado para agendamento do próximo passo da IA.
// Observação: MoveShip foi implementado em internal/entity/Board (PlayerEntityBoard),
// portanto chamamos ali. Se quiser refletir no PlayerBoard (visual), sincronize depois.
func (s *DynamicMatchService) MovePlayerShip(m *entity.Match, ship *entity.Ship, newRow int, newCol int, now time.Time) error {
	if m == nil {
		return ErrMatchNotFound
	}
	if m.IsFinished() {
		return ErrMatchFinished
	}
	if m.Status != entity.MatchStatusInProgress {
		return ErrMatchNotInProgress
	}
	if m.Turn != entity.TurnPlayer {
		return ErrNotPlayersTurn
	}
	// precisa das referências runtime presentes
	if m.PlayerEntityBoard == nil || m.PlayerBoard == nil {
		return ErrMatchNotReady
	}
	if ship == nil {
		return fmt.Errorf("ship nil")
	}

	// checa se ship pertence à frota do jogador
	found := false
	if m.PlayerFleet != nil {
		for _, sh := range m.PlayerFleet.GetFleetShips() {
			if sh == ship {
				found = true
				break
			}
		}
	}
	if !found {
		return fmt.Errorf("ship not in player fleet")
	}

	// delega a movimentação para PlayerEntityBoard (onde MoveShip existe)
	if err := m.PlayerEntityBoard.MoveShip(ship, newRow, newCol); err != nil {
		return err
	}

	// consumir o turno do jogador: passa para IA e agenda próximo ataque
	m.Turn = entity.TurnEnemy
	m.NextAction = entity.NextActionEnemyAttack
	m.NextActionAt = now.Add(s.aiDelay)

	return nil
}

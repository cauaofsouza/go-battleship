package medal

import (
	"time"

	"github.com/allanjose001/go-battleship/internal/entity"
)

// medal_repository funciona como "repository" de medalhas hardcoded
//TODO: Adicionar caminhos dos icones das medalhas

//para evitar dor de cabeça foi preferido deixá-las hardcoded assim

// MedalsList lista de todas as medalhas do jogo
var MedalsList = []*Medal{
	{
		Name:        "Almirante",
		Description: "Venceu sem perder navios",
		IconPath:    "",
		Verification: func(stats entity.PlayerStats) bool {
			return stats.WinWithoutLosses
		},
	},
	{
		Name:        "Capitão",
		Description: "Acertou 7 tiros seguidos",
		IconPath:    "",
		Verification: func(stats entity.PlayerStats) bool {
			return stats.HigherHitSequence >= 7
		},
	},
	{
		Name:        "Capitão de Mar e Guerra",
		Description: "Acertou 8 tiros seguidos",
		IconPath:    "",
		Verification: func(stats entity.PlayerStats) bool {
			return stats.HigherHitSequence >= 8
		},
	},
	{
		Name:        "Marinheiro",
		Description: "Venceu em 1 minuto",
		IconPath:    "",
		Verification: func(stats entity.PlayerStats) bool {
			return stats.FasterTime <= time.Minute.Milliseconds()
		},
	},
}

// MedalsMap Map para acesso rápido pelo nome
var MedalsMap = make(map[string]*Medal)

// init inicializa map para facilitar load  profile do json com medalhas
func init() {
	for _, m := range MedalsList {
		MedalsMap[m.Name] = m
	}
}

// GetMedals serve para pegar os objetos medal pelo nome
// (para usar com lista de strings medalha do player)
func GetMedals(names []string) []*Medal {
	var result []*Medal
	for _, n := range names {
		if m, ok := MedalsMap[n]; ok {
			result = append(result, m)
		}
	}
	return result
}

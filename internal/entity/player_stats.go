package entity

import "fmt"

// PlayerStats struct que encapsula stats acumulados do player
type PlayerStats struct {
	Matches           int   `json:"matches"`
	Wins              int   `json:"wins"`
	TotalShots        int   `json:"total_shots"`
	TotalHits         int   `json:"total_hits"`
	HighScore         int   `json:"high_score"`
	TotalScore        int   `json:"total_score"`
	HigherHitSequence int   `json:"higher_hit_sequence"`
	FasterTime        int64 `json:"faster_time"` //tempo em milissegundos
	WinWithoutLosses  bool  `json:"win_without_losses"`
}

// WinRate retorna winrate do player
func (s *PlayerStats) WinRate() float32 {
	if s.Matches == 0 {
		return 0
	}
	return float32(s.Wins) / float32(s.Matches) * 100
}

// Accuracy retorna hitrate (relação tiros / acertos)
func (s *PlayerStats) Accuracy() float32 {
	if s.TotalShots == 0 {
		return 0
	}
	return float32(s.TotalHits) / float32(s.TotalShots) * 100
}

func (s *PlayerStats) ApplyMatch(r MatchResult) {
	s.Matches++

	if r.Win {
		s.Wins++
	}

	s.TotalShots += r.ShotsFired

	s.TotalHits += r.Hits

	s.TotalScore += r.Score

	if r.Score > s.HighScore {
		s.HighScore = r.Score
	}

	if r.HigherHitSequence > s.HigherHitSequence {
		s.HigherHitSequence = r.HigherHitSequence
	}

	if r.LostShips == 0 && r.Win {
		s.WinWithoutLosses = true
	}
	if r.Duration < s.FasterTime {
		s.FasterTime = r.Duration
	}

}

// FormattedFasterTime retorna string para ser usada no front
func (s *PlayerStats) FormattedFasterTime() string {
	totalSec := s.FasterTime / 1000
	minutes := totalSec / 60
	seconds := totalSec % 60

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

package service

import (
	"sync"

	"github.com/kyeett/highscore-server/model"
)

type ScoreDB struct {
	*sync.RWMutex
	Score []*model.Score
}

func NewScoreDB() ScoreDB {
	return ScoreDB{
		RWMutex: &sync.RWMutex{},
		Score:   []*model.Score{},
	}
}

func (db *ScoreDB) Add(m model.Score) {
	db.Lock()
	defer db.Unlock()

	db.Score = append(db.Score, &m)
}

func (db *ScoreDB) Count(g model.Game) int {
	db.RLock()
	defer db.RUnlock()

	var count int
	for _, s := range db.Score {
		if s.Game.Name != "" && (s.Game.Name != g.Name) {
			continue
		}
		count++
	}

	return count
}

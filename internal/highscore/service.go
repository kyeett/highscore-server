package highscore

import (
	"github.com/jmoiron/sqlx"
	"github.com/kyeett/highscore-server/model"
)

var _ Service = &BasicService{}

type Service interface {
	Add(*model.Score) error
	ListByGame(g model.Game) ([]*model.Score, error)
}

type BasicService struct {
	db *sqlx.DB
}

func NewBasic(db *sqlx.DB) *BasicService {
	return &BasicService{db: db}
}

func (h *BasicService) Add(m *model.Score) error {
	query := `
	INSERT INTO score (
		id,
		created_at,
		score,
		game_name,
		user_id
	) VALUES (
		:id,
		:created_at,
		:score,
		:game_name,
		:user_id
	)	
	`

	if _, err := h.db.NamedExec(query, m); err != nil {
		return err
	}
	return nil
}

func (h *BasicService) ListByGame(g model.Game) ([]*model.Score, error) {
	query := `SELECT * FROM score WHERE game_name = $1 ORDER BY score DESC`
	var list []*model.Score
	if err := h.db.Select(&list, query, g.Name); err != nil {
		return nil, err
	}
	return list, nil
}

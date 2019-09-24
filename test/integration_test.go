package test

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kyeett/highscore-server/client"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	"github.com/kyeett/highscore-server/internal/service"
	"github.com/kyeett/highscore-server/model"

	_ "github.com/lib/pq"
)

var svc *service.Service

const (
	databaseUrl = "postgres://testuser:testpass@localhost:5432/testuser?sslmode=disable"
)

func init() {
	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil {
		panic(fmt.Sprintf("could not connect to database: %v", err))
	}
	svc = service.New(db)
}

var (
	testGame1 = model.Game{Name: "My Awesome Game"}
	testGame2 = model.Game{Name: "My Not So Awesome Game"}

	scoreGame1 = &model.Score{
		ID:   uuid.New(),
		Game: testGame1,
		// User:      testUser1,
	}

	scoreGame2 = &model.Score{
		ID:   uuid.New(),
		Game: testGame2,
		// User:      testUser1,
	}
)

func scoreFromGame(score float64, game model.Game) *model.Score {
	u := model.User{ID: uuid.New().String()}

	return &model.Score{
		ID:    uuid.New(),
		Score: score,
		Game:  game,
		User:  u,
	}

}

func TestService(t *testing.T) {
	game1 := model.Game{Name: "My Awesome Game"}
	game2 := model.Game{Name: "My Awesomer Game"}
	score1 := scoreFromGame(10, game1)
	score2 := scoreFromGame(5, game2)

	assert.NoError(t, svc.Highscore.Add(score1))
	assert.NoError(t, svc.Highscore.Add(score2))

	game1Scores, err := svc.Highscore.ListByGame(game1)
	assert.NoError(t, err)

	game2Scores, err := svc.Highscore.ListByGame(game2)
	assert.NoError(t, err)

	assert.Len(t, game1Scores, 1)
	assert.EqualValues(t, score1, unsetCreatedAt(game1Scores[0]))

	assert.Len(t, game2Scores, 1)
	assert.EqualValues(t, score2, unsetCreatedAt(game2Scores[0]))
}

func TestResource(t *testing.T) {
	server := httptest.NewServer(svc.Router)
	fmt.Println(server.URL)
	// server.URL
	defer server.Close()

	c, err := client.New(server.URL, "My HTTP game")
	assert.NoError(t, err, "failed to create client: %v", err)

	addedScore := []float64{20, 21, 22, 23, 24}
	for _, score := range addedScore {
		if err := c.AddSimple(score); err != nil {
			t.Fatalf("failed to add score: %v", err)
		}
	}

	retrievedScore, err := c.ListSimple()
	if err != nil {
		t.Fatalf("failed to retrieve highscore: %v", err)
	}
	assert.Len(t, retrievedScore, len(addedScore))
}

func unsetCreatedAt(m *model.Score) *model.Score {
	m.CreatedAt = time.Time{}
	return m
}

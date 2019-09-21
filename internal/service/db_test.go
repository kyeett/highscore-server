package service_test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/kyeett/highscore-server/internal/service"
	"github.com/kyeett/highscore-server/model"
)

var (
	testGame1 = model.Game{Name: "My Awesome Game"}
	testGame2 = model.Game{Name: "My Not So Awesome Game"}

	testUser1 = model.User{Name: "Magnus", ID: uuid.New()}
	testUser2 = model.User{Name: "John", ID: uuid.New()}
)

func TestAdd(t *testing.T) {
	scoreGame1 := model.Score{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Game:      testGame1,
		User:      testUser1,
	}

	scoreGame2 := model.Score{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Game:      testGame2,
		User:      testUser1,
	}

	db := service.NewScoreDB()
	db.Add(scoreGame1)
	db.Add(scoreGame2)
	db.Add(scoreGame2)

	// Expected 1 score for game 1
	count := db.Count(testGame1)
	if count != 1 {
		t.Fatalf("expected 1 entries for game 1, got %d", count)
	}

	// Expected 1 score for game 1
	count = db.Count(testGame2)
	if count != 2 {
		t.Fatalf("expected 2 entries for game 2, got %d", count)
	}
}

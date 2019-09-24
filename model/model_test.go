package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {

	m1 := Score{}
	m1.Game.Name = "My HTTP game"
	m1.Score = 123.0

	JSONPayload := fmt.Sprintf(`{
		"id": "10b3c86b-c25e-4e9b-ad60-5cd62b29e0ab",
		"created_at": "0001-01-01T00:00:00Z",
		"score": %f,
		"user_id": "65aa943acce98f5a3b8798b216bb373dcdb39f1049a3e92d5746afe65236ca49",
		"user_name": "",
		"game_name": "%s"
		}`, m1.Score, m1.Game.Name)

	m2 := Score{}
	if err := json.Unmarshal([]byte(JSONPayload), &m2); err != nil {
		t.Fatalf("failed to unmarshal payload: %v", err)
	}

	if m2.Game.Name != m1.Game.Name {
		t.Fatalf("Unmarshalled game name not identical, Expected %s, got %s", m1.Game.Name, m2.Game.Name)
	}
}

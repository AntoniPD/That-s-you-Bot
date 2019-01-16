package models

// Game -
// ID - id of the game
// Players - players in current game
// Standings - user points standings
type Game struct {
	ID        string   `json:"id,omitempty"`
	Players   []Player `json:"players,omitempty"`
	Standings []int    `json:"standings,omitempty"`
}

// StartNewGame - {WIP}
// When all players are agreed to start game,
// saves the game in the database.
func StartNewGame(game Game) string {
	return "gameId"
}

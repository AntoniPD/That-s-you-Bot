package models

import models "msgbot/models/messenger"

// Player -
// UserData - info from facebook
// Points - points for current game
// Jokers - jokers for current game
// Images - images send in current game
type Player struct {
	UserData models.UserInfo `json:"user_info,omitempty"`
	Points   int             `json:"points,omitempty"`
	Jokers   int             `json:"jokers,omitempty"`
	Images   []string        `json:"images,omitempty"`
	InGame   bool            `json:"in_game,omitempty"`
}

// AddPlayer - {WIP}
// Add player if not exist in the database.
// returns the saved player id.
func AddPlayer(player Player) string {
	return "playerId"
}

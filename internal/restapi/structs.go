package faceit

type (
	// User is the response faceituser payload
	User struct {
		ID         string `json:"id"`
		Username   string `json:"username"`
		SkillLevel int    `json:"skilllevel"`
		Steam      string `json:"steam"`
	}

	// Game is the game struct
	Game struct {
		SkillLevel int `json:"skill_level" validate:"required"`
	}

	// RawUser is the raw response from the faceit api
	RawUser struct {
		ID       string          `json:"player_id" validate:"required"`
		Username string          `json:"nickname" validate:"required"`
		Games    map[string]Game `json:"games" validate:"required"`
	}

	// ReqInvitePayload is the raw payload used to get invites
	ReqInvitePayload struct {
		EntityID   string `json:"entity_id"`
		EntityType string `json:"entity_type"`
		Type       string `json:"type"`
		MaxAge     int    `json:"max_age"`
		MaxUses    int    `json:"max_uses"`
	}

	// Invite is the wrapper payload returned by rest API
	Invite struct {
		Payload InvitePayload `json:"payload"`
	}

	// InvitePayload is the payload
	InvitePayload struct {
		Code string `json:"code"`
	}
)

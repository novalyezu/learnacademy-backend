package community

import (
	"time"

	"github.com/novalyezu/learnacademy-backend/user"
)

type Community struct {
	ID               string
	UserID           string
	Name             string
	Slug             string
	Thumbnail        string
	ShortDescription string
	Description      string
	IsPublic         bool
	TotalMember      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	User             user.User
}

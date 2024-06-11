package community

import (
	"time"

	"github.com/novalyezu/learnacademy-backend/user"
)

type CreateCommunityInput struct {
	Name             string `form:"name" binding:"required"`
	ShortDescription string `form:"short_description" binding:"required"`
	Description      string `form:"description" binding:"required"`
	IsPublic         bool   `form:"is_public" binding:"required"`
	Thumbnail        string
	UserID           string
}

type CommunityOutput struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	Slug             string          `json:"slug"`
	Thumbnail        string          `json:"thumbnail"`
	ShortDescription string          `json:"short_description"`
	Description      string          `json:"description"`
	IsPublic         bool            `json:"is_public"`
	TotalMember      int             `json:"total_member"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	User             user.UserOutput `json:"user"`
}

func FormatToCommunityOutput(community Community) CommunityOutput {
	user := user.UserOutput{
		ID:        community.User.ID,
		Firstname: community.User.Firstname,
		Lastname:  community.User.Lastname,
		Email:     community.User.Email,
	}
	return CommunityOutput{
		ID:               community.ID,
		Name:             community.Name,
		Slug:             community.Slug,
		Thumbnail:        community.Thumbnail,
		ShortDescription: community.ShortDescription,
		Description:      community.Description,
		IsPublic:         community.IsPublic,
		TotalMember:      community.TotalMember,
		CreatedAt:        community.CreatedAt,
		UpdatedAt:        community.UpdatedAt,
		User:             user,
	}
}

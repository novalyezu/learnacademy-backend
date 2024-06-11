package community

import (
	"strconv"
	"time"

	"github.com/gosimple/slug"
	"github.com/novalyezu/learnacademy-backend/helper"
)

type CommunityService interface {
	Create(input CreateCommunityInput) (Community, error)
}

type communityServiceImpl struct {
	communityRepository CommunityRepository
}

func NewCommunityService(communityRepository CommunityRepository) CommunityService {
	return &communityServiceImpl{
		communityRepository: communityRepository,
	}
}

func (s *communityServiceImpl) Create(input CreateCommunityInput) (Community, error) {
	ID, errUlid := helper.NewID()
	if errUlid != nil {
		return Community{}, errUlid
	}

	community := Community{
		ID:               ID,
		UserID:           input.UserID,
		Name:             input.Name,
		Slug:             slug.Make(input.Name + strconv.Itoa(int(time.Now().Unix()))),
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		IsPublic:         input.IsPublic,
		Thumbnail:        input.Thumbnail,
		TotalMember:      0,
	}

	newCommunity, errSave := s.communityRepository.Save(community)
	if errSave != nil {
		return Community{}, errSave
	}

	return newCommunity, nil
}

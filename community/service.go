package community

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/novalyezu/learnacademy-backend/helper"
)

type CommunityService interface {
	Create(input CreateCommunityInput) (CommunityDetailOutput, error)
	GetAll(input GetCommunityInput) (CommunitiesOutput, error)
}

type communityServiceImpl struct {
	communityRepository CommunityRepository
}

func NewCommunityService(communityRepository CommunityRepository) CommunityService {
	return &communityServiceImpl{
		communityRepository: communityRepository,
	}
}

func (s *communityServiceImpl) Create(input CreateCommunityInput) (CommunityDetailOutput, error) {
	ID, errUlid := helper.NewID()
	if errUlid != nil {
		return CommunityDetailOutput{}, errUlid
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
		return CommunityDetailOutput{}, errSave
	}

	communityOutput := FormatToCommunityDetailOutput(newCommunity)

	return communityOutput, nil
}

func (s *communityServiceImpl) GetAll(input GetCommunityInput) (CommunitiesOutput, error) {
	page := helper.Ternary(input.Page < 1, 1, input.Page).(int)
	limit := helper.Ternary(input.Limit < 1, 10, input.Limit).(int)
	orderBy := helper.Ternary(input.OrderBy == "", "created_at__desc", input.OrderBy).(string)
	orderBy = strings.ReplaceAll(orderBy, "__", " ")
	offset := (page - 1) * limit

	var (
		condition         string
		conditionArgs     []any
		communitiesOutput []CommunityOutput
	)

	if input.Search != "" {
		condition += "name LIKE ? OR description LIKE ?"
		conditionArgs = append(conditionArgs, "%"+input.Search+"%")
		conditionArgs = append(conditionArgs, "%"+input.Search+"%")
	}

	communities, errFind := s.communityRepository.FindAll(FindAllParams{
		Limit:         limit,
		Offset:        offset,
		OrderBy:       orderBy,
		Condition:     condition,
		ConditionArgs: conditionArgs,
	})
	if errFind != nil {
		return CommunitiesOutput{}, errFind
	}

	count, errCount := s.communityRepository.Count(condition, conditionArgs)
	if errCount != nil {
		return CommunitiesOutput{}, errCount
	}
	totalPage := math.Ceil(float64(count) / float64(limit))

	for _, community := range communities {
		communitiesOutput = append(communitiesOutput, FormatToCommunityOutput(community))
	}

	return CommunitiesOutput{
		Communities: communitiesOutput,
		Meta: helper.PaginationMeta{
			CurrentPage: page,
			NextPage:    helper.Ternary(int(totalPage) > page, page+1, -1).(int),
			PrevPage:    helper.Ternary(page > 1, page-1, -1).(int),
			TotalData:   int(count),
			TotalPage:   int(totalPage),
		},
	}, nil
}

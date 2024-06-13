package community

import "gorm.io/gorm"

type FindAllParams struct {
	Limit         int
	Offset        int
	OrderBy       string
	Condition     string
	ConditionArgs []any
}

type CommunityRepository interface {
	Save(community Community) (Community, error)
	FindBySlug(slug string) (Community, error)
	FindAll(params FindAllParams) ([]Community, error)
	Count(condition string, conditionArgs []any) (int64, error)
}

type communityRepositoryImpl struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) CommunityRepository {
	return &communityRepositoryImpl{
		db: db,
	}
}

func (r *communityRepositoryImpl) Save(community Community) (Community, error) {
	err := r.db.Create(&community).Error
	if err != nil {
		return Community{}, err
	}
	return community, nil
}

func (r *communityRepositoryImpl) FindBySlug(slug string) (Community, error) {
	var community Community
	err := r.db.Where("slug = ?", slug).First(&community).Error
	if err != nil {
		return community, err
	}
	return community, nil
}

func (r *communityRepositoryImpl) FindAll(params FindAllParams) ([]Community, error) {
	var communities []Community

	err := r.db.Where(params.Condition, params.ConditionArgs...).
		Order(params.OrderBy).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&communities).
		Error
	if err != nil {
		return []Community{}, err
	}

	return communities, nil
}

func (r *communityRepositoryImpl) Count(condition string, conditionArgs []any) (int64, error) {
	var count int64

	err := r.db.Model(&Community{}).Where(condition, conditionArgs...).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

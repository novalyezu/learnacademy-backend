package community

import "gorm.io/gorm"

type CommunityRepository interface {
	Save(community Community) (Community, error)
	FindBySlug(slug string) (Community, error)
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

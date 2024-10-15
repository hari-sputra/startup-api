package campaign

import "gorm.io/gorm"

type CampaignRepository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type campaignRepo struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *campaignRepo {
	return &campaignRepo{db}
}

func (r *campaignRepo) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepo) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

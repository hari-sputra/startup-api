package campaign

type CampaignService interface {
	FindCampaign(userID int) ([]Campaign, error)
}

type campaignService struct {
	campaignRepo CampaignRepository
}

func NewCampaignService(campaignRepo CampaignRepository) *campaignService {
	return &campaignService{campaignRepo}
}

func (s *campaignService) FindCampaign(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaign, err := s.campaignRepo.FindByUserID(userID)
		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaign, err := s.campaignRepo.FindAll()
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

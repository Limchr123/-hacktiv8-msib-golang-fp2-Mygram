package campaign

type CampaignFormatter struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

func FormatterUser(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:       campaign.ID,
		Title:    campaign.Title,
		Caption:  campaign.Title,
		PhotoUrl: campaign.PhotoUrl,
		UserId:   campaign.UserId,
	}
	return formatter
}

package campaign

type CampaignFormatter struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

func FormatterCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:       campaign.ID,
		Title:    campaign.Title,
		Caption:  campaign.Title,
		PhotoUrl: campaign.PhotoUrl,
		UserId:   campaign.UserId,
	}
	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		//inisiasi campaignFormatter untuk dimasukkan ke dalam fungsi FormatCampaign diatas agar dibaca
		//masih kurang lengkap penjelasannya
		campaignFormatter := FormatterCampaign(campaign)
		//menggunakan append kalau ada lagi maka masukkan campaignFormatter
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

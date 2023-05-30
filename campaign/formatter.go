package campaign

import (
	"time"
)

type CampaignCreateFormatter struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatterCreateCampaign(campaign Campaign) CampaignCreateFormatter {
	formatterCreate := CampaignCreateFormatter{
		ID:        campaign.ID,
		Title:     campaign.Title,
		Caption:   campaign.Title,
		PhotoUrl:  campaign.PhotoUrl,
		UserId:    campaign.UserId,
		CreatedAt: campaign.CreatedAt,
	}
	return formatterCreate
}

type CampaignUserFormatter struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
}

type CampaignGetFormatter struct {
	ID        int                   `json:"id"`
	Title     string                `json:"title"`
	Caption   string                `json:"caption"`
	PhotoUrl  string                `json:"photo_url"`
	UserId    int                   `json:"user_id"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	User      CampaignUserFormatter `json:"user"`
}

func FormatterGet(campaign Campaign) CampaignGetFormatter {
	formatterGet := CampaignGetFormatter{}
	formatterGet.ID = campaign.ID
	formatterGet.Title = campaign.Title
	formatterGet.Caption = campaign.Caption
	formatterGet.PhotoUrl = campaign.PhotoUrl
	formatterGet.UserId = campaign.UserId
	formatterGet.CreatedAt = campaign.CreatedAt
	formatterGet.UpdatedAt = campaign.UpdatedAt

	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Email = user.Email
	campaignUserFormatter.UserName = user.Username

	formatterGet.User = campaignUserFormatter

	return formatterGet
}

func FormatterGetCampaign(campaigns []Campaign) []CampaignGetFormatter {
	campaignGetFormatter := []CampaignGetFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatterGet(campaign)
		campaignGetFormatter = append(campaignGetFormatter, campaignFormatter)
	}

	return campaignGetFormatter
}

type CampaignUpdatedFormatter struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatterUpdatedCampaign(campaign Campaign) CampaignUpdatedFormatter {
	formatterUpdated := CampaignUpdatedFormatter{
		ID:        campaign.ID,
		Title:     campaign.Title,
		Caption:   campaign.Title,
		PhotoUrl:  campaign.PhotoUrl,
		UserId:    campaign.UserId,
		UpdatedAt: campaign.UpdatedAt,
	}
	return formatterUpdated
}

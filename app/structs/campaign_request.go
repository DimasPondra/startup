package structs

type CampaignStoreRequest struct {
	Name             string `json:"name" validate:"required,campaign_name_available"`
	ShortDescription string `json:"short_description" validate:"required"`
	Description      string `json:"description" validate:"required"`
	GoalAmount       int    `json:"goal_amount" validate:"required,gt=1000000"`
	Perks            string `json:"perks" validate:"required"`
	User             User
}

type CampaignUpdateRequest struct {
	Name             string `json:"name" validate:"required,campaign_name_available"`
	ShortDescription string `json:"short_description" validate:"required"`
	Description      string `json:"description" validate:"required"`
	GoalAmount       int    `json:"goal_amount" validate:"required,gt=1000000"`
	Perks            string `json:"perks" validate:"required"`
}
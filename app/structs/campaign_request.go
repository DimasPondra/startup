package structs

type CampaignStoreRequest struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required,number,gt=0"`
	Perks            string `json:"perks" binding:"required"`
	User             User
}

type CampaignUpdateRequest struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required,number,gt=0"`
	Perks            string `json:"perks" binding:"required"`
}
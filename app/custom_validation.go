package app

import (
	"startup/app/services"
	"startup/app/structs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegisterEmailAvailableValidation(validate *validator.Validate, userService services.UserService) error {
	err := validate.RegisterValidation("email_available", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())

		checkEmail := structs.CheckEmailRequest{
			Email: value.String(),
		}

		emailAvailable, _ := userService.IsEmailAvailable(checkEmail)
		return emailAvailable
	})

	return err
}

func RegisterCampaignNameAvailableValidation(validate *validator.Validate, campaignService services.CampaignService, c *gin.Context) error {
	slug := c.Param("slug")

	err := validate.RegisterValidation("campaign_name_available", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())

		checkCampaign, _ := campaignService.GetCampaignByName(value.String())

		// for update
		if slug != "" {
			if checkCampaign.ID == 0 {
				return true
			} else {
				campaignBySlug, _ := campaignService.GetCampaignBySlug(slug)

				return checkCampaign.Name == value.String() && checkCampaign.ID == campaignBySlug.ID
			}
		}

		// for create
		return checkCampaign.ID == 0
	})

	return err
}

func RegisterExistsInCampaignsValidation(validate *validator.Validate, campaignService services.CampaignService) error {
	err := validate.RegisterValidation("exists_in_campaigns", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())

		campaign, _ := campaignService.GetCampaignByID(int(value.Int()))

		return campaign.ID != 0
	})

	return err
}
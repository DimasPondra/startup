package app

import (
	"mime/multipart"
	"startup/app/services"
	"startup/app/structs"
	"strings"

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

func RegisterRoleNameAvailableValidation(validate *validator.Validate, roleService services.RoleService) error {
	err := validate.RegisterValidation("role_name_available", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())

		role, _ := roleService.GetRoleByName(value.String())

		return role.ID == 0
	})

	return err
}

func RegisterImageTypeValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("image_type", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())

		files := value.Interface().([]*multipart.FileHeader)

		for _, file := range files {
			resultOfSplit := strings.Split(file.Filename, ".")
			lengthOfSplit := len(resultOfSplit)

			imageType := resultOfSplit[lengthOfSplit-1]

			if imageType != "jpg" && imageType != "png" {
				return false
			}
		}

		return true
	})

	return err
}

func RegisterExistsInFilesValidation(validate *validator.Validate, fileService services.FileService) error {
	err := validate.RegisterValidation("exists_in_files", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())

		file, _ := fileService.GetFileByID(int(value.Int()))

		return file.ID != 0
	})

	return err
}

func RegisterIDsExistsInFilesValidation(validate *validator.Validate, fileService services.FileService) error {
	err := validate.RegisterValidation("ids_exists_in_files", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())
		IDs := value.Interface().([]int)

		for _, ID := range IDs {
			file, _ := fileService.GetFileByID(ID)

			if file.ID == 0 {
				return false
			}
		}

		return true
	})

	return err
}

package validation

import (
	"golang-web-service/model/modelsocialmedia"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateSocialMediaCreate(data modelsocialmedia.Request) error {
	return validation.Errors{
		"name":             validation.Validate(data.Name, validation.Required),
		"social_media_url": validation.Validate(data.SocialMediaUrl, validation.Required),
	}.Filter()
}

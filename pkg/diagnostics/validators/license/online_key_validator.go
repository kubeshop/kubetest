package license

import (
	"regexp"
	"strings"

	"github.com/kubeshop/testkube/pkg/diagnostics/validators"
)

const exampleLicense = "AB24F3-405E39-C3F657-94D113-F06C13-V3"

func NewOnlineLicenseKeyValidator() LocalLicenseKeyValidator {
	return LocalLicenseKeyValidator{}
}

type LocalLicenseKeyValidator struct {
}

// Validate validates a given license key for format / length correctness without calling external services
func (v LocalLicenseKeyValidator) Validate(subject any) validators.ValidationResult {
	// get key
	key, ok := subject.(string)
	if !ok {
		return ErrInvalidLicenseFormat
	}

	if key == "" {
		return validators.ValidationResult{
			Status: validators.StatusInvalid,
			Errors: []validators.ErrorWithSuggesstion{
				{
					Error:       ErrLicenseKeyNotFound,
					Suggestions: Suggestions[ErrLicenseKeyNotFound],
				},
			},
		}

	}

	// Check if the license key is the correct length and validate
	if len(key) == 29 {
		// Check if the license key matches the expected format
		match, _ := regexp.MatchString(`^[A-Z0-9]{6}-[A-Z0-9]{6}-[A-Z0-9]{6}-[A-Z0-9]{6}-[A-Z0-9]{6}-[A-Z0-9]{1-2}$`, key)
		if !match {
			println(match)
			return validators.ValidationResult{
				Status: validators.StatusInvalid,
				Errors: []validators.ErrorWithSuggesstion{
					{
						Error:       ErrLicenseKeyInvalidFormat,
						Suggestions: Suggestions[ErrLicenseKeyInvalidFormat],
					},
				},
			}
		}
	}

	// key can be in enrypted format
	if strings.HasPrefix(key, "key/") {
		// TODO validate air gapped key

	}

	return validators.NewValidResponse()
}

package validation

import (
	"regexp"
	"strings"

	"github.com/Oudwins/zog"
)

var RegisterUserRequestSchema = zog.Struct(zog.Shape{
	"email": zog.String().
		Trim().
		Required().
		Email().
		Min(5).
		Max(254).
		Transform(func(valPtr *string, ctx zog.Ctx) error {
			*valPtr = strings.ToLower(*valPtr)
			return nil
		}),
	"username": zog.String().
		Trim().
		Required().
		Min(3).
		Max(30).
		Match(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)),
	"password": zog.String().
		Trim().
		Required().
		Min(6).
		Max(128),
})

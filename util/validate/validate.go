package validate

import (
	"context"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/uniplaces/carbon"

	"github.com/kanok-p/go-clean-architecture/config"
	"github.com/kanok-p/go-clean-architecture/repository/users"
)

//go:generate mockery --name=Validator
type Validator interface {
	Validate(ctx context.Context, params interface{}) error
}

type VLDService struct {
	usrRepo users.Repository
	config  *config.Config
	formats map[string][]string
}

func New(repo users.Repository, config *config.Config, formats ...map[string][]string) *VLDService {
	return &VLDService{
		usrRepo: repo,
		config:  config,
		formats: formats[0],
	}
}

func (vld VLDService) Validate(ctx context.Context, params interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("unique", vld.unique(ctx))
	_ = validate.RegisterValidation("date-format", vld.isDateFormat())
	_ = validate.RegisterValidation("password-format", vld.format(vld.getFormat("password")))
	err := validate.Struct(params)

	return err
}

func (vld VLDService) getFormat(name string) []string {
	return vld.formats[name]
}

func (vld VLDService) unique(ctx context.Context) validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface()
		field := fl.StructFieldName()
		_, err := vld.usrRepo.Get(ctx, field, value)
		if err != nil {
			return true
		}

		return false
	}
}

func (vld VLDService) isDateFormat() validator.Func {
	return func(fl validator.FieldLevel) bool {
		format := fl.Param()
		dateString := fl.Field().String()
		_, err := carbon.CreateFromFormat(format, dateString, vld.config.TimeZone)

		return err == nil
	}
}

func (vld VLDService) format(formats []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		for _, format := range formats {
			re := regexp.MustCompile(format)
			if !re.MatchString(field) {
				return false
			}
		}

		return true
	}
}

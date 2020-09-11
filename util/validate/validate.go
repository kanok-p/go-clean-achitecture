package validate

import (
	"context"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/uniplaces/carbon"

	"github.com/kanok-p/go-clean-achitecture/config"
	"github.com/kanok-p/go-clean-achitecture/repository/users"
)

type Validator interface {
	Validate(ctx context.Context, params interface{}) error
	Unique(ctx context.Context) validator.Func
	IsDateFormat() validator.Func
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
	_ = validate.RegisterValidation("unique", vld.Unique(ctx))
	_ = validate.RegisterValidation("date-format", vld.IsDateFormat())
	_ = validate.RegisterValidation("password-format", vld.Format(vld.getFormat("password")))
	err := validate.Struct(params)

	return err
}

func (vld VLDService) getFormat(name string) []string {
	return vld.formats[name]
}

func (vld VLDService) Unique(ctx context.Context) validator.Func {
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

func (vld VLDService) IsDateFormat() validator.Func {
	return func(fl validator.FieldLevel) bool {
		format := fl.Param()
		dateString := fl.Field().String()
		_, err := carbon.CreateFromFormat(format, dateString, vld.config.TimeZone)

		return err == nil
	}
}

func (vld VLDService) Format(formats []string) validator.Func {
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

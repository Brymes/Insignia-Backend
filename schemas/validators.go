package schemas

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type BaseSchema interface {
	Validate(c *gin.Context, response *Handler)
}

func ValidateAndBind(validatorTarget, target interface{}, c *gin.Context, response *Handler) bool {
	var (
		ve     validator.ValidationErrors
		fields []string
	)

	if err := c.ShouldBindBodyWith(validatorTarget, binding.JSON); err != nil {
		response.Logger.Println(err)
		if errors.As(err, &ve) {
			for _, fieldError := range ve {
				fields = append(fields, fieldError.Field())
			}
		}
		response.Logger.Println(fields)
		failValidation(response, fields)
		return false
	}

	if target != nil {
		if err := c.ShouldBindBodyWith(target, binding.JSON); err != nil {
			if errors.As(err, &ve) {
				for _, fieldError := range ve {
					fields = append(fields, fieldError.Field())
				}
			}
			failValidation(response, fields)
		}
	}

	return true
}

func failValidation(response *Handler, fields []string) {
	response.Status, response.Message = 400, "Invalid Request Payload"
	response.ResponseData = gin.H{
		"fields": fields,
	}
	panic("Invalid Request Body")
}

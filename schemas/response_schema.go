package schemas

import (
	"Insignia-Backend/config"
	"Insignia-Backend/models"
	"Insignia-Backend/utils"
	"bytes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	models.SQLModel `json:"-"`
	buff            *bytes.Buffer `json:"-"`
	Logger          *log.Logger   `json:"-"`
	Status          int           `json:"-"`
	Success         bool          `json:"success"`
	Message         string        `json:"message"`
	ResponseData    interface{}   `json:"data"`
}

func NewHandler(logName, defaultMessage string) *Handler {
	buff, logger := config.InitRequestLogger(logName)
	if defaultMessage == "" {
		defaultMessage = "Unable to Process Request"
	}

	return &Handler{
		SQLModel: models.SQLModel{},
		buff:     buff,
		Logger:   logger,
		Success:  false,
		Message:  defaultMessage,
	}
}

func (handler *Handler) HandlePanic(err interface{}) {
	if utils.HandlePanicMacro(err, handler.Logger) {
		if handler.Status == 0 {
			handler.Status = 500
		}
		handler.Success = false
		if handler.Message == "" {
			handler.Message = "Error processing your request at the current time. Kindly contact support"
		}
	}
}
func (handler *Handler) LogToConsole() {
	log.Println(handler.buff)
}

func (handler *Handler) SendResponse(c *gin.Context) {
	defer log.Println(handler.buff)
	err := recover()
	handler.HandlePanic(err)

	if !handler.Success || handler.Status != 0 {
		c.IndentedJSON(handler.Status, handler)
	} else {
		handler.Success = true
		c.IndentedJSON(http.StatusOK, handler)
	}
}

func (handler *Handler) SendDBResponse(c *gin.Context) {
	handler.Status, handler.Message = handler.DBRStatus, handler.DBRMessage
	handler.SendResponse(c)
}

func (handler *Handler) SendResponseWithDefaults(c *gin.Context, status int, message string) {
	handler.Message, handler.Status = message, status
	handler.SendResponse(c)
}

func (handler *Handler) ValidatorMacro(c *gin.Context, validator BaseSchema) {
	validator.Validate(c, handler)
}

func (handler *Handler) HandleError(err error, code int) {
	if err != nil {
		handler.Status = code
	}
}

func (handler *Handler) HandleDbErrorMacro(result *gorm.DB, message string) {
	handler.Message, handler.Status = utils.HandleDbError(result, handler.Logger)

	if handler.Status != 0 {
		handler.Success, handler.ResponseData = false, nil
	} else {
		handler.Message, handler.Success = message, true
	}
}

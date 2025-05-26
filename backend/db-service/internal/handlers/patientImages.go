package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/jwt"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/logger"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/resources"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatientImagesHandlerInterface interface {
	Create(context *gin.Context)
	GetByPatientId(context *gin.Context)
	GetById(context *gin.Context)
	GetByUserId(context *gin.Context)
}

type PatientImagesHandler struct {
	PatientImagesService service.PatientImagesServiceInterface
	Logger               logger.LoggerInterface
	JWTManager           jwt.JWTManagerInterface
	ResponseManager      response.ResponseManagerInterface
	ResourceManager      resources.ResourceManagerInterface
}

func NewPatientImagesHandler(
	patientImagesService service.PatientImagesServiceInterface,
	logger logger.LoggerInterface,
	jwtManager jwt.JWTManagerInterface,
	responseManager response.ResponseManagerInterface,
	resourceManager resources.ResourceManagerInterface,
) PatientImagesHandlerInterface {
	return &PatientImagesHandler{
		PatientImagesService: patientImagesService.(*service.PatientImagesService),
		Logger:               logger,
		JWTManager:           jwtManager,
		ResponseManager:      responseManager,
		ResourceManager:      resourceManager,
	}
}

func (handler *PatientImagesHandler) Create(context *gin.Context) {
	action := "Create Patient Image"
	var patientImage models.PatientImages
	if err := context.ShouldBindJSON(&patientImage); err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Input", nil, err)
		return
	}

	createdPatientImage, err := handler.PatientImagesService.Create(&patientImage)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to create patient image", nil, err)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetPatientImagesResource(), map[string]interface{}{"patientImage": createdPatientImage})
	handler.ResponseManager.Send(context, http.StatusOK, "Patient Image Created", createdPatientImage, nil)

}

func (handler *PatientImagesHandler) GetByPatientId(context *gin.Context) {
	action := "Get Patient Images By Patient Id"
	patientIdString := context.Param("patientId")
	if patientIdString == "" {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Patient Id Parameter", nil, err)
		return
	}

	patientId, err := uuid.Parse(patientIdString)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Patient Id", nil, err)
		return
	}

	patientImages, err := handler.PatientImagesService.GetByPatientId(patientId)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get patient images", nil, err)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetPatientImagesResource(), map[string]interface{}{"patientImages": patientImages})
	handler.ResponseManager.Send(context, http.StatusOK, "Patient Images Retrieved", patientImages, nil)
}

func (handler *PatientImagesHandler) GetById(context *gin.Context) {
	action := "Get Patient Image By Id"
	idString := context.Param("id")
	if idString == "" {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Id Parameter", nil, err)
		return
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Id", nil, err)
		return
	}

	patientImage, err := handler.PatientImagesService.GetById(id)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get patient image", nil, err)
		return
	}

	if patientImage == nil {
		err := errors.New("patient image not found")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusNotFound, "Patient Image not found", nil, nil)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetPatientImagesResource(), map[string]interface{}{"patientImage": patientImage})
	handler.ResponseManager.Send(context, http.StatusOK, "Patient Image Retrieved", patientImage, nil)
}

func (handler *PatientImagesHandler) GetByUserId(context *gin.Context) {
	action := "Get Patient Images By User Id"
	userIdString := context.Param("userId")
	if userIdString == "" {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid User Id Parameter", nil, err)
		return
	}

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid User Id", nil, err)
		return
	}

	patientImages, err := handler.PatientImagesService.GetByUserId(userId)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientImagesResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get patient images", nil, err)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetPatientImagesResource(), map[string]interface{}{"patientImages": patientImages})
	handler.ResponseManager.Send(context, http.StatusOK, "Patient Images Retrieved", patientImages, nil)
}

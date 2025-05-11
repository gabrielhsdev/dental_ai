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

type PatientHandlerInterface interface {
	GetPatientById(context *gin.Context)
	GetPatientsByUserId(context *gin.Context)
	CreatePatient(context *gin.Context)
}

type PatientHandler struct {
	PatientService  service.PatientsServiceInterface
	Logger          logger.LoggerInterface
	JWTManager      jwt.JWTManagerInterface
	ResponseManager response.ResponseManagerInterface
	ResourceManager resources.ResourceManagerInterface
}

func NewPatientHandler(
	patientService service.PatientsServiceInterface,
	logger logger.LoggerInterface,
	jwtManager jwt.JWTManagerInterface,
	responseManager response.ResponseManagerInterface,
	resourceManager resources.ResourceManagerInterface,
) PatientHandlerInterface {
	return &PatientHandler{
		PatientService:  patientService.(*service.PatientService),
		Logger:          logger,
		JWTManager:      jwtManager,
		ResponseManager: responseManager,
		ResourceManager: resourceManager,
	}
}

func (handler *PatientHandler) GetPatientsByUserId(context *gin.Context) {
	userIdString := context.Param("userId")
	action := "Get Patients By User Id"

	if userIdString == "" {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid User Id Parameter", nil, err)
	}

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid User Id", nil, err)
		return
	}

	patients, err := handler.PatientService.GetPatientsByUserId(userId)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get patients", nil, err)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetPatientResource(), map[string]interface{}{"patients": patients})
	handler.ResponseManager.Send(context, http.StatusOK, "Patients found", patients, nil)
}

func (handler *PatientHandler) GetPatientById(context *gin.Context) {
	idString := context.Param("id")
	action := "Get Patient By Id"

	if idString == "" {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Id Parameter", nil, err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Id", nil, err)
		return
	}

	patient, err := handler.PatientService.GetPatientById(id)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get patient", nil, err)
		return
	}
	if patient == nil {
		err := errors.New("patient not found")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusNotFound, "Patient not found", nil, nil)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetPatientResource(), map[string]interface{}{"patient": patient})
	handler.ResponseManager.Send(context, http.StatusOK, "Patient found", patient, nil)
}

func (handler *PatientHandler) CreatePatient(context *gin.Context) {
	action := "Create Patient"
	token := context.Request.Header.Get("Authorization")
	var patient models.Patient

	if token == "" {
		err := errors.New("unauthorized, token not found")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Unauthorized, Token Not Found", nil, nil)
		return
	}

	claims, err := handler.JWTManager.Validate(token)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Unauthorized", nil, err)
		return
	}

	if err := context.ShouldBindJSON(&patient); err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Input", nil, err)
		return
	}

	patient.UserId = claims.Sub
	createdPatient, err := handler.PatientService.CreatePatient(&patient)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetPatientResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to create patient", nil, err)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetPatientResource(), map[string]interface{}{"patient": createdPatient})
	handler.ResponseManager.Send(context, http.StatusOK, "Patient Created", createdPatient, nil)
}

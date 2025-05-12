package resources

type ResourceType string

const (
	Authentication ResourceType = "authentication"
	User           ResourceType = "user"
	Patient        ResourceType = "patient"
	PatientImages  ResourceType = "patient_images"
)

type ResourceManagerInterface interface {
	GetAuthenticationResource() ResourceType
	GetUserResource() ResourceType
	GetPatientResource() ResourceType
	GetPatientImagesResource() ResourceType
	ValidateResource(resource ResourceType) bool
}

type ResourceManager struct {
	Authentication ResourceType
	User           ResourceType
	Patient        ResourceType
	PatientImages  ResourceType
}

func NewResourceManager() ResourceManagerInterface {
	return &ResourceManager{
		Authentication: Authentication,
		User:           User,
		Patient:        Patient,
		PatientImages:  PatientImages,
	}
}

func (rm *ResourceManager) ValidateResource(resource ResourceType) bool {
	switch resource {
	case rm.Authentication, User, Patient, PatientImages: // Add more valid resources here
		return true
	default:
		return false
	}
}

func (rm *ResourceManager) GetAuthenticationResource() ResourceType {
	return rm.Authentication
}

func (rm *ResourceManager) GetUserResource() ResourceType {
	return rm.User
}

func (rm *ResourceManager) GetPatientResource() ResourceType {
	return rm.Patient
}

func (rm *ResourceManager) GetPatientImagesResource() ResourceType {
	return rm.PatientImages
}

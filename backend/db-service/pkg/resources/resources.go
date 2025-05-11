package resources

type ResourceType string

const (
	Authentication ResourceType = "authentication"
	User           ResourceType = "user"
	Patient        ResourceType = "patient"
)

type ResourceManagerInterface interface {
	GetAuthenticationResource() ResourceType
	GetUserResource() ResourceType
	GetPatientResource() ResourceType
	ValidateResource(resource ResourceType) bool
}

type ResourceManager struct {
	// Add other resources here if needed
	Authentication ResourceType
	User           ResourceType
	Patient        ResourceType
}

func NewResourceManager() ResourceManagerInterface {
	return &ResourceManager{
		Authentication: Authentication,
		User:           User,
		Patient:        Patient,
	}
}

func (rm *ResourceManager) ValidateResource(resource ResourceType) bool {
	switch resource {
	case rm.Authentication, User, Patient: // Add more valid resources here
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

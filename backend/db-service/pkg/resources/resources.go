package resources

type ResourceType string

const (
	Authentication ResourceType = "authentication"
	User           ResourceType = "user"
)

type ResourceManagerInterface interface {
	GetAuthenticationResource() ResourceType
	GetUserResource() ResourceType
	ValidateResource(resource ResourceType) bool
}

type ResourceManager struct {
	// Add other resources here if needed
	Authentication ResourceType
	User           ResourceType
}

func NewResourceManager() ResourceManagerInterface {
	return &ResourceManager{
		Authentication: Authentication,
		User:           User,
	}
}

func (rm *ResourceManager) ValidateResource(resource ResourceType) bool {
	switch resource {
	case rm.Authentication, User: // Add more valid resources here
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

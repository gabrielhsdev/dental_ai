from app.config import Config
from app.repositories.ImageRepository import ImageRepository
from app.repositories.ModelRepository import ModelRepository
from app.services.ProcessingServiceFactory import ProcessingServiceFactory

class ApiHandler:
    """Base handler for API requests"""
    
    def __init__(self, config: Config):
        self.config = config
        self.model_repository = ModelRepository(config)
        self.image_repository = ImageRepository(config)
        self.service_factory = ProcessingServiceFactory(self.model_repository, config)
    
    def allowed_file(self, filename: str) -> bool:
        """Check if file has an allowed extension"""
        return '.' in filename and \
               filename.rsplit('.', 1)[1].lower() in self.config.ALLOWED_EXTENSIONS

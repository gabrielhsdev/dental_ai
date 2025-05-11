from app.services.ImageProcessorService import ImageProcessorService
from app.repositories.ModelRepository import ModelRepository
from app.config.Config import Config
from app.services.ClassificationService import ClassificationService
from app.services.DetectionService import DetectionService
from app.services.SegmentationService import SegmentationService

class ProcessingServiceFactory:
    """Factory for creating the appropriate image processing service"""
    
    def __init__(self, model_repository: ModelRepository, config: Config):
        self.model_repository = model_repository
        self.config = config
    
    def create_service(self, model_id: str) -> ImageProcessorService:
        """Create an image processing service for the given model ID"""
        model = self.model_repository.get_model(model_id)
        
        if model_id == "classification":
            return ClassificationService(model, self.config)
        elif model_id == "detection":
            return DetectionService(model, self.config)
        elif model_id == "segmentation":
            return SegmentationService(model, self.config)
        else:
            raise ValueError(f"Unknown model ID: {model_id}")

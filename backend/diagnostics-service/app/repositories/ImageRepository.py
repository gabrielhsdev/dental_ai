import os
import uuid
from PIL import Image
from app.config import Config
from werkzeug.utils import secure_filename

class ImageRepository:
    """Repository for handling image storage and retrieval"""
    
    def __init__(self, config: Config):
        self.config = config
    
    def save_upload(self, file) -> str:
        """Save an uploaded image and return its path"""
        filename = secure_filename(file.filename)
        unique_filename = f"{uuid.uuid4()}_{filename}"
        filepath = os.path.join(self.config.UPLOAD_FOLDER, unique_filename)
        file.save(filepath)
        return filepath
    
    def save_result(self, image: Image.Image) -> str:
        """Save a result image and return its filename"""
        result_filename = f"{uuid.uuid4()}.png"
        result_path = os.path.join(self.config.RESULTS_FOLDER, result_filename)
        image.save(result_path)
        return result_filename
    
    def get_result(self, filename: str) -> str:
        """Get the path to a result image"""
        return os.path.join(self.config.RESULTS_FOLDER, filename)

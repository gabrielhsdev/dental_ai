import numpy as np
import cv2
from PIL import Image
from typing import Dict, Any
from typing import Dict

from app.models import Model
from app.config import Config
from app.services.ImageProcessorService import ImageProcessorService

class SegmentationService(ImageProcessorService):
    """Service for image segmentation"""
    
    def __init__(self, model: Model, config: Config):
        self.model = model
        self.config = config
    
    def process(self, image: Image.Image) -> Image.Image:
        """Segment an image and highlight the teeth"""
        model_instance = self.model.load()
        img = np.asarray(image)
        img_cv = self._convert_one_channel(img)
        img_cv = cv2.resize(img_cv, (512, 512), interpolation=cv2.INTER_LANCZOS4)
        img_cv = np.float32(img_cv / 255)
        img_cv = np.reshape(img_cv, (1, 512, 512, 1))
        
        prediction = model_instance.predict(img_cv)
        predicted = prediction[0]
        predicted = cv2.resize(predicted, (img.shape[1], img.shape[0]), interpolation=cv2.INTER_LANCZOS4)
        
        mask = np.uint8(predicted * 255)
        _, mask = cv2.threshold(mask, thresh=0, maxval=255, type=cv2.THRESH_BINARY + cv2.THRESH_OTSU)
        
        kernel = np.ones((5, 5), dtype=np.float32)
        mask = cv2.morphologyEx(mask, cv2.MORPH_OPEN, kernel, iterations=1)
        mask = cv2.morphologyEx(mask, cv2.MORPH_CLOSE, kernel, iterations=1)
        
        cnts, _ = cv2.findContours(mask, cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)
        
        # Make a writable copy of the image
        img_writable = self._convert_rgb(img).copy()
        output = cv2.drawContours(img_writable, cnts, -1, (255, 0, 0), 3)
        
        return Image.fromarray(output)
    
    def predict_raw(self, image: Image.Image) -> Dict[str, Any]:
        """Return raw prediction data"""
        model_instance = self.model.load()
        img = np.asarray(image)
        img_cv = self._convert_one_channel(img)
        img_cv = cv2.resize(img_cv, (512, 512), interpolation=cv2.INTER_LANCZOS4)
        img_cv = np.float32(img_cv / 255)
        img_cv = np.reshape(img_cv, (1, 512, 512, 1))
        
        prediction = model_instance.predict(img_cv)
        
        # Convert mask to binary format for storage
        predicted = prediction[0]
        mask = np.uint8(predicted * 255)
        _, binary_mask = cv2.threshold(mask, thresh=0, maxval=1, type=cv2.THRESH_BINARY + cv2.THRESH_OTSU)
        
        # Calculate statistics
        total_pixels = binary_mask.size
        segmented_pixels = np.sum(binary_mask)
        segmentation_ratio = float(segmented_pixels / total_pixels)
        
        # Find contours for counting teeth
        resized_mask = cv2.resize(binary_mask, (img.shape[1], img.shape[0]), interpolation=cv2.INTER_LANCZOS4)
        kernel = np.ones((5, 5), dtype=np.float32)
        processed_mask = cv2.morphologyEx(resized_mask, cv2.MORPH_CLOSE, kernel, iterations=1)
        cnts, _ = cv2.findContours(np.uint8(processed_mask * 255), cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
        
        return {
            "segmentation_ratio": segmentation_ratio,
            "segmented_pixels": int(segmented_pixels),
            "total_pixels": int(total_pixels),
            "contour_count": len(cnts)
        }
    
    def _convert_one_channel(self, img: np.ndarray) -> np.ndarray:
        """Convert image to grayscale if needed"""
        if len(img.shape) > 2:
            img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
        return img
    
    def _convert_rgb(self, img: np.ndarray) -> np.ndarray:
        """Convert image to RGB if needed"""
        if len(img.shape) == 2:
            img = cv2.cvtColor(img, cv2.COLOR_GRAY2RGB)
        return img

from tensorflow.keras.preprocessing import image as image_processor # type: ignore
from tensorflow.keras.applications.vgg16 import preprocess_input # type: ignore
from PIL import Image, ImageDraw, ImageFont
from typing import Dict, Any
from typing import Dict
import numpy as np

from app.models import Model
from app.config import Config
from app.services.ImageProcessorService import ImageProcessorService

class ClassificationService(ImageProcessorService):
    """Service for image classification"""
    
    def __init__(self, model: Model, config: Config):
        self.model = model
        self.config = config
    
    def process(self, image: Image.Image) -> Image.Image:
        """Classify an image and annotate it with the result"""
        model_instance = self.model.load()
        img = image.resize((224, 224))
        x = image_processor.img_to_array(img)
        x = np.expand_dims(x, axis=0)
        img_data = preprocess_input(x)
        
        result = model_instance.predict(img_data)
        if result[0][0] > result[0][1]:
            prediction = 'Calculus'
        else:
            prediction = 'Caries'
        
        # Draw the classification result on the image
        draw = ImageDraw.Draw(image)
        font = ImageFont.truetype(self.config.FONT_DIR, 20)
        text = f"Classified as: {prediction}"
        bbox = draw.textbbox((0, 0), text, font=font)
        text_width = bbox[2] - bbox[0]
        text_height = bbox[3] - bbox[1]
        draw.rectangle([(0, 0), (text_width, text_height)], fill="black")
        draw.text((0, 0), text, fill="white", font=font)
        
        return image

    def predict_raw(self, image: Image.Image) -> Dict[str, Any]:
        """Return raw prediction data"""
        model_instance = self.model.load()
        img = image.resize((224, 224))
        x = image_processor.img_to_array(img)
        x = np.expand_dims(x, axis=0)
        img_data = preprocess_input(x)
        
        result = model_instance.predict(img_data)
        predictions = {
            "calculus": float(result[0][0]),
            "caries": float(result[0][1]),
            "prediction": "Calculus" if result[0][0] > result[0][1] else "Caries"
        }
        
        return predictions

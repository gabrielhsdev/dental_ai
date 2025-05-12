from PIL import Image, ImageDraw, ImageFont
from typing import Dict, Any
from typing import Dict

from app.models import Model
from app.config import Config
from app.services.ImageProcessorService import ImageProcessorService

class DetectionService(ImageProcessorService):
    """Service for object detection"""
    
    def __init__(self, model: Model, config: Config):
        self.model = model
        self.config = config
    
    def process(self, image: Image.Image) -> Image.Image:
        """Detect caries in an image and annotate it with bounding boxes"""
        model_instance = self.model.load()
        results = model_instance.predict(image)
        result = results[0]
        
        # Debug where the font is loaded
        print(f"Font directory: {self.config.FONT_DIR}")
        draw = ImageDraw.Draw(image)
        font = ImageFont.truetype(self.config.FONT_DIR, 8)

        for box in result.boxes:
            x1, y1, x2, y2 = [round(x) for x in box.xyxy[0].tolist()]
            class_id = box.cls[0].item()
            prob = round(box.conf[0].item(), 2)
            label = f"{result.names[class_id]}: {prob}"
            
            draw.rectangle([x1, y1, x2, y2], outline="red", width=2)
            bbox = draw.textbbox((0, 0), label, font=font)
            text_width = bbox[2] - bbox[0]
            text_height = bbox[3] - bbox[1]
            draw.rectangle([(x1, y1 - text_height), (x1 + text_width, y1)], fill="red")
            draw.text((x1, y1 - text_height), label, fill="white", font=font)

        return image
    
    def predict_raw(self, image: Image.Image) -> Dict[str, Any]:
        """Return raw prediction data"""
        model_instance = self.model.load()
        results = model_instance.predict(image)
        result = results[0]
        
        detections = []
        for box in result.boxes:
            x1, y1, x2, y2 = [round(x) for x in box.xyxy[0].tolist()]
            class_id = int(box.cls[0].item())
            confidence = float(box.conf[0].item())
            class_name = result.names[class_id]
            
            detections.append({
                "class_id": class_id,
                "class_name": class_name,
                "confidence": confidence,
                "bbox": [x1, y1, x2, y2]
            })
        
        return {
            "detections": detections,
            "count": len(detections)
        }

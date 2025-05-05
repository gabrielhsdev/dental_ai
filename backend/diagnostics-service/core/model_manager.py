import os
import tensorflow as tf
from tensorflow.keras.models import load_model
from huggingface_hub import from_pretrained_keras
from ultralytics import YOLO

from app.config import Config  # Adjust this if you structure it differently

class ModelManager:
    @staticmethod
    def load_model(model_name: str):
        model_path = os.path.join(Config.MODELS_DIR, Config.MODELS[model_name])
        
        if model_name == "Dental X-Ray Segmentation":
            try:
                return from_pretrained_keras("SerdarHelli/Segmentation-of-Teeth-in-Panoramic-X-ray-Image-Using-U-Net")
            except Exception:
                return tf.keras.models.load_model(model_path)
        
        elif model_name == "Caries Detection":
            return YOLO(model_path)
        
        else:
            return load_model(model_path)

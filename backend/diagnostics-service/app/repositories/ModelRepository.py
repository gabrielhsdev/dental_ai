import os
from typing import Dict, List

from app.config import Config
from app.models.Model import Model
from app.models.KerasModel import KerasModel
from app.models.YoloModel import YoloModel

class ModelRepository:
    """Repository for accessing models"""
    
    def __init__(self, config: Config):
        self.config = config
        self._models = {}
    
    def get_model(self, model_id: str) -> Model:
        """Get a model by ID"""
        if model_id not in self.config.MODEL_CONFIG:
            raise ValueError(f"Unknown model ID: {model_id}")
        
        if model_id not in self._models:
            model_config = self.config.MODEL_CONFIG[model_id]
            model_path = os.path.join(self.config.MODELS_DIR, model_config["file"])
            
            if model_config["type"] == "keras":
                huggingface_id = model_config.get("huggingface_id")
                self._models[model_id] = KerasModel(model_path, huggingface_id)
            elif model_config["type"] == "yolo":
                self._models[model_id] = YoloModel(model_path)
            else:
                raise ValueError(f"Unknown model type: {model_config['type']}")
        
        return self._models[model_id]
    
    def get_available_models(self) -> List[Dict[str, str]]:
        """Get a list of available models"""
        models = []
        for model_id, config in self.config.MODEL_CONFIG.items():
            models.append({
                "id": model_id,
                "name": config["name"],
                "type": config["type"]
            })
        return models
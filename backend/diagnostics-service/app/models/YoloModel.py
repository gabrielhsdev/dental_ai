from typing import Literal
from ultralytics import YOLO
from Model import Model

class YoloModel(Model):
    """Class for YOLO models"""
    
    def __init__(self, model_path: str):
        self.model_path = model_path
        self._model = None
    
    def load(self) -> YOLO:
        """Load a YOLO model from file"""
        if self._model is None:
            try:
                self._model = YOLO(str(self.model_path))
            except Exception as e:
                raise Exception(f"Failed to load YOLO model: {str(e)}")
        return self._model
    
    @property
    def model_type(self) -> Literal["yolo"]:
        return "yolo"
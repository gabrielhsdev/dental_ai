from abc import ABC, abstractmethod
from typing import Literal, Union
from ultralytics import YOLO
import tensorflow as tf

class Model(ABC):
    """Abstract base class for all models"""
    
    @abstractmethod
    def load(self) -> Union[tf.keras.Model, YOLO]:
        """Load the model"""
        pass
    
    @property
    @abstractmethod
    def model_type(self) -> Literal["keras", "yolo"]:
        """Return the model type"""
        pass

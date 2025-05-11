import tensorflow as tf
from typing import Optional, Literal
from tensorflow.keras.models import load_model # type: ignore
from huggingface_hub import from_pretrained_keras
from app.models.Model import Model

class KerasModel(Model):
    """Class for Keras models"""
    
    def __init__(self, model_path: str, huggingface_id: Optional[str] = None):
        self.model_path = model_path
        self.huggingface_id = huggingface_id
        self._model = None
    
    def load(self) -> tf.keras.Model:
        """Load a Keras model from file or Hugging Face"""
        if self._model is None:
            try:
                if self.huggingface_id:
                    self._model = from_pretrained_keras(str(self.huggingface_id))
                else:
                    self._model = load_model(str(self.model_path))
            except Exception as e:
                raise Exception(f"Failed to load Keras model: {str(e)}")
        return self._model
    
    @property
    def model_type(self) -> Literal["keras"]:
        return "keras"

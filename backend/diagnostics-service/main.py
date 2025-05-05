import os
import numpy as np
import cv2
import tensorflow as tf
from flask import Flask, request, jsonify, send_file
from tensorflow.keras.preprocessing import image as image_processor
from tensorflow.keras.applications.vgg16 import preprocess_input
from tensorflow.keras.models import load_model
from PIL import Image, ImageDraw, ImageFont
from ultralytics import YOLO
from huggingface_hub import from_pretrained_keras
from werkzeug.utils import secure_filename
import io
import uuid
from abc import ABC, abstractmethod
from typing import Dict, List, Any, Optional, Tuple
from typing import Set, Dict

# ============================================================
# Configuration
# ============================================================

class Config:
    """Application configuration"""
    BASE_DIR: str = os.path.dirname(os.path.abspath(__file__))
    ASSETS_DIR: str = os.path.join(BASE_DIR, 'assets')
    MODELS_DIR: str = os.path.join(ASSETS_DIR, 'models')
    FONT_DIR: str = os.path.join(ASSETS_DIR, 'arial.ttf')
    REQUEST_FOLDER: str = os.path.join(ASSETS_DIR, 'results')
    UPLOAD_FOLDER: str = os.path.join(REQUEST_FOLDER, 'uploads')
    RESULTS_FOLDER: str = os.path.join(REQUEST_FOLDER, 'results')
    ALLOWED_EXTENSIONS: Set[str] = {'png', 'jpg', 'jpeg'}
    
    # Create necessary directories
    os.makedirs(UPLOAD_FOLDER, exist_ok=True)
    os.makedirs(RESULTS_FOLDER, exist_ok=True)
    
    MODEL_CONFIG: Dict[str, Dict[str, str]] = {
        "classification": {
            "name": "Calculus and Caries Classification",
            "file": "classification.h5",
            "type": "keras"
        },
        "detection": {
            "name": "Caries Detection",
            "file": "detection.pt",
            "type": "yolo"
        },
        "segmentation": {
            "name": "Dental X-Ray Segmentation",
            "file": "dental_xray_seg.h5",
            "type": "keras",
            "huggingface_id": "SerdarHelli/Segmentation-of-Teeth-in-Panoramic-X-ray-Image-Using-U-Net"
        }
    }

# ============================================================
# Models
# ============================================================

class Model(ABC):
    """Abstract base class for all models"""
    
    @abstractmethod
    def load(self) -> Any:
        """Load the model"""
        pass
    
    @property
    @abstractmethod
    def model_type(self) -> str:
        """Return the model type"""
        pass

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
                    self._model = from_pretrained_keras(self.huggingface_id)
                else:
                    self._model = load_model(self.model_path)
            except Exception as e:
                raise Exception(f"Failed to load Keras model: {str(e)}")
        return self._model
    
    @property
    def model_type(self) -> str:
        return "keras"

class YoloModel(Model):
    """Class for YOLO models"""
    
    def __init__(self, model_path: str):
        self.model_path = model_path
        self._model = None
    
    def load(self) -> YOLO:
        """Load a YOLO model from file"""
        if self._model is None:
            try:
                self._model = YOLO(self.model_path)
            except Exception as e:
                raise Exception(f"Failed to load YOLO model: {str(e)}")
        return self._model
    
    @property
    def model_type(self) -> str:
        return "yolo"

# ============================================================
# Repositories
# ============================================================

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

# ============================================================
# Services
# ============================================================

class ImageProcessorService(ABC):
    """Abstract base class for image processors"""
    
    @abstractmethod
    def process(self, image: Image.Image) -> Image.Image:
        """Process an image and return the result"""
        pass

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

# ============================================================
# Handlers
# ============================================================

class ApiHandler:
    """Base handler for API requests"""
    
    def __init__(self, config: Config):
        self.config = config
        self.model_repository = ModelRepository(config)
        self.image_repository = ImageRepository(config)
        self.service_factory = ProcessingServiceFactory(self.model_repository, config)
    
    def allowed_file(self, filename: str) -> bool:
        """Check if file has an allowed extension"""
        return '.' in filename and \
               filename.rsplit('.', 1)[1].lower() in self.config.ALLOWED_EXTENSIONS

class ModelsHandler(ApiHandler):
    """Handler for model-related endpoints"""
    
    def get_models(self):
        """Get a list of available models"""
        try:
            models = self.model_repository.get_available_models()
            return jsonify({
                "success": True,
                "models": models
            }), 200
        except Exception as e:
            return jsonify({
                "success": False,
                "error": str(e)
            }), 500

class ImageProcessingHandler(ApiHandler):
    """Handler for image processing endpoints"""
    
    def process_image(self, model_id: str):
        """Process an image with the specified model"""
        # Check if the request has a file
        if 'file' not in request.files:
            return jsonify({
                "success": False,
                "error": "No file part in the request"
            }), 400
        
        file = request.files['file']
        
        # Check if a file was selected
        if file.filename == '':
            return jsonify({
                "success": False,
                "error": "No file selected"
            }), 400
        
        # Check if the file has an allowed extension
        if not file or not self.allowed_file(file.filename):
            return jsonify({
                "success": False,
                "error": "File type not allowed"
            }), 400
        
        try:
            # Save the uploaded file
            file_path = self.image_repository.save_upload(file)
            
            # Open the image
            image = Image.open(file_path)
            
            # Get the appropriate service
            service = self.service_factory.create_service(model_id)
            
            # Process the image
            result_image = service.process(image)
            
            # Save the result
            result_filename = self.image_repository.save_result(result_image)
            
            # Return the result
            return jsonify({
                "success": True,
                "result_image": f"/api/results/{result_filename}",
                "model_id": model_id
            }), 200
            
        except Exception as e:
            return jsonify({
                "success": False,
                "error": str(e)
            }), 500
    
    def predict(self, model_id: str):
        """Get raw predictions from a model for an image"""
        # Check if the request has a file
        if 'file' not in request.files:
            return jsonify({
                "success": False,
                "error": "No file part in the request"
            }), 400
        
        file = request.files['file']
        
        # Check if a file was selected
        if file.filename == '':
            return jsonify({
                "success": False,
                "error": "No file selected"
            }), 400
        
        # Check if the file has an allowed extension
        if not file or not self.allowed_file(file.filename):
            return jsonify({
                "success": False,
                "error": "File type not allowed"
            }), 400
        
        try:
            # Save the uploaded file
            file_path = self.image_repository.save_upload(file)
            
            # Open the image
            image = Image.open(file_path)
            
            # Get the appropriate service
            service = self.service_factory.create_service(model_id)
            
            # Get raw predictions
            predictions = service.predict_raw(image)
            
            # Return the predictions
            return jsonify({
                "success": True,
                "model_id": model_id,
                "predictions": predictions
            }), 200
            
        except Exception as e:
            return jsonify({
                "success": False,
                "error": str(e)
            }), 500
    
    def get_result(self, filename: str):
        """Get a result image"""
        try:
            file_path = self.image_repository.get_result(filename)
            return send_file(file_path, mimetype='image/png')
        except FileNotFoundError:
            return jsonify({
                "success": False,
                "error": "Image not found"
            }), 404
        except Exception as e:
            return jsonify({
                "success": False,
                "error": str(e)
            }), 500

# ============================================================
# Routes
# ============================================================

def setup_routes(app: Flask, config: Config):
    """Set up routes for the Flask app"""
    
    models_handler = ModelsHandler(config)
    image_handler = ImageProcessingHandler(config)
    
    # API routes
    @app.route('/api/models', methods=['GET'])
    def get_models():
        return models_handler.get_models()
    
    @app.route('/api/process/<model_id>', methods=['POST'])
    def process_image(model_id):
        return image_handler.process_image(model_id)
    
    @app.route('/api/predict/<model_id>', methods=['POST'])
    def predict(model_id):
        return image_handler.predict(model_id)
    
    @app.route('/api/results/<filename>', methods=['GET'])
    def get_result(filename):
        return image_handler.get_result(filename)
    
    # Root route
    @app.route('/')
    def index():
        return jsonify({
            "message": "Dental AI API",
            "endpoints": {
                "GET /api/models": "Get a list of available models",
                "POST /api/process/<model_id>": "Process an image with a model",
                "POST /api/predict/<model_id>": "Get raw predictions from a model",
                "GET /api/results/<filename>": "Get a result image"
            }
        })

# ============================================================
# Application
# ============================================================

def create_app():
    """Create and configure the Flask application"""
    app = Flask(__name__)
    
    # Load configuration
    config = Config()
    
    # Set up routes
    setup_routes(app, config)
    
    return app

# Run the application
if __name__ == "__main__":
    app = create_app()
    app.run(debug=True, host='0.0.0.0', port=5123)
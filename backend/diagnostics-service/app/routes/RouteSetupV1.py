from flask import Flask, jsonify
from flask import Blueprint

from app.config import Config
from app.handlers.ModelsHandler import ModelsHandler
from app.handlers.ImageProcessingHandler import ImageProcessingHandler

class RouteSetupV1:
    def __init__(self, app: Flask, config: Config):
        self.app = app
        self.config = config
        self.models_handler = ModelsHandler(config)
        self.image_handler = ImageProcessingHandler(config)
        self.setup_routes()

    def setup_routes(self):
        """Set up routes for the Flask app"""
        # Create v1 blueprint
        v1 = Blueprint('v1', __name__, url_prefix='/api/v1')
        
        # Define routes
        v1.route('/models', methods=['GET'])(self.get_models)
        v1.route('/process/<model_id>', methods=['POST'])(self.process_image)
        v1.route('/predict/<model_id>', methods=['POST'])(self.predict)
        v1.route('/results/<filename>', methods=['GET'])(self.get_result)
        
        # Register blueprint
        self.app.register_blueprint(v1)
        
        # Root route
        self.app.route('/')(self.index)

    def get_models(self):
        return self.models_handler.get_models()
    
    def process_image(self, model_id):
        return self.image_handler.process_image(model_id)
    
    def predict(self, model_id):
        return self.image_handler.predict(model_id)
    
    def get_result(self, filename):
        return self.image_handler.get_result(filename)
    
    def index(self):
        return jsonify({
            "message": "Dental AI API",
            "endpoints": {
                "GET /api/v1/models": "Get a list of available models",
                "POST /api/v1/process/<model_id>": "Process an image with a model",
                "POST /api/v1/predict/<model_id>": "Get raw predictions from a model",
                "GET /api/v1/results/<filename>": "Get a result image"
            }
        })
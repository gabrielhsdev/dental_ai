from flask import Flask

# ============================================================
# Configs
# ============================================================
from app.config.Config import Config

# ============================================================
# Models
# ============================================================
from app.models.Model import Model
from app.models.KerasModel import KerasModel
from app.models.YoloModel import YoloModel

# ============================================================
# Repositories
# ============================================================
from app.repositories.ModelRepository import ModelRepository
from app.repositories.ImageRepository import ImageRepository

# ============================================================
# Services
# ============================================================
from app.services.ImageProcessorService import ImageProcessorService
from app.services.ClassificationService import ClassificationService
from app.services.DetectionService import DetectionService
from app.services.SegmentationService import SegmentationService
from app.services.ProcessingServiceFactory import ProcessingServiceFactory

# ============================================================
# Handlers
# ============================================================
from app.handlers.ApiHandler import ApiHandler
from app.handlers.ModelsHandler import ModelsHandler
from app.handlers.ImageProcessingHandler import ImageProcessingHandler

# ============================================================
# Routes
# ============================================================
from app.routes.RouteSetupV1 import RouteSetupV1

# ============================================================
# Application
# ============================================================

def create_app():
    """Create and configure the Flask application"""
    app = Flask(__name__)
    config = Config()
    
    # Set up routes
    route_setup = RouteSetupV1(app, config)
    route_setup.setup_routes()
    
    return app

# Run the application
if __name__ == "__main__":
    app = create_app()
    app.run(debug=True, host='0.0.0.0', port=5123)
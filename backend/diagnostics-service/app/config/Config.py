from ast import Dict, Set
import os


class Config:
    BASE_DIR: str = str(os.path.dirname(os.path.abspath(__file__)))
    ASSETS_DIR: str = str(os.path.join(BASE_DIR, 'assets'))
    MODELS_DIR: str = str(os.path.join(ASSETS_DIR, 'models'))
    FONT_DIR: str = str(os.path.join(ASSETS_DIR, 'arial.ttf'))
    REQUEST_FOLDER: str = str(os.path.join(ASSETS_DIR, 'results'))
    UPLOAD_FOLDER: str = str(os.path.join(REQUEST_FOLDER, 'uploads'))
    RESULTS_FOLDER: str = str(os.path.join(REQUEST_FOLDER, 'results'))
    ALLOWED_EXTENSIONS: Set[str] = set({'png', 'jpg', 'jpeg'})
    
    # Create necessary directories
    os.makedirs(UPLOAD_FOLDER, exist_ok=True)
    os.makedirs(RESULTS_FOLDER, exist_ok=True)
    
    MODEL_CONFIG: Dict[str, Dict[str, str]] = {
        "classification": {
            "name": str("Calculus and Caries Classification"),
            "file": str("classification.h5"),
            "type": str("keras")
        },
        "detection": {
            "name": str("Caries Detection"),
            "file": str("detection.pt"),
            "type": str("yolo")
        },
        "segmentation": {
            "name": str("Dental X-Ray Segmentation"),
            "file": str("dental_xray_seg.h5"),
            "type": str("keras"),
            "huggingface_id": str("SerdarHelli/Segmentation-of-Teeth-in-Panoramic-X-ray-Image-Using-U-Net")
        }
    }

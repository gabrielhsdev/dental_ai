import os

class Config:
    ASSETS_DIR = os.path.join(os.path.dirname(__file__), '../assets')
    MODELS_DIR = os.path.join(ASSETS_DIR, 'models')
    FONT_DIR = os.path.join(ASSETS_DIR, 'arial.ttf')
    MODELS = {
        "Calculus and Caries Classification": "classification.h5",
        "Caries Detection": "detection.pt",
        "Dental X-Ray Segmentation": "dental_xray_seg.h5"
    }
    EXAMPLES = {
        "Calculus and Caries Classification": os.path.join(ASSETS_DIR, 'classification'),
        "Caries Detection": os.path.join(ASSETS_DIR, 'detection'),
        "Dental X-Ray Segmentation": os.path.join(ASSETS_DIR, 'segmentation')
    }
    # Debug our path, if needed
    print(f"Assets Directory: {ASSETS_DIR}")
    print(f"Models Directory: {MODELS_DIR}")
    print(f"Font Directory: {FONT_DIR}")
    print(f"Examples Directory: {EXAMPLES}")
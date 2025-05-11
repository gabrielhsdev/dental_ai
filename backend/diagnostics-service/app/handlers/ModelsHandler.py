from flask import jsonify
from app.handlers.ApiHandler import ApiHandler

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
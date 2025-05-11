from app.handlers.ApiHandler import ApiHandler
from flask import jsonify, request, send_file
from PIL import Image

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

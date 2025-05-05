from core.model_manager import ModelManager
from app.config import Config
from tensorflow.keras.preprocessing import image as image_processor
import numpy as np
from tensorflow.keras.applications.vgg16 import preprocess_input
from PIL import Image, ImageDraw, ImageFont
import cv2

class ImageProcessor:
    def process_image(self, image: Image.Image, model_name: str):
        if model_name == "Calculus and Caries Classification":
            return self.classify_image(image, model_name)
        elif model_name == "Caries Detection":
            return self.detect_caries(image)
        elif model_name == "Dental X-Ray Segmentation":
            return self.segment_dental_xray(image)

    def classify_image(self, image: Image.Image, model_name: str):
        model = ModelManager.load_model(model_name)
        img = image.resize((224, 224))
        x = image_processor.img_to_array(img)
        x = np.expand_dims(x, axis=0)
        img_data = preprocess_input(x)
        result = model.predict(img_data)
        if result[0][0] > result[0][1]:
            prediction = 'Calculus'
        else:
            prediction = 'Caries'
        
        # Draw the classification result on the image
        draw = ImageDraw.Draw(image)
        font = ImageFont.truetype(Config.FONT_DIR, 20)
        text = f"Classified as: {prediction}"
        bbox = draw.textbbox((0, 0), text, font=font)
        text_width = bbox[2] - bbox[0]
        text_height = bbox[3] - bbox[1]
        draw.rectangle([(0, 0), (text_width, text_height)], fill="black")
        draw.text((0, 0), text, fill="white", font=font)
        
        return image

    def detect_caries(self, image: Image.Image):
        model = ModelManager.load_model("Caries Detection")
        results = model.predict(image)
        result = results[0]
        draw = ImageDraw.Draw(image)
        font = ImageFont.truetype(Config.FONT_DIR, 20)

        for box in result.boxes:
            x1, y1, x2, y2 = [round(x) for x in box.xyxy[0].tolist()]
            class_id = box.cls[0].item()
            prob = round(box.conf[0].item(), 2)
            label = f"{result.names[class_id]}: {prob}"
            draw.rectangle([x1, y1, x2, y2], outline="red", width=2)
            # text_width, text_height = draw.textsize(label, font=font)
            bbox = draw.textbbox((0, 0), label, font=font)
            text_width = bbox[2] - bbox[0]
            text_height = bbox[3] - bbox[1]
            draw.rectangle([(x1, y1 - text_height), (x1 + text_width, y1)], fill="red")
            draw.text((x1, y1 - text_height), label, fill="white", font=font)

        return image

    def segment_dental_xray(self, image: Image.Image):
        model = ModelManager.load_model("Dental X-Ray Segmentation")
        img = np.asarray(image)
        img_cv = self.convert_one_channel(img)
        img_cv = cv2.resize(img_cv, (512, 512), interpolation=cv2.INTER_LANCZOS4)
        img_cv = np.float32(img_cv / 255)
        img_cv = np.reshape(img_cv, (1, 512, 512, 1))
        prediction = model.predict(img_cv)
        predicted = prediction[0]
        predicted = cv2.resize(predicted, (img.shape[1], img.shape[0]), interpolation=cv2.INTER_LANCZOS4)
        mask = np.uint8(predicted * 255)
        _, mask = cv2.threshold(mask, thresh=0, maxval=255, type=cv2.THRESH_BINARY + cv2.THRESH_OTSU)
        kernel = np.ones((5, 5), dtype=np.float32)
        mask = cv2.morphologyEx(mask, cv2.MORPH_OPEN, kernel, iterations=1)
        mask = cv2.morphologyEx(mask, cv2.MORPH_CLOSE, kernel, iterations=1)
        cnts, _ = cv2.findContours(mask, cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)
        
        # Make a writable copy of the image
        img_writable = self.convert_rgb(img).copy()
        output = cv2.drawContours(img_writable, cnts, -1, (255, 0, 0), 3)
        return Image.fromarray(output)

    def convert_one_channel(self, img):
        if len(img.shape) > 2:
            img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
        return img

    def convert_rgb(self, img):
        if len(img.shape) == 2:
            img = cv2.cvtColor(img, cv2.COLOR_GRAY2RGB)
        return img

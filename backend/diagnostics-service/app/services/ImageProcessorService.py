from abc import ABC, abstractmethod
from PIL import Image

class ImageProcessorService(ABC):
    """Abstract base class for image processors"""
    
    @abstractmethod
    def process(self, image: Image.Image) -> Image.Image:
        """Process an image and return the result"""
        pass
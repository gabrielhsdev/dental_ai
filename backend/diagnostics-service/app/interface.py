from core.image_processor import ImageProcessor
from app.config import Config
import os
import gradio as gr

class GradioInterface:
    def __init__(self):
        self.image_processor = ImageProcessor()
        self.preloaded_examples = self.preload_examples()

    def preload_examples(self):
        preloaded = {}
        for model_name, example_dir in Config.EXAMPLES.items():
            examples = [os.path.join(example_dir, img) for img in os.listdir(example_dir)]
            preloaded[model_name] = examples
        return preloaded

    def create_interface(self):
        app_styles = """
        <style>
            /* Global Styles */
            body, #root {
                font-family: Helvetica, Arial, sans-serif;
                background-color: #1a1a1a;
                color: #fafafa;
            }
            /* Header Styles */
            .app-header {
                background: linear-gradient(45deg, #1a1a1a 0%, #333333 100%);
                padding: 24px;
                border-radius: 8px;
                margin-bottom: 24px;
                text-align: center;
            }
            .app-title {
                font-size: 48px;
                margin: 0;
                color: #fafafa;
            }
            .app-subtitle {
                font-size: 24px;
                margin: 8px 0 16px;
                color: #fafafa;
            }
            .app-description {
                font-size: 16px;
                line-height: 1.6;
                opacity: 0.8;
                margin-bottom: 24px;
            }
            /* Button Styles */
            .publication-links {
                display: flex;
                justify-content: center;
                flex-wrap: wrap;
                gap: 8px;
                margin-bottom: 16px;
            }
            .publication-link {
                display: inline-flex;
                align-items: center;
                padding: 8px 16px;
                background-color: #333;
                color: #fff !important;
                text-decoration: none !important;
                border-radius: 20px;
                font-size: 14px;
                transition: background-color 0.3s;
            }
            .publication-link:hover {
                background-color: #555;
            }
            .publication-link i {
                margin-right: 8px;
            }
            /* Content Styles */
            .content-container {
                background-color: #2a2a2a;
                border-radius: 8px;
                padding: 24px;
                margin-bottom: 24px;
            }
            /* Image Styles */
            .image-preview img {
                max-width: 512px;
                max-height: 512px;  
                margin: 0 auto;
                border-radius: 4px;
                display: block;
                object-fit: contain;  
            }
            /* Control Styles */
            .control-panel {
                background-color: #333;
                padding: 16px;
                border-radius: 8px;
                margin-top: 16px;
            }
            /* Gradio Component Overrides */
            .gr-button {
                background-color: #4a4a4a;
                color: #fff;
                border: none;
                border-radius: 4px;
                padding: 8px 16px;
                cursor: pointer;
                transition: background-color 0.3s;
            }
            .gr-button:hover {
                background-color: #5a5a5a;
            }
            .gr-input, .gr-dropdown {
                background-color: #3a3a3a;
                color: #fff;
                border: 1px solid #4a4a4a;
                border-radius: 4px;
                padding: 8px;
            }
            .gr-form {
                background-color: transparent;
            }
            .gr-panel {
                border: none;
                background-color: transparent;
            }
            /* Override any conflicting styles from Bulma */
            .button.is-normal.is-rounded.is-dark {
                color: #fff !important;
                text-decoration: none !important;
            }
        </style>
        """

        header_html = f"""
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css">
        <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.15.4/css/all.css">
        {app_styles}
        <div class="app-header">
            <h1 class="app-title">AI in Dentistry</h1>
            <h2 class="app-subtitle"> Advancing Imaging and Clinical Transcription</h2>
            <p class="app-description">
                This application demonstrates the use of AI in dentistry for tasks such as classification, detection, and segmentation.
            </p>
        </div>
        """

        def process_image(image, model_name):
            result = self.image_processor.process_image(image, model_name)
            return result

        def update_examples(model_name):
            examples = self.preloaded_examples[model_name]
            return gr.Dataset(samples=[[example] for example in examples])

        with gr.Blocks() as demo:
            gr.HTML(header_html)
            with gr.Row(elem_classes="content-container"):
                with gr.Column():
                    input_image = gr.Image(label="Input Image", type="pil", format="png", elem_classes="image-preview")
                    with gr.Row(elem_classes="control-panel"):
                        model_name = gr.Dropdown(
                            label="Model",
                            choices=list(Config.MODELS.keys()),
                            value="Calculus and Caries Classification",
                        )
                    examples = gr.Examples(
                        inputs=input_image,
                        examples=self.preloaded_examples["Calculus and Caries Classification"],
                    )
                with gr.Column():
                    result = gr.Image(label="Result", elem_classes="image-preview")
                    run_button = gr.Button("Run", elem_classes="gr-button")

            model_name.change(
                fn=update_examples,
                inputs=model_name,
                outputs=examples.dataset,
            )

            run_button.click(
                fn=process_image,
                inputs=[input_image, model_name],
                outputs=result,
            )

        return demo

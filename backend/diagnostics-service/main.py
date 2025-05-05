from app.interface import GradioInterface

def main():
    interface = GradioInterface()
    demo = interface.create_interface()
    demo.launch(share=False)

if __name__ == "__main__":
    main()
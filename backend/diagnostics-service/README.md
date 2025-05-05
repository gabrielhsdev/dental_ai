1. **Create a virtual environment** (inside a folder named `.venv`):

   ```bash
   python3 -m venv .venv
   ```

2. **Activate the virtual environment**:

   * On **Linux/macOS**:

     ```bash
     source .venv/bin/activate
     ```

   * On **Windows**:

     ```bash
     .venv\Scripts\activate
     ```

3. **Install dependencies from `requirements.txt`**:

   ```bash
   pip install -r requirements.txt
   ```

4. **Run the pythom main.py**:

   ```bash
   python main.py
   ```

5. **(Optional) Deactivate the virtual environment**:

   ```bash
   deactivate
   ```
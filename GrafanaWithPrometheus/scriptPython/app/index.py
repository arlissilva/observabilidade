import os
import time
from flask import Flask, jsonify

app = Flask(__name__)

def create_edit_delete_files():
    for i in range(1000):
        filename1 = f"file1_{i}.txt"
        filename2 = f"file2_{i}.txt"
        
        # Criação dos arquivos
        with open(filename1, 'w') as f1, open(filename2, 'w') as f2:
            f1.write("Some content here\n" * 1000)
            f2.write("Some content here\n" * 1000)
        
        # Edição dos arquivos
        with open(filename1, 'a') as f1, open(filename2, 'a') as f2:
            f1.write("Appending some more content\n" * 1000)
            f2.write("Appending some more content\n" * 1000)
        
        # Exclusão dos arquivos
        os.remove(filename1)
        os.remove(filename2)
        
        # Pequena pausa para simular uso contínuo
        time.sleep(0.1)

@app.route('/health', methods=['GET'])
def health_check():
    return jsonify(status="OK"), 200

if __name__ == "__main__":
    from threading import Thread
    file_thread = Thread(target=create_edit_delete_files)
    file_thread.start()
    
    app.run(host='0.0.0.0', port=8000)
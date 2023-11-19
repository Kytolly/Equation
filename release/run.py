import os
import subprocess
import webbrowser
import time

folder_path = "../back"
os.chdir(folder_path)
run_process = subprocess.Popen(["go", "run", "main.go"])
time.sleep(2)
webbrowser.open("http://localhost:8080/front")
input("按任意键关闭服务器...\n")
run_process.terminate()
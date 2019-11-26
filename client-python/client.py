import requests
import time

while True:
    time.sleep(1)
    try:
        response = requests.get("http://192.168.25.98:3000/api/getVersion")
        print ("client-python: " + response.text)
    except:
        print ("client-python: Error")
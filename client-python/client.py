import requests
import time
import os

getVersionURL = os.getenv("API_SERVER") + "/api/getVersion"

print("Starting to poll " + getVersionURL)

while True:
    time.sleep(1)
    try:
        response = requests.get(getVersionURL)
        print ("client-python: " + response.text)
    except:
        print ("client-python: Error")
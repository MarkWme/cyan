FROM python:alpine
ADD client.py requirements.txt /
RUN python3 -m pip install -r requirements.txt
CMD [ "python", "-u", "./client.py" ]
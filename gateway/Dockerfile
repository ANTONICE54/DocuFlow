FROM python:3.10-slim

WORKDIR /app

COPY ./gateway .

RUN pip3 install --upgrade pip && pip install --no-cache-dir -r /app/requirments.txt


EXPOSE 5000

CMD ["gunicorn", "gateway_server:docu_flow_app", "-b", "0.0.0.0:5000", "-w", "4"]



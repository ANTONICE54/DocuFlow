FROM python:3.10-slim

WORKDIR /app

COPY ./python_api .

RUN pip3 install --upgrade pip && pip install --no-cache-dir -r /app/requirments.txt


EXPOSE 5001

CMD ["gunicorn", "-k", "uvicorn.workers.UvicornWorker", "--bind", "0.0.0.0:5001", "--log-level", "debug", "DocuFlow:api_app"]




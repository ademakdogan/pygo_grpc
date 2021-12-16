#-----------------------
FROM python:3.8-slim-buster
RUN apt-get update && apt-get install make
RUN apt-get install gcc -y
WORKDIR /opt/pygo_grpc
COPY . .
RUN make install
CMD ["python","src/server/server.py"]
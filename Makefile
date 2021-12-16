

install:
	@pip3 install --default-timeout=900 -r requirements.txt

clean:
	@rm -rf src/__pycache__
	@rm -rf src/models/__pycache__
	@rm -rf src/predictors/__pycache__
	@rm -rf src/utils/__pycache__

docker:
	@docker build . -t pygo_grpc

docker_run:
	@docker run  --rm -t -p ${p}:9001 pygo_grpc

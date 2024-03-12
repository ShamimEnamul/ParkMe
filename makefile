IMG_NAME:=nginx
IMG_TAG:=1.0

docker_build:
	@ echo "Building docker image ${IMG}"
	@ docker build -t ${IMG_NAME}:${IMG_TAG} .

# docker_push:
# 	@ docker push ${IMG_NAME}:${IMG_TAG}
#
 docker_login:
 	@ docker login

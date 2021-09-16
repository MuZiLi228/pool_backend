#!/bin/sh 
  
readonly docker_addr=docker.test.com #docker.isecsp.com  #nexus的docker镜像仓库地址 
readonly slu_name=nft
readonly app_name=pool_backend
readonly env=master
readonly tag=latest #c0070119
readonly port=7001
readonly file_path=/home/linkai/data/workspace/goWork/src/pool_backend

# docker pull ${docker_addr}/${slu_name}/${app_name}-${env}:${tag}
docker stop ${slu_name}-${app_name}
docker rm ${slu_name}-${app_name}
docker run --restart=always -d \
--network host \
--name=${slu_name}-${app_name} \
-e CONSUL_HOST=192.168.1.7 \
-e CONSUL_PORT=8500 \
-e CONSUL_CONFIG_PATH=test/test_config \
-v ${file_path}/configs:/data/app/configs \
-v ${file_path}/logs:/data/app/logs \
${docker_addr}/${slu_name}/${app_name}-${env}:${tag}
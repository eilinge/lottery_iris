# docker 云端部署

## nginx.conf

    index.html: /usr/share/nginx/html/index.html
    ```
    server {
        location / {
            root /usr/share/nginx/html
            index index.html
        }
    }    
    ```

## Dockerfile

    ```
    FROM nginx:alpine

    COPY ./nginx.conf /etc/nginx/conf.d/default.conf
    COPY ../dist /usr/share/nginx/html
    ```

## Aliyun Ghost

    ```
    sudo docker login --username=eilinge registry.cn-shanghai.aliyuncs.com
    sudo docker tag [ImageId] registry.cn-shanghai.aliyuncs.com/eilingeloveduzi/eilinge:[镜像版本号]
    sudo docker push registry.cn-shanghai.aliyuncs.com/eilingeloveduzi/eilinge:[镜像版本号]
    ```

## build.sh

    ```
    npm run build
    cp -r ../dist ./
    docker build -t registry.cn-shanghai.aliyuncs.com/eilingeloveduzi/lottery:v0.0.1
    docker push registry.cn-shanghai.aliyuncs.com/eilingeloveduzi/lottery:v0.0.1
    ```

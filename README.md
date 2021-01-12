docker build -t tallongsun/slow-http-server:0.1 .
docker login
docker push tallongsun/slow-http-server:0.1

kubectl apply -f deploy.yaml
curl --resolve your.example.com:32134:0.0.0.0 http://your.example.com:32134?ts=3000

# Production API

## Build and Run

### Docker

Network
```bash
docker network create production-network
```

Mongo Container
```bash
docker run \
  --name production-mongo \
  --network production-network \
  -p 27017:27017 \
  -d \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=123 \
  mongo:6.0.13
```

RabbitMQ Container
```bash
docker run \
  --name production-rabbitmq \
  --network production-network \
  -p 5672:5672 \
  -p 15672:15672 \
  -d \
  rabbitmq:3.12-management
```

API Image
```bash
docker build -t production-api:latest .
```

API Container
```bash
docker run \
  --name=production-api \
  --network=production-network \
  -p 8080:8080 \
  -d \
  -e MONGO_HOST=production-mongo \
  -e MONGO_PORT=27017 \
  -e MONGO_USERNAME=root \
  -e MONGO_PASSWORD=123 \
  -e RABBITMQ_HOST=production-rabbitmq \
  -e RABBITMQ_PORT=5672 \
  -e RABBITMQ_USERNAME=guest \
  -e RABBITMQ_PASSWORD=guest \
  production-api
```

### Docker Compose
```bash
docker-compose up -d
```

### Kubernetes

#### Secrets DB
```bash
kubectl create secret generic production-mongo \
  --from-literal=username=CHANGE_HERE \
  --from-literal=password=CHANGE_HERE
```

#### Secrets RabbitMQ
```bash
kubectl create secret generic production-rabbitmq \
  --from-literal=username=CHANGE_HERE \
  --from-literal=password=CHANGE_HERE
```

#### Mongo Pods and Services
```bash
kubectl apply -f kubernetes/db/pod.yaml
kubectl apply -f kubernetes/db/service.yaml
kubectl apply -f kubernetes/db/nodeport.yaml # Optional to local access
```

#### RabbitMQ Pods and Services
```bash
kubectl apply -f kubernetes/rabbitmq/pod.yaml
kubectl apply -f kubernetes/rabbitmq/service.yaml
kubectl apply -f kubernetes/rabbitmq/nodeport.yaml # Optional to local access
kubectl apply -f kubernetes/rabbitmq/load-balancer-service.yaml # Optional to local access
kubectl apply -f kubernetes/rabbitmq/load-balancer-service-2.yaml # Optional to local access
```

#### API Pods and Services
```bash
kubectl apply -f kubernetes/api/deployment.yaml
kubectl apply -f kubernetes/api/load-balancer-service.yaml
kubectl apply -f kubernetes/api/service.yaml
```


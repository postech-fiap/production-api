# Production API

## Build and Run

### Docker

Network
```bash
docker network create production-network
```

Mongo Container
```bash
docker run --name=production-mongo --network=production-network -p 27017:27017 -d \
  -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123 \
  mongo:6.0.13
```

API Image
```bash
docker build -t production-api:latest .
```

API Container
```bash
docker run --name=production-api --network=production-network -p 8080:8080 -d \
  -e DATABASE_MONGO_HOST=production-mongo -e DATABASE_MONGO_PORT=27017 -e DATABASE_MONGO_USERNAME=root -e DATABASE_MONGO_PASSWORD=123 \
  production-api
```

### Docker Compose
```bash
docker-compose up -d
```

### Kubernetes

#### Secrets
```bash
kubectl create secret generic production-db --from-literal=mongo-username=CHANGE_HERE --from-literal=mongo-password=CHANGE_HERE
```

#### Mongo Pods and Services
```bash
kubeclt apply -f kubernetes/db/pod.yaml
kubeclt apply -f kubernetes/db/service.yaml
kubeclt apply -f kubernetes/db/nodeport.yaml # Optional to local access
```

#### API Pods and Services
```bash
kubeclt apply -f kubernetes/api/deployment.yaml
kubeclt apply -f kubernetes/api/load-balancer-service.yaml
kubeclt apply -f kubernetes/api/service.yaml
```


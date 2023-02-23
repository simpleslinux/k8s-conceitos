# Containers

```
docker build -t teste-app .
```

```
docker run -p 6377:6377 teste-app
```

```
docker tag teste-app fmnapoli/teste-app 
```


```
docker push fmnapoli/teste-app 
```

```
docker compose  -f "./docker-compose/docker-compose.yaml" up -d  
```

```
docker compose  -f "./docker-compose/minio-docker-compose.yaml" up -d  
```

```
docker compose  -f "./docker-compose/wp-docker-compose.yaml" up -d 
```

# Kubernetes

```
cd k8s
```

```
kubectl run teste-app --image=fmnapoli/teste-app
```

```
kubectl delete pod teste-app
```

```
kubectlrun teste-app --image=fmnapoli/teste-app --dry-run=client -o yaml > pod.yaml
```

```
kubectl apply -f pod.yaml
```

```
kubectl port-forward pods/teste-app 7777:6377
```

```
kubectl delete -f pod.yaml
```

```
kubectl apply -f replicaset.yaml
```

```
kubectl delete -f replicaset.yaml
```

```
kubectl apply -f deployment.yaml
```

```
kubectl expose deployment teste-app-deployment  --port=80 --target-port=6377 --dry-run=client -o yaml
```

```
kubectl apply -f service.yaml
```
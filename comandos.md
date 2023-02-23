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
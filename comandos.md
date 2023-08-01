# Containers

Executando um container
```bash
docker run -p 6378:80 nginx
```

Executando um container e removendo ao final
```bash
docker run --rm -p 6378:80 nginx
```

Executando um container em background
```bash
docker run -d -p 6378:80 nginx
```

Listando containers em execução
```bash
docker container ls
```
Listando todos os containers
```bash
docker container ls -a
```

Executando um container e removendo ao final
```bash
docker run --rm -p 6378:80 nginx
```

Construindo uma Imagem
```bash
docker build -t teste-app .
```

Listando as imagens locais
```bash
docker images
```
Criando um container baseado na imagem criada
```bash
docker run --rm -p 6377:5000 teste-app
```
Colocando uma Tag na imagem
```bash
docker tag teste-app fmnapoli/teste-app 
```

Enviando uma imagem para o Registry

```bash
docker push fmnapoli/teste-app 
```
## Docker Compose

Iniciando um Compose do App Teste

```bash
docker compose  -f "./docker-compose/docker-compose.yaml" up -d
```
Iniciando um Compose MinIO
```bash
docker compose  -f "./docker-compose/minio-docker-compose.yaml" up -d  
```
Iniciando um Compose WordPress
```bash
docker compose  -f "./docker-compose/wp-docker-compose.yaml" up -d 
```

Finalizando um Compose

```bash
docker compose  -f NOMEDOARQUIVO down
```

# Kubernetes

## Criação do cluster K3D

Windows:
```powershell
$env:KUBECONFIG=$("$env:USERPROFILE\.kube\config-devops-lab")
k3d cluster create devops-lab --agents 2 --port "80:80@loadbalancer"
. $profile
```

Linux:
```bash
KUBECONFIG="$HOME/.kube/config-k8s-lab"
k3d cluster create devops-lab --agents 2 --port "80:80@loadbalancer"
source ~/.zshrc
```

```bash
cd k8s
```
## Manupulando o K8S

Verificando o Cluster

```bash
kubectl cluster-info
```

Listando nós do Cluster

```bash
kubectl get nodes
```

Criando um POD
```bash
kubectl run  teste-app  --image=fmnapoli/teste-app
```

Criando um POD passando ENV
```bash
kubectl run  teste-app  --image=fmnapoli/teste-app --env TITULO="K8S POD" --env MENSAGEM="Testando POD"
```

Acessando via encaminhamento de portas

```bash
kubectl port-forward pod/teste-app 5000:5000
```
No navegador: http://localhost:5000

Removendo um POD
```bash
kubectl delete pod teste-app
```

```bash
kubectl run teste-app --image=fmnapoli/teste-app --dry-run=client -o yaml > pod.yaml
```

```bash
kubectl create -f pod.yaml
```

```bash
kubectl port-forward pod/teste-app 5000:5000
```
No navegador: http://localhost:5000

Apagando o POD

```bash
kubectl delete -f pod.yaml
```

```bash
kubectl apply -f replicaset-v1.yaml
```
```bash
kubectl get rs
```
```bash
kubectl delete -f replicaset-v1.yaml
```

```bash
kubectl apply -f replicaset-v2.yaml
```

```bash
kubectl delete -f replicaset-v2.yaml
```

```bash
kubectl apply -f deployment-v1.yaml
```

```bash
kubectl get deployment
```

```bash
kubectl apply -f deployment-v2.yaml
```

```bash
kubectl get deployment
```
```bash
kubectl rollout undo deployment teste-app-deployment
```
```bash
kubectl delete -f deployment-v2.yaml
```


```bash
kubectl apply -f service.yaml
```
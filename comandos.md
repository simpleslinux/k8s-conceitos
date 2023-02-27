# Containers

Listando containers em execução
```bash
docker container ls
```
Listando todos os containers
```bash
docker container ls -a
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
kubectl deploy create  teste-app  --image=fmnapoli/teste-app
```
Removendo um POD
```bash
kubectl delete pod teste-app
```

```bash
kubectlrun teste-app --image=fmnapoli/teste-app --dry-run=client -o yaml > pod.yaml
```

```bash
kubectl apply -f pod.yaml
```

```bash
kubectl port-forward pods/teste-app 7777:5000
```

```bash
kubectl delete -f pod.yaml
```

```bash
kubectl apply -f replicaset.yaml
```

```bash
kubectl delete -f replicaset.yaml
```

```bash
kubectl apply -f deployment.yaml
```

```bash
kubectl expose deployment teste-app-deployment  --port=80 --target-port=6377 --dry-run=client -o yaml
```

```bash
kubectl apply -f service.yaml
```
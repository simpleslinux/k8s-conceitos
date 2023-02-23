# Kuberneter


# Dica

Usar k9s ( https://k9scli.io/ )



Windows

Acrescentar no $profile
```
New-Alias k kubectl.exe
$env:KUBECONFIG=$(Get-ChildItem -Path "$env:USERPROFILE\.kube\config*").FullName -join ";"
```

Linux 

Install k3d
```
wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

Acrescentar no  ~/.zshrc
```
alias k="kubectl"
alias ns="kubens"
alias ctx="kubectx"

source <(kubectl completion zsh)
export KUBECONFIG=$(for i in $(find $HOME/.kube/ -iname 'config*') ; do echo -n ":$i"; done | cut -c 2-)

```


# Criação do Cluster

Create Cluster - Windows

```
$env:KUBECONFIG=$("$env:USERPROFILE\.kube\config-k8s-lab")
k3d cluster create k8s-lab --agents 2 --port "6377:80@loadbalancer"

. $profile
```

Create Cluster - Linux

```
KUBECONFIG="$HOME/.kube/config-k8s-lab"
k3d cluster create k8s-lab --agents 2 --port "6377:80@loadbalancer"

source ~/.zshrc
```
## Instalando o K3D Windows

Powehshell como Adm

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

Fechar e abrir outro Powershell como Adm

```powershell
choco install k3d -y
choco install jq -y
choco install yq -y
choco install kubernetes-helm -y

if ( -not ( Test-Path $Profile ) ) { New-Item -Path $Profile -Type File -Force }
k3d completion powershell | Out-File -Append $Profile

```

Depois:

```powershell
$env:KUBECONFIG=$("$env:USERPROFILE\.kube\config-devops-lab")
k3d cluster create devops-lab --agents 2 --port "80:80@loadbalancer"
. $profile

```

## Instalando o K3D no linux

```bash
sudo apt update 
sudo apt upgrade -y
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash


```


Dica:

Usar k9s ( https://k9scli.io/ )

Acrescentar no $profile
```bash
New-Alias k kubectl.exe
$env:KUBECONFIG=$(Get-ChildItem -Path "$env:USERPROFILE\.kube\config*").FullName -join ";"
```

Linux 

Install k3d

```bash
wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

Acrescentar no  ~/.zshrc

```bash
alias k="kubectl"
alias ns="kubens"
alias ctx="kubectx"

source <(kubectl completion zsh)
export KUBECONFIG=$(for i in $(find $HOME/.kube/ -iname 'config*') ; do echo -n ":$i"; done | cut -c 2-)
```
## Instalar k3d no windows

Instalar WSL

```
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```
Reiniciar

Instalar WSL2 Kernel update

https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi

Depois setar como padrão

```
wsl --set-default-version 2
```

Instalar o Rancher Desktop

https://rancherdesktop.io/

E usar o Docker como container runtime 

Instalando o K3D

Powehshell como Adm

```
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

Fechar e abrir outro Powershell como Adm

```

choco install k3d -y
choco install jq -y
choco install yq -y
choco install kubernetes-helm -y

if ( -not ( Test-Path $Profile ) ) { New-Item -Path $Profile -Type File -Force }
k3d completion powershell | Out-File -Append $Profile

```

Depois:

```

$env:KUBECONFIG=$("$env:USERPROFILE\.kube\config-devops-lab")
k3d cluster create devops-lab --agents 2 --port "6380:80@loadbalancer"
. $profile

```

Dica:

Usar k9s ( https://k9scli.io/ )

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
# Preparar Windows

## Instalar WSL

```powershell
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

Reiniciar

Instalar WSL2 Kernel update

https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi


Depois setar como padrão

```powershell
wsl --set-default-version 2
```

## Instalar Ubuntu

https://apps.microsoft.com/detail/9PN20MSR04DW?hl=pt-br&gl=BR


## Instalar o Rancher Desktop

https://rancherdesktop.io/

E usar o Docker como container runtime

## Instalar Windows Terminal

https://aka.ms/terminal




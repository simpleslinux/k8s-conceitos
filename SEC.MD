# Scanner de Imagens

## Trivy

```
trivy image fmnapoli/teste-app:3.0.0-inicial

```

## Grype

Site: https://github.com/anchore/grype

```bash

grype fmnapoli/teste-app:3.0.0-wolfi

```
Grype com template

```bash
grype fmnapoli/teste-app:3.0.0-wolfi -o template -t grype.tmpl > grype.html

firefox grype.html

```
# SBOM

## BOM

Site: https://github.com/kubernetes-sigs/bom

```bash
bom generate --image fmnapoli/teste-app:3.0.0-wolfi > teste-app.spdx
```
# Configuration GitHub Actions

## Go Modules

Les modules Go utilisent les tags GitHub comme versions. Aucun secret n'est requis.

## Déploiement

Le workflow se déclenche automatiquement lors de la création d'un tag `v*` (exemple : `v1.0.0`).
Les utilisateurs peuvent installer avec : `go get github.com/factpulse/sdk-go@v1.0.0`

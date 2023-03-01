# IFT630-SCAM

**S**imulateur de **C**onstruction **A**léatoire d'une **M**étropole

## Autheurs

- Kevin Bissonnette-Leblanc - 19 092 232
- Jonathan Bouthiette - 20 055 574
- Guillaume Cléroux - 20 067 819
- Olivier Cimon - 15 092 472

## Instructions Docker

C'est surement plus simple d'utiliser docker pour le développement,
alors j'ai créer un Dockerfile. Vous pouvez surement trouver
une façon de runner le tout en local si vous ne voulez pas utiliser docker.

```bash
# Build the docker image
docker compose build
```

```bash
# Run the image
docker compose run ift630-scam
```

## CI

Tous les tests doivent être ajoutés dans le dossier `test/` et suivre
la [convention Go](https://pkg.go.dev/testing).

On a un petit pipeline de CI qui va rouler sur toutes les branches à l'exception
de `main`.

Le pipeline s'exécute uniquement lors d'un `push`.

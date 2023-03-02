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

### Tests

Tous les tests doivent être ajoutés dans le dossier `test/` et suivre
la [convention Go](https://pkg.go.dev/testing).

De plus, toutes les PR doivent passer les tests avant de merge.

### Rapport

L'écriture du rapport est dans le dossier `docs/`. Afin de rendre la mise en page
plus facile, on peut directement écrire dans le fichier `rapport.md`.

Lors du CI, un runner va automatiquement mettre à jour la table des matières
et générer un `rapport.pdf` à jour. Vous avez seulement à écrire votre texte
et créer les entêtes au besoin pour ajouter une nouvelle entrée dans la table
des matières. L'édition des liens se fera de façon automatique.

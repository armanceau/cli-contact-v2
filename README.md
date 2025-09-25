# cli-contact-v2

Un petit CRM en ligne de commande développé en Go pour gérer des contacts (ID, Nom, Email).  
Permet d’ajouter, lister, supprimer et mettre à jour des contacts directement depuis le terminal.

## Fonctionnalités

- Afficher un menu interactif en boucle
- Ajouter un contact
- Lister tous les contacts
- Supprimer un contact par ID
- Mettre à jour un contact
- Quitter l’application

## Prérequis

- Go >= 1.25.1
- SQLite (si utilisation de GORM, ici nous allons passer outre via `github.com/glebarez/sqlite`, cette option est précisée dans la documentation de gorm : [ici](https://gorm.io/docs/connecting_to_the_database.html))

## Installation

1. Cloner le repository :

```bash
git clone https://github.com/armanceau/cli-contact-V2.git
cd cli-contact-v2
go mod tidy
```

2. Lancer le projet :

```bash
go run main.go
```

## Backend

Le stockage peut se faire via :

- Fichier JSON
- Base SQLite avec GORM (par défaut)

Pour changer le backend, modifiez `root.go` :

```go
store := storage.NewJsonStore("./contacts.json")

database.ConnectDB()
store := storage.NewGormStore()
```

## Structure du projet

```markdown
cli-contact-v2/
├── cmd/
│ └── crm/
├── internal/
│ ├── app/
│ ├── database/
│ ├── models/
│ └── storage/
├── go.mod
├── go.sum
├── main.go
└── .gitignore
```

## Utilisation

### Menu interactif

_Lancer le programme. Puis suivre les instructions pour :_

1. Ajouter un contact
2. Lister les contacts
3. Supprimer un contact
4. Mettre à jour un contact
5. Quitter

## Auteur

Arthur Manceau 🙉

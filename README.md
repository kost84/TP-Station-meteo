# TP-Station-meteo

---

## 1. Présentation du projet

L'application a pour objectif de permettre la recherche d'une commune française par son nom, de gérer les homonymes via leur code INSEE unique, puis d'afficher les prévisions météorologiques sur plusieurs jours associées à ses coordonnées géographiques.

### Fonctionnalités prévues :
- Recherche textuelle de communes avec validation des saisies (longueur comprise entre 2 et 60 caractères).
- Prise en charge des homonymes grâce à l'identification par le code INSEE.
- Récupération et conversion des prévisions météorologiques quotidiennes via Open-Meteo.
- Affichage des dates au format français, des températures et des libellés de conditions météo (codes WMO).
- Intégration d'une carte géographique et de représentations graphiques des températures sans recours à JavaScript.
- Gestion centralisée des erreurs techniques avec pages dédiées (400, 404, 500) et journalisation interne.

---

## 2. APIs externes exploitées

1. **API Découpage administratif (geo.api.gouv.fr)**
   - Recherche par nom : `https://geo.api.gouv.fr/communes?nom={nom}&fields=nom,code,codesPostaux,codeDepartement,centre,population&boost=population`
   - Recherche par code INSEE : `https://geo.api.gouv.fr/communes/{code}?fields=nom,code,codesPostaux,population,centre`
   - Données extraites : nom de la commune, département, codes postaux, population et coordonnées géographiques (longitude et latitude issues du champ `centre`).

2. **API Open-Meteo (open-meteo.com)**
   - Prévisions : `https://api.open-meteo.com/v1/forecast?latitude={lat}&longitude={lon}&daily=temperature_2m_max,temperature_2m_min,weather_code&timezone=Europe%2FParis`
   - Données extraites : températures journalières minimales et maximales, codes météorologiques WMO.

---

## 3. Architecture du projet

Le projet est structuré selon une architecture en couches strictes, conformément aux exigences :

```text
.
|-- api/            # Appels HTTP sortants vers les APIs externes
|-- models/         # Structures de données Go (Commune, Day, etc.) et erreurs métier
|-- services/       # Logique applicative, agrégation et traitement des données
|-- controllers/    # Réception des requêtes HTTP, validation et orchestration
|-- routes/         # Déclaration du routeur et configuration des routes
|-- helpers/        # Gestion des gabarits HTML (rendu par buffer) et pages d'erreur
|-- templates/      # Fichiers HTML (layout commun, accueil, liste, météo, erreurs)
|-- assets/         # Fichiers statiques CSS et images (servis sous /static/)
|-- main.go         # Point d'entrée de l'application
|-- go.mod          # Définition du module Go
`-- README.md       # Documentation du projet
```

---

## 4. Routes de l'application

| Méthode | Route | Description |
| :--- | :--- | :--- |
| `GET` | `/` ou `/{$}` | Page d'accueil contenant le formulaire de recherche |
| `GET` | `/ville?q={nom}` | Affichage de la liste des communes correspondant à la recherche |
| `GET` | `/meteo?code={insee}` | Affichage détaillé des prévisions météo pour la commune sélectionnée |
| `GET` | `/static/*` | Distribution des ressources statiques (CSS) |
| `*` | `/*` | Route de capture renvoyant une réponse 404 (Not Found) |

### Exemples d'URL de test :
- Accueil : `http://localhost:8080/`
- Recherche d'une ville sans ambiguïté : `http://localhost:8080/ville?q=Nantes`
- Recherche avec homonymes multiples : `http://localhost:8080/ville?q=Saint-Martin`
- Consultation météo par code INSEE : `http://localhost:8080/meteo?code=13001` (Aix-en-Provence)
- URL invalide (test 404) : `http://localhost:8080/page-inconnue`

---

## 5. Installation et lancement

### Prérequis :
- Go version 1.22 ou supérieure installé sur la machine.

### Commandes de démarrage :

1. Cloner ou ouvrir le dossier du projet à la racine.
2. Télécharger les éventuelles dépendances (uniquement si applicable) :
   ```bash
   go mod tidy
   ```
3. Lancer le serveur web :
   ```bash
   go run .
   ```
4. Ouvrir un navigateur web à l'adresse suivante :
   ```text
   http://localhost:8080
   ```
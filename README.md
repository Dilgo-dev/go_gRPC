# 🔮 Microservices gRPC Demo

> Un projet d'exploration de la communication inter-microservices avec gRPC

## 📋 À propos

Ce projet illustre la communication entre microservices en utilisant **gRPC** plutôt que les traditionnelles requêtes HTTP. Il met en pratique les avantages de gRPC pour les architectures distribuées modernes.

### 🎯 Objectifs pédagogiques

- 🔍 **Découvrir gRPC** : Alternative performante aux API REST
- 🏗️ **Architecture microservices** : Communication inter-services efficace
- ⚡ **Performance** : Sérialisation binaire avec Protocol Buffers
- 🔄 **Communication bidirectionnelle** : Au-delà des requêtes HTTP simples

## 🏛️ Architecture

```
┌─────────────────┐    gRPC     ┌─────────────────┐
│   Node.js API   │ ──────────► │   Go Service    │
│   (Port 3000)   │             │   (Port 3333)   │
│                 │             │ + gRPC (50051)  │
└─────────────────┘             └─────────────────┘
        │                               │
        │ HTTP REST                     │ SQLite
        ▼                               ▼
   🌐 Utilisateur                   💾 Base de données
```

### 🧩 Composants

- **🟢 Service Node.js** : API REST exposée aux clients
- **🔵 Service Go** : Service de logs avec gRPC et base de données
- **📡 gRPC** : Communication inter-services haute performance
- **🐳 Docker** : Conteneurisation et orchestration

## 🚀 Démarrage rapide

### 📋 Prérequis

- 🐳 Docker et Docker Compose
- 📁 Git

### 🏃‍♂️ Installation

```bash
# Clone du projet
git clone https://github.com/Dilgo-dev/go_gRPC
cd go_gRPC

# Copie des variables d'environnement
cp .env.sample .env

# Lancement avec Docker Compose
docker-compose up --build
```

### 🧪 Test de l'application

```bash
# Page d'accueil (génère un log automatiquement)
curl http://localhost:3000/

# Création d'un log manuel
curl -X POST http://localhost:3000/log \
  -H "Content-Type: application/json" \
  -d '{"message": "Mon premier log via gRPC!"}'

# Consultation de tous les logs
curl http://localhost:3000/logs
```

## 📡 API Endpoints

### 🟢 Node.js Service (Port 3000)

| Méthode | Endpoint | Description |
|---------|----------|-------------|
| `GET` | `/` | Page d'accueil (génère un log) |
| `POST` | `/log` | Création d'un log manuel |
| `GET` | `/logs` | Récupération de tous les logs |

### 🔵 Go Service (Port 3333)

| Méthode | Endpoint | Description |
|---------|----------|-------------|
| `GET` | `/api/logs` | API REST directe pour les logs |

### 📡 gRPC Service (Port 50051)

| Service | Méthode | Description |
|---------|---------|-------------|
| `LogService` | `CreateLog` | Création d'un log via gRPC |
| `LogService` | `GetLogs` | Récupération des logs via gRPC |

---

**🎯 Objectif atteint** : Découverte pratique de gRPC comme alternative performante aux API REST traditionnelles dans les architectures microservices !

---

*Développé avec ❤️ pour l'apprentissage des architectures distribuées modernes*
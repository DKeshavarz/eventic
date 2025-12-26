# Eventic
Eventic is a platform designed to make event management simple, social, and organized. It allows users to participate in events, create their own events, and even form organizations to manage events professionally.
Eventic will be available both as a website and a Telegram.

## Instaltion Guid
Choose from multiple installation methods depending on your environment and requirements.
### Native Installation
**System Requirements**
- Go Version: 1.24 or higher

**Comands**
```sh
# Step 1: Clone the Repository
git clone git@github.com:DKeshavarz/eventic.git
cd eventic

# step 2 : Configure Go Environment, use .env.sample
cp .env.example .env
nano .env

# step 3: Download dependencies
make dep
# Or 
go mod tidy

# step 4: Run program
make
# Or 
go run main.go
```
###  Docker Installation
**System Requirements**
- Docker 

**Comands**
```sh
# Step 1: Pull the latest image from GitHub Container Registry
docker pull ghcr.io/dkeshavarz/eventic:latest

# Step 2: Create environment file

# Step 3: Run the container
docker run --name eventic \
  --env-file .env \
  -p 8080:8080 \
  -d \
  --restart unless-stopped \
  ghcr.io/dkeshavarz/eventic:latest
```



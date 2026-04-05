# EduChain

A blockchain-based educational certification system for issuing, managing, and verifying student diplomas using Ethereum smart contracts, IPFS, and Kubernetes.

## Overview

EduChain enables educational institutions to:

- Register as verified diploma providers
- Enroll students and associate them with Ethereum wallets
- Issue diplomas as ERC-721 NFTs stored on-chain, with the diploma documents stored on IPFS
- Verify diploma authenticity through the blockchain

The system consists of two Go microservices backed by a private Ethereum node, an IPFS node, and a high-availability PostgreSQL cluster — all deployed on Kubernetes.

## Architecture

```
                        ┌───────────────────────────────────┐
                        │           Kubernetes Cluster      │
                        │                                   │
  Client ──────────────►│  Smart Contract API  (:8081)      │
                        │     ├── PostgreSQL (3 replicas)   │
                        │     ├── Ethereum / Geth (:8545)   │
                        │     └── IPFS API (:8080)          │
                        │           └── IPFS Node (:5001)   │
                        └───────────────────────────────────┘
```

| Component | Role |
|-----------|------|
| **Smart Contract API** | Main REST API — auth, student/diploma management, Ethereum integration |
| **IPFS API** | File upload/download wrapper around the IPFS node |
| **Geth Node** | Private Ethereum node (PoA dev mode) |
| **IPFS Node** | Kubo node for distributed diploma document storage |
| **PostgreSQL** | CloudNative-PG cluster — providers, students, diplomas metadata |

## Prerequisites

- [Go](https://go.dev/) 1.24+
- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [Helm](https://helm.sh/docs/intro/install/)

## Getting Started

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `IPFS_NODE_IP` | IPFS node hostname | `ipfs-node` / `localhost` |
| `IPFS_NODE_PORT` | IPFS node RPC port | `5001` |
| `IPFS_API_HOST` | IPFS API hostname | `ipfs-api` / `localhost` |
| `IPFS_API_PORT` | IPFS API port | `8080` |
| `DB_HOST` | PostgreSQL hostname | `postgres-rw` / `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | PostgreSQL username | `admin` |
| `DB_PASSWORD` | PostgreSQL password | `admin` |
| `DB_NAME` | Database name | `educhain` |
| `NET_URL` | Ethereum RPC endpoint | `http://geth-node:8545` |
| `JWT_SIGNING_KEY` | JWT secret key | `superSecretKey` |
| `MASTER_KEY` | Path to encrypted wallet key file | `/master-key.json` |
| `MASTER_KEY_PASSWD` | Password for wallet key file | `testPassword123` |

### Local Development

```bash
# Clone the repo
git clone https://github.com/Gabi-Limberea/educhain.git
cd educhain

# Source environment variables
source .envrc

# Start Smart Contract API (http://localhost:8081)
cd smart-contract-api
go mod download
make run

# In a separate terminal — Start IPFS API (http://localhost:8080)
cd ipfs-api
go mod download
make run
```

### Kubernetes Deployment

The deployment script sets up a full local cluster using `kind`:

```bash
cd infra-config
./deploy.sh
```

This will:
1. Create a Kind cluster named `educhain`
2. Install the CloudNative-PG operator
3. Deploy a 3-replica PostgreSQL cluster
4. Deploy the IPFS node, Geth node, and IPFS API
5. Extract the Ethereum master key to `key-file.json`

After the script completes, deploy the Smart Contract API (update the ConfigMap in `smart-contract-api.yaml` with the master key from `key-file.json` first):

```bash
kubectl apply -f smart-contract-api.yaml
```

**Access services locally via port-forward:**

```bash
kubectl port-forward svc/smart-contract-api 8081:8081
kubectl port-forward svc/ipfs-api 8080:8080
```

### Building Docker Images

```bash
# Smart Contract API
cd smart-contract-api
make docker-build VERSION=v1.0.9
make docker-push VERSION=v1.0.9

# IPFS API
cd ipfs-api
make docker-build VERSION=v1.0.1
make docker-push VERSION=v1.0.1
```

## API Reference

### Smart Contract API

#### Authentication

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| `POST` | `/providers` | Register a new educational provider | None |
| `POST` | `/providers/token` | Issue a JWT token | Basic Auth |
| `GET` | `/account` | Get the authenticated provider's info | Bearer JWT |
| `GET` | `/healthz` | Health check | None |

#### Students

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| `POST` | `/students` | Batch enroll students | Bearer JWT |
| `GET` | `/students` | List all enrolled students | Bearer JWT |
| `GET` | `/students/{student_id}` | Get a specific student | Bearer JWT |

#### Diplomas

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| `POST` | `/students/{student_id}/diplomas` | Issue a diploma to a student | Bearer JWT |
| `GET` | `/students/{student_id}/diplomas` | Get all diplomas for a student | Bearer JWT |
| `POST` | `/diplomas` | Bulk issue diplomas | Bearer JWT |

Full OpenAPI 3.0 specification with interactive HTML docs available in the
`docs` directory:

- [smart-contract-api/docs/spec.yaml](smart-contract-api/docs/spec.yaml)
- [smart-contract-api/docs/redoc-static.html](smart-contract-api/docs/redoc-static.html)

### IPFS API

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/file` | Upload a file or ZIP archive to IPFS |
| `GET` | `/api/file/{cid}` | Download a file by its IPFS CID |
| `GET` | `/healthz` | Health check |

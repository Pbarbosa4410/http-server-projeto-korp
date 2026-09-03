# HTTP Server Projeto Korp

Projeto de laboratório DevOps utilizando Go, Docker, Nginx, Prometheus, Grafana e Ansible.

## Arquitetura

Fluxo principal:

Cliente
→ Nginx
→ Aplicação Go
→ Prometheus
→ Grafana

## Tecnologias

- Go 1.22
- Docker
- Docker Compose
- Nginx
- Prometheus
- Grafana
- Ansible

## Estrutura do projeto

```text
http-server-projeto-korp/
├── ansible/
│   └── playbook.yml
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── server/
│       └── server.go
├── monitoring/
│   ├── grafana/
│   │   └── provisioning/
│   │       ├── dashboards/
│   │       └── datasources/
│   │           └── datasource.yml
│   └── prometheus/
│       └── prometheus.yml
├── nginx/
│   └── conf.d/
│       └── default.conf
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
# Rate Limiter em Go

## Descrição

Este projeto implementa um **Rate Limiter** em Go para controlar o número de requisições feitas a um serviço web com base no endereço IP ou em um token de acesso. Ele utiliza o **Redis** como mecanismo de persistência e está configurado para ser executado com **Docker Compose**.

---

## Funcionalidades

- Limitação por endereço IP e/ou token de acesso.
- Configuração flexível via variáveis de ambiente.
- Mensagem apropriada e código HTTP `429` ao exceder o limite de requisições.
- Suporte para bloqueio temporário de IPs ou tokens após exceder o limite.
- Middleware para fácil integração ao servidor web.
- Implementação modular para troca de estratégias de persistência (ex.: Redis).
- Testes automatizados para validação de comportamento.

---

## Requisitos

### Tecnologias utilizadas

- Go (versão 1.20 ou superior)
- Redis
- Docker e Docker Compose

---

## Configuração do Ambiente

1. **Instale as dependências do projeto:**
   ```bash
   go mod tidy
2. **Crie um arquivo .env na raiz do projeto: Exemplo de configurações:**
    REDIS_HOST=redis
    REDIS_PORT=6379
    RATE_LIMIT_IP=10            # Limite de requisições por segundo por IP
    RATE_LIMIT_TOKEN=100        # Limite de requisições por segundo por Token
    BLOCK_DURATION=300          # Duração do bloqueio em segundos
3. **Suba os serviços com Docker Compose:**
    docker-compose up --build

### Usabilidade
1. **Testar o servidor web**
    O servidor estará disponível na porta 8080.

    Para enviar uma requisição:

    curl http://localhost:8080
    Para testar o limite por token:

    curl -H "API_KEY: <TOKEN>" http://localhost:8080
2. **Respostas esperadas**
    Dentro do limite:

    {
    "message": "Bem-vindo ao Rate Limited API!"
    }
    Quando o limite for excedido (HTTP 429):

    {
    "message": "you have reached the maximum number of requests or actions allowed within a certain time frame"
}
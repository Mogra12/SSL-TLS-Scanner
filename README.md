# SSL/TLS Scanner

## Descrição

Este é um scanner simples para verificar a configuração SSL/TLS de um servidor. O programa realiza uma conexão segura com o host especificado e exibe informações sobre o certificado SSL e os detalhes da conexão TLS.

## Funcionalidades

- Exibe informações sobre o certificado SSL, incluindo:
  - Nome comum (CN) do emissor e do sujeito
  - Validade do certificado (início e fim)
  - Algoritmo de assinatura
- Exibe detalhes da conexão TLS, como:
  - Versão do protocolo TLS utilizado
  - Suíte de cifra usada na conexão
- Suporte a conexões com verificação SSL desativada (*InsecureSkipVerify*)
- Mensagens coloridas para melhor visualização dos erros

## Requisitos

- Go 1.18+
- Biblioteca `github.com/fatih/color` para formatação colorida no terminal

## Instalação

Clone o repositório e instale as dependências necessárias:

```sh
# Clonar o repositório
github.com/Mogra12/SSL-TLS-Scanner.git
cd SSL-TLS-Scanner

# Instalar dependências
go mod tidy
```

## Uso

Execute o scanner passando o hostname e a porta do servidor como argumento:

```sh
go run main.go -h example.com:443
```

### Exemplo de Saída

```
=== SSL Certificate ===
Issued to: example.com
Issued by: Let's Encrypt Authority X3
Valid from: 2024-01-01 12:00:00 +0000 UTC
Valid until: 2024-03-31 12:00:00 +0000 UTC
Signature algorithm: SHA256-RSA

=== TLS Configuration ===
TLS Version: TLS 1.2
Cipher Suite: TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
```

## Licença

Este projeto está licenciado sob a MIT License


# Generate PDF - Documentação

## 📌 Sobre o Projeto

Este simples projeto permite a geração de PDFs utilizando a biblioteca [Maroto](https://github.com/johnfercher/maroto). Ele cria um documento PDF contendo uma lista de instrumentos musicais para um projeto comunitário.

## 🚀 Tecnologias Utilizadas

- **Go** (Golang)
- **Maroto** (Biblioteca para geração de PDFs)

## 📂 Estrutura do Projeto

```
/generate-pdf
│-- /data
│   ├── data.go  # Contém a lista de instrumentos musicais
│-- /images
│   ├── logo_div_generate.png  # Logotipo utilizado no cabeçalho do PDF
│-- /pdfs
│   ├── div_generate.pdf  # PDF gerado
│-- main.go  # Código principal
│-- go.mod  # Gerenciamento de dependências
│-- go.sum  # Checksum das dependências
```

## 📥 Instalação

### **Pré-requisitos**

- Ter o **Go** instalado ([Baixar Go](https://go.dev/dl/))
- Ter o **Git** instalado ([Baixar Git](https://git-scm.com/))

### **Passo 1: Clonar o repositório**

```bash
git clone https://github.com/EmmanoelDan/generate-pdf.git
cd generate-pdf
```

### **Passo 2: Instalar as dependências**

```bash
go mod tidy
```

Isso garantirá que todas as dependências necessárias sejam baixadas.

## ▶️ Como Executar o Projeto

Após a instalação das dependências, execute o seguinte comando:

```bash
go run main.go
```

Se tudo estiver correto, será gerado um PDF na pasta `pdfs/` com o nome `div_generate.pdf`.

## 🛠️ Testando o Projeto

1. Execute o comando `go run main.go`
2. Verifique a pasta `pdfs/` e abra o arquivo `div_generate.pdf`
3. O PDF deve conter a lista de instrumentos musicais formatada corretamente

## 📝 Personalizações

- Para alterar a lista de instrumentos, edite o arquivo `data/data.go`.
- Para mudar o logotipo, substitua o arquivo `images/logo_div_generate.png`.

## 📄 Licença

Este projeto é de código aberto e está licenciado sob a **MIT License**.

---

Criado por **Emmanoel Dan** 🚀

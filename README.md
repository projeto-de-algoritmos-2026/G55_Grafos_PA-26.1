# 3D-Print-Scheduler

Número da Lista: 1<br>
Conteúdo da Disciplina: Grafos<br>

## Alunos

| Matrícula | Aluno |
| -- | -- |
| 19/0102977 | Artur Ricardo dos Santos Lopes |

## Sobre

O **3D Print Scheduler** é um agendador de ordem de impressão 3D baseado em teoria dos grafos. Em impressoras FDM, certas peças ou partes de uma montagem dependem de outras já estarem prontas na mesa — seja por questões de suporte físico ou encaixe sequencial. O projeto modela esse problema como um **Grafo Acíclico Dirigido (DAG)**, onde cada nó representa uma peça a ser impressa e cada aresta dirigida `A → B` indica que a peça `A` deve ser impressa antes da peça `B`.

O algoritmo central é a **Ordenação Topológica** implementada via busca em profundidade (DFS), que determina uma sequência de impressão válida respeitando todas as dependências. O ponto alto do projeto é a **detecção de ciclos**: caso o usuário defina uma dependência circular (ex: peça A depende de B, que depende de A), o sistema identifica o ciclo e retorna um erro descritivo, prevenindo uma ordem de impressão impossível.

## Screenshots

> 🖼️ *Screenshot 1 — Grafo de dependências construído e exibido no terminal*

![Screenshot 1](docs/screenshots/screenshot1.png)

> 🖼️ *Screenshot 2 — Ordenação topológica executada com sucesso*

![Screenshot 2](docs/screenshots/screenshot2.png)

> 🖼️ *Screenshot 3 — Detecção de ciclo com mensagem de erro*

![Screenshot 3](docs/screenshots/screenshot3.png)

## Instalação

Linguagem: Go (Golang) 1.22+<br>
Framework: Nenhum (biblioteca padrão apenas)<br>

**Pré-requisitos:**

- Ter o [Go](https://go.dev/dl/) instalado (versão 1.22 ou superior)
- Verificar instalação com `go version`

**Passos:**

```bash
# Clone o repositório
git clone https://github.com/projeto-de-algoritmos-2026/G55_Grafos_PA-26.1.git

# Entre na pasta do projeto
cd G55_Grafos_PA-26.1/3d-print-scheduler

# Baixe as dependências (caso existam)
go mod tidy

# Compile e execute
go run main.go
```

## Uso

Ao executar o projeto, o programa roda o cenário de exemplo definido em `main.go`, exibindo:

1. O grafo de dependências montado (lista de adjacência)
2. A ordem de impressão sugerida pela ordenação topológica
3. Caso exista um ciclo nas dependências, uma mensagem de erro indicando as peças envolvidas

Para testar com suas próprias peças, edite o slice `pecas` e o slice `dependencias` diretamente em `main.go`:

```go
// Adicione suas peças
pecas := []*scheduler.PrintPart{
    {ID: "minha_base", Name: "Base Customizada", EstimatedHours: 2.5},
    // ...
}

// Defina as dependências
dependencias := [][2]string{
    {"minha_base", "outra_peca"},
    // ...
}
```

## Outros

- O projeto foi desenvolvido como trabalho prático da disciplina **Projeto de Algoritmos (2026.1)** na Universidade de Brasília (UnB).
- A estrutura do grafo utiliza **lista de adjacência** (`map[string][]string`) para eficiência em grafos esparsos, que é o caso típico em montagens de impressão 3D.
- A detecção de ciclos é feita com coloração de nós (branco / cinza / preto), abordagem clássica descrita em *Introduction to Algorithms* (CLRS).
- Trabalho futuro inclui leitura de grafos via arquivo YAML e uma interface visual do DAG no terminal.
// Package scheduler implementa o núcleo do agendador de impressão 3D.
// A ideia central é modelar as peças como nós de um grafo dirigido,
// onde uma aresta A → B significa "A deve ser impresso antes de B".
package scheduler

import "fmt"

// PrintPart representa uma peça que será impressa.
// Em FDM, uma peça pode depender de outras já estarem na mesa
// (ex: suportes, encaixes, bases).
type PrintPart struct {
	ID          string  // Identificador único, ex: "base_frame"
	Name        string  // Nome legível para exibição
	EstimatedHours float64 // Tempo estimado de impressão em horas

	// TODO: adicionar campos como filamento, temperatura, escala
}

// DependencyGraph representa o grafo de dependências entre peças.
// Usamos lista de adjacência: para cada peça, guardamos quais peças
// ela "desbloqueia" (seus dependentes diretos).
//
// Exemplo visual:
//   base_frame → side_panel → top_cover
//              ↘ front_door
//
// Isso significa: imprima base_frame antes de side_panel e front_door.
type DependencyGraph struct {
	// parts é o registro central de todas as peças, indexado pelo ID.
	parts map[string]*PrintPart

	// adjacency é a lista de adjacência do grafo.
	// adjacency["A"] = ["B", "C"] significa que B e C dependem de A.
	adjacency map[string][]string

	// TODO: carregar grafo de um arquivo YAML/JSON para facilitar testes
	// TODO: adicionar suporte a pesos nas arestas (ex: tempo de espera entre impressões)
}

// NewDependencyGraph cria e retorna um grafo vazio, pronto para uso.
func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		parts:     make(map[string]*PrintPart),
		adjacency: make(map[string][]string),
	}
}

// AddPart registra uma nova peça no grafo.
// Retorna erro se uma peça com o mesmo ID já existir (evita duplicatas silenciosas).
func (g *DependencyGraph) AddPart(part *PrintPart) error {
	if _, exists := g.parts[part.ID]; exists {
		return fmt.Errorf("peça com ID '%s' já existe no grafo", part.ID)
	}

	g.parts[part.ID] = part
	// Garante que a peça tenha entrada na lista de adjacência,
	// mesmo que ela não tenha dependentes (nó folha).
	g.adjacency[part.ID] = []string{}

	return nil
}

// AddDependency declara que 'afterID' só pode ser impresso depois de 'beforeID'.
// Cria a aresta directed: beforeID → afterID.
//
// Retorna erro se algum dos IDs não estiver registrado no grafo.
func (g *DependencyGraph) AddDependency(beforeID, afterID string) error {
	if _, exists := g.parts[beforeID]; !exists {
		return fmt.Errorf("peça '%s' não encontrada — adicione-a antes de criar dependências", beforeID)
	}
	if _, exists := g.parts[afterID]; !exists {
		return fmt.Errorf("peça '%s' não encontrada — adicione-a antes de criar dependências", afterID)
	}

	// Adiciona a aresta na lista de adjacência de 'beforeID'
	g.adjacency[beforeID] = append(g.adjacency[beforeID], afterID)

	return nil
}

// Parts retorna todas as peças registradas.
// Útil para iteração externa sem expor o map interno diretamente.
func (g *DependencyGraph) Parts() map[string]*PrintPart {
	return g.parts
}

// Neighbors retorna as peças que dependem diretamente de 'partID'.
// Em termos de grafo: retorna os vizinhos na lista de adjacência.
func (g *DependencyGraph) Neighbors(partID string) []string {
	return g.adjacency[partID]
}
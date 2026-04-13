// Este arquivo implementa o algoritmo central do projeto:
// Ordenação Topológica via DFS com detecção de ciclos.
//
// A ideia de coloração de nós vem do livro CLRS (Introduction to Algorithms):
//   - branco (0): nó ainda não visitado
//   - cinza  (1): nó sendo visitado (está na pilha de recursão atual)
//   - preto  (2): nó completamente processado
//
// Se durante a DFS encontrarmos uma aresta que aponta para um nó CINZA,
// significa que encontramos um ciclo — impossível ordenar topologicamente.
package scheduler

import "fmt"

// cores dos nós durante a DFS
const (
	branco = 0 // não visitado
	cinza  = 1 // em processamento (na pilha de recursão)
	preto  = 2 // finalizado
)

// TopoResult é o resultado da ordenação topológica.
// Se houver ciclo, Err estará preenchido e Order estará vazio.
type TopoResult struct {
	Order []string // IDs das peças na ordem de impressão
	Err   error    // erro de ciclo detectado, se houver
}

// TopologicalSort executa a ordenação topológica no grafo.
// Retorna a ordem de impressão das peças ou um erro se houver ciclo.
func (g *DependencyGraph) TopologicalSort() TopoResult {
	cor := make(map[string]int)
	pilha := []string{} // acumula os nós em ordem de finalização

	// Inicializa todos os nós como brancos (não visitados)
	for id := range g.parts {
		cor[id] = branco
	}

	// Dispara a DFS a partir de cada nó ainda não visitado
	for id := range g.parts {
		if cor[id] == branco {
			if err := g.dfs(id, cor, &pilha); err != nil {
				return TopoResult{Err: err}
			}
		}
	}

	// A pilha está em ordem reversa de finalização — invertemos para
	// obter a ordem topológica correta (quem termina por último, imprime primeiro)
	order := make([]string, len(pilha))
	for i, id := range pilha {
		order[len(pilha)-1-i] = id
	}

	return TopoResult{Order: order}
}

// dfs é a busca em profundidade recursiva.
// Retorna erro imediatamente ao detectar um ciclo.
func (g *DependencyGraph) dfs(id string, cor map[string]int, pilha *[]string) error {
	cor[id] = cinza // marca como "em visita"

	for _, vizinho := range g.adjacency[id] {
		if cor[vizinho] == cinza {
			// Aresta para um nó cinza = ciclo detectado!
			return fmt.Errorf(
				"ciclo detectado: '%s' → '%s' cria uma dependência circular impossível de imprimir",
				id, vizinho,
			)
		}
		if cor[vizinho] == branco {
			if err := g.dfs(vizinho, cor, pilha); err != nil {
				return err // propaga o erro ciclo acima na recursão
			}
		}
	}

	cor[id] = preto          // marca como completamente processado
	*pilha = append(*pilha, id) // empilha ao finalizar
	return nil
}
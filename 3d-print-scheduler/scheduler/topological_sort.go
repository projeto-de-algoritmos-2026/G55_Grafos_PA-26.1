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

import (
	"fmt"
	"strings"
)

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
	pilha := []string{}
	caminho := []string{} // rastrea o caminho atual da DFS para montar a mensagem de ciclo

	// Inicializa todos os nós como brancos (não visitados)
	for id := range g.parts {
		cor[id] = branco
	}

	// Dispara a DFS a partir de cada nó ainda não visitado
	for id := range g.parts {
		if cor[id] == branco {
			if err := g.dfs(id, cor, &pilha, &caminho); err != nil {
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
// Mantém o slice 'caminho' para reconstruir o ciclo caso ele seja detectado.
func (g *DependencyGraph) dfs(id string, cor map[string]int, pilha *[]string, caminho *[]string) error {
	cor[id] = cinza
	*caminho = append(*caminho, id) // entra no caminho atual

	for _, vizinho := range g.adjacency[id] {
		if cor[vizinho] == cinza {
			// Encontrou ciclo — reconstrói o caminho a partir do nó que fecha o loop
			inicioCiclo := vizinho
			indiceCiclo := -1
			for i, no := range *caminho {
				if no == inicioCiclo {
					indiceCiclo = i
					break
				}
			}

			// Monta a string "A → B → C → A" mostrando o loop completo
			trecho := append((*caminho)[indiceCiclo:], inicioCiclo)
			return fmt.Errorf(
				"ciclo detectado: %s — essas peças formam uma dependência circular impossível de imprimir",
				strings.Join(trecho, " → "),
			)
		}
		if cor[vizinho] == branco {
			if err := g.dfs(vizinho, cor, pilha, caminho); err != nil {
				return err
			}
		}
	}

	cor[id] = preto
	*pilha = append(*pilha, id)
	*caminho = (*caminho)[:len(*caminho)-1] // sai do caminho atual (backtrack)
	return nil
}
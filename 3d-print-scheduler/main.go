package main

import (
	"fmt"
	"log"

	"github.com/projeto-de-algoritmos-2026/G55_Grafos_PA-26.1/scheduler"
)

func main() {
	fmt.Println("=== 3D Print Scheduler ===")
	fmt.Println("Algoritmo: Ordenação Topológica + Detecção de Ciclos")
	fmt.Println()

	// Montando um exemplo real: peças de um robô simples.
	// A ordem de impressão importa: a base tem que existir antes das laterais.
	g := scheduler.NewDependencyGraph()

	pecas := []*scheduler.PrintPart{
		{ID: "base",        Name: "Base do Robô",      EstimatedHours: 3.5},
		{ID: "lateral_esq", Name: "Lateral Esquerda",  EstimatedHours: 2.0},
		{ID: "lateral_dir", Name: "Lateral Direita",   EstimatedHours: 2.0},
		{ID: "topo",        Name: "Tampa Superior",     EstimatedHours: 1.5},
		{ID: "cabeca",      Name: "Cabeça",             EstimatedHours: 4.0},
	}

	for _, p := range pecas {
		if err := g.AddPart(p); err != nil {
			log.Fatalf("erro ao adicionar peça: %v", err)
		}
	}

	// Definindo a ordem de precedência:
	// base → laterais → topo → cabeça
	dependencias := [][2]string{
		{"base", "lateral_esq"},
		{"base", "lateral_dir"},
		{"lateral_esq", "topo"},
		{"lateral_dir", "topo"},
		{"topo", "cabeca"},
	}

	for _, dep := range dependencias {
		if err := g.AddDependency(dep[0], dep[1]); err != nil {
			log.Fatalf("erro ao adicionar dependência: %v", err)
		}
	}

	// Exibe o grafo construído
	fmt.Println("Grafo de dependências construído:")
	for id, part := range g.Parts() {
		vizinhos := g.Neighbors(id)
		fmt.Printf("  [%s] %s (%.1fh) → %v\n", id, part.Name, part.EstimatedHours, vizinhos)
	}

	// Executa a ordenação topológica
	resultado := g.TopologicalSort()
	if resultado.Err != nil {
		fmt.Printf("\n❌ erro: %v\n", resultado.Err)
		return
	}

	fmt.Println("\nOrdem de impressão sugerida:")
	for i, id := range resultado.Order {
		part := g.Parts()[id]
		fmt.Printf("  %d. [%s] %s (%.1fh)\n", i+1, id, part.Name, part.EstimatedHours)
	}

	fmt.Println()
}
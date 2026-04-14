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

	fmt.Println("--- Cenário 1: grafo válido (DAG) ---")
	rodarCenarioValido()

	fmt.Println()
	fmt.Println("--- Cenário 2: grafo com ciclo ---")
	rodarCenarioComCiclo()
}

// rodarCenarioValido demonstra a ordenação topológica num grafo sem ciclos.
func rodarCenarioValido() {
	g := scheduler.NewDependencyGraph()

	pecas := []*scheduler.PrintPart{
		{ID: "base",        Name: "Base do Robô",     EstimatedHours: 3.5},
		{ID: "lateral_esq", Name: "Lateral Esquerda", EstimatedHours: 2.0},
		{ID: "lateral_dir", Name: "Lateral Direita",  EstimatedHours: 2.0},
		{ID: "topo",        Name: "Tampa Superior",   EstimatedHours: 1.5},
		{ID: "cabeca",      Name: "Cabeça",           EstimatedHours: 4.0},
	}

	for _, p := range pecas {
		if err := g.AddPart(p); err != nil {
			log.Fatalf("erro ao adicionar peça: %v", err)
		}
	}

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

	resultado := g.TopologicalSort()
	if resultado.Err != nil {
		fmt.Printf("❌ erro inesperado: %v\n", resultado.Err)
		return
	}

	fmt.Println("Ordem de impressão sugerida:")
	for i, id := range resultado.Order {
		part := g.Parts()[id]
		fmt.Printf("  %d. [%s] %s (%.1fh)\n", i+1, id, part.Name, part.EstimatedHours)
	}

	// Calcula e exibe as métricas de tempo da ordem sugerida
	metrics, err := g.CalculateMetrics(resultado.Order)
	if err != nil {
		fmt.Printf("erro ao calcular métricas: %v\n", err)
		return
	}

	fmt.Printf("\nTempo total estimado: %.1fh\n", metrics.TotalHours)
	fmt.Println("Tempo acumulado por etapa:")
	for i, id := range resultado.Order {
		part := g.Parts()[id]
		fmt.Printf("  %d. [%s] %s → acumulado: %.1fh\n",
			i+1, id, part.Name, metrics.CumulativeMap[id])
	}
}

// rodarCenarioComCiclo demonstra a detecção de dependência circular.
// cabeça → base cria um ciclo: base → ... → cabeça → base
func rodarCenarioComCiclo() {
	g := scheduler.NewDependencyGraph()

	pecas := []*scheduler.PrintPart{
		{ID: "base",   Name: "Base do Robô", EstimatedHours: 3.5},
		{ID: "topo",   Name: "Tampa",        EstimatedHours: 1.5},
		{ID: "cabeca", Name: "Cabeça",       EstimatedHours: 4.0},
	}

	for _, p := range pecas {
		if err := g.AddPart(p); err != nil {
			log.Fatalf("erro ao adicionar peça: %v", err)
		}
	}

	// Grafo: base → topo → cabeça → base (ciclo!)
	dependencias := [][2]string{
		{"base", "topo"},
		{"topo", "cabeca"},
		{"cabeca", "base"}, // esta aresta cria o ciclo
	}

	for _, dep := range dependencias {
		if err := g.AddDependency(dep[0], dep[1]); err != nil {
			log.Fatalf("erro ao adicionar dependência: %v", err)
		}
	}

	resultado := g.TopologicalSort()
	if resultado.Err != nil {
		fmt.Printf("❌ %v\n", resultado.Err)
		return
	}

	fmt.Println("ordenação concluída (não deveria chegar aqui)")
}
// Package scheduler — este arquivo calcula métricas de tempo sobre
// a ordem de impressão gerada pela ordenação topológica.
package scheduler

import "fmt"

// PrintMetrics agrupa as métricas calculadas sobre uma ordem de impressão.
type PrintMetrics struct {
	TotalHours    float64            // tempo total somado de todas as peças
	CumulativeMap map[string]float64 // tempo acumulado até cada peça (inclusive)
}

// CalculateMetrics recebe a ordem topológica e calcula o tempo total
// e o tempo acumulado etapa a etapa.
//
// Exemplo: se a ordem é [base(3.5h), topo(1.5h), cabeça(4.0h)],
// o mapa acumulado será: base=3.5, topo=5.0, cabeça=9.0
func (g *DependencyGraph) CalculateMetrics(order []string) (PrintMetrics, error) {
	metrics := PrintMetrics{
		CumulativeMap: make(map[string]float64),
	}

	acumulado := 0.0
	for _, id := range order {
		part, exists := g.parts[id]
		if !exists {
			return PrintMetrics{}, fmt.Errorf("peça '%s' da ordem não encontrada no grafo", id)
		}

		acumulado += part.EstimatedHours
		metrics.CumulativeMap[id] = acumulado
	}

	metrics.TotalHours = acumulado
	return metrics, nil
}
package k8sconfigmap

import "k8s.io/api/core/v1"

func toCells(std []v1.ConfigMap) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = configMapCell(std[i])
	}
	return cells
}

func fromCells(cells []DataCell) []v1.ConfigMap {
	configMaps := make([]v1.ConfigMap, len(cells))
	for i := range cells {
		configMaps[i] = v1.ConfigMap(cells[i].(configMapCell))
	}

	return configMaps
}

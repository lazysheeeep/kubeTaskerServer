package k8sservice

import "k8s.io/api/core/v1"

func toCells(std []v1.Service) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = serviceCell(std[i])
	}
	return cells
}

func fromCells(cells []DataCell) []v1.Service {
	services := make([]v1.Service, len(cells))
	for i := range cells {
		services[i] = v1.Service(cells[i].(serviceCell))
	}

	return services
}

package k8sdeployment

import "k8s.io/api/apps/v1"

func toCells(std []v1.Deployment) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = deploymentCell(std[i])
	}
	return cells
}

func fromCells(cells []DataCell) []v1.Deployment {
	deployments := make([]v1.Deployment, len(cells))
	for i := range cells {
		deployments[i] = v1.Deployment(cells[i].(deploymentCell))
	}
	return deployments
}

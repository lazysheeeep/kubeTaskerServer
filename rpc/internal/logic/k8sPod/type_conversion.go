package k8sPod

import "k8s.io/api/core/v1"

// 类型转换方法corev1.Pod --> DataCell,DataCell-->corev1.Pod
func toCells(pods []v1.Pod) []DataCell {
	cells := make([]DataCell, len(pods))
	for i := range pods {
		cells[i] = podCell(pods[i])
	}
	return cells
}

func fromCells(cells []DataCell) []v1.Pod {
	pods := make([]v1.Pod, len(cells))
	for i := range cells {
		// cells[i].(podCell)是将DataCell类型转换成podCell
		pods[i] = v1.Pod(cells[i].(podCell))
	}
	return pods
}

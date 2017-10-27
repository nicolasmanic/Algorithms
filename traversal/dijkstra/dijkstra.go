package dijkstra

import (
	"personal/Algorithms/traversal"
)

var MAX_WEIGHT = 100000

func GetShortestPaths(graph *traversal.WeighedGraph, startingKey int ) map[int]int {
	numberOfNodes := len(graph.Nodes)
	dist := make(map[int]int)
	sptSet := make(map[int]bool)

	//initialization of arrays
	for _, node := range graph.Nodes {
		dist[node.Key] = MAX_WEIGHT
		sptSet[node.Key] = false
	}

	//The distance from startingPoint to startingPoint is 0
	dist[startingKey] = 0

	for i:=0;i<=numberOfNodes -1 ; i++ {
		minNodeKey := returnMinimumNonFinaleKey(dist, sptSet)
		if minNodeKey<0{
			panic("Min Key less than 0")
		}
		minNode := graph.GetNode(minNodeKey)
		sptSet[minNode.Key] = true
		neighbors,_:=graph.GetNeighbors(minNode.Key)

		for neighbor,_ := range neighbors{
			if dist[minNode.Key] + minNode.Neighbors[neighbor] < dist[neighbor.Key] && !sptSet[neighbor.Key]{
				dist[neighbor.Key] = dist[minNode.Key] + minNode.Neighbors[neighbor]
			}
		}
	}
	return  dist

}
func returnMinimumNonFinaleKey(distances map[int]int, sptSet map[int]bool) int {

	minWeight:= MAX_WEIGHT
	minKey := -1
	for key, weight := range distances{
		if weight<minWeight && !sptSet[key]{
			minKey = key
			minWeight = weight
		}
	}
	return minKey
}
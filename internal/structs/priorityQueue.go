package structs

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
   return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
   pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
   node := x.(*Point)
   *pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
   old := *pq
   n := len(old)
   node := old[n-1]
   *pq = old[0 : n-1]
   return node
}
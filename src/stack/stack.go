package stack
import(
	"container/heap"
	"fmt"
)

type Item struct{
	Value string
	Priority int
	Index int
}

type PriorityQueue [] *Item

func (pq PriorityQueue) Len() int{
	return len(pq)
}

func (pq PriorityQueue) Less(i,j int) bool{
	return pq[i].Priority>pq[j].Priority
}

func (pq PriorityQueue) Swap(i,j int){
	pq[i],pq[j]=pq[j],pq[i]
	pq[i].Index=i
	pq[j].Index=j
}

func(pq *PriorityQueue) Push(x interface{}){
	n:=len(*pq)
	item:=x.(*Item)
	item.Index=n
	*pq=append(*pq,item)
	
}

func (pq *PriorityQueue) Pop() interface{}{
	old:=*pq
	n:=len(old)
	item:=old[n-1]
	item.Index=-1
	*pq=old[0:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item,value string,priority int){
	item.Value=value
	item.Priority=priority
	heap.Fix(pq,item.Index)
}

func Test(){
	items:=map[string]int{"banana":3,"apple":2,"pear":4}
	pq:=make(PriorityQueue,len(items))
	i:=0
	for value,priority:=range items{
		pq[i]=&Item{
			Value: value,
			Priority:priority,
			Index: i,
		}
		i++
	}
	heap.Init(&pq)
	
	//Insert a new item and then modify its priority
	item:=&Item{
		Value:"orange",
		Priority:1,
	}
	heap.Push(&pq,item)
	pq.update(item,item.Value,5)
	
	//Take the items out
	for pq.Len()>0{
		item:=heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s",item.Priority,item.Value)
	}
}

package main
import "fmt"
var graph = make(map[string]map[string]bool)
func addEdge(from, to string) {
    edges := graph[from]
    if edges == nil {
        edges = make(map[string]bool)
        graph[from] = edges
    }
    edges[to] = true
}

func hasEdge(from, to string) bool {
    return graph[from][to]
}

func main() {
    //create a directed graph
    addEdge("A","B");
    addEdge("A","C");
    addEdge("B","C");

    var node = [3]string{"A","B","C"}
    for _,s_node := range node {
       for _, e_node := range node {
            if s_node == e_node {
                continue
            }
            if(hasEdge(s_node,e_node)) {
                fmt.Printf("%s->%s exists.\n",s_node, e_node)
            }
       }
    }
}

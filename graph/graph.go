package graph

type (
	Node struct {
		Name         string
		children     []*Node
		predecessors []*Node
	}
)

// AddChild links two nodes by updating the children and predecessors
// properties.
func (n *Node) AddChild(c *Node) *Node {
	n.children = append(n.children, c)
	c.predecessors = append(c.predecessors, n)
	return n
}

// Ordered returns all distinct Node names. The order is preferably such that
// all managers of a report have a lower list-index than the report itself.
func (n *Node) Ordered() []string {
	// TODO implement
	return nil
}

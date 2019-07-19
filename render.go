package gas

// RenderCore render station
type RenderCore struct {
	BE BackEnd

	queue []*RenderNode
}

// RenderNode node storing changes
type RenderNode struct {
	index int // The index of the item in the heap.

	Type RenderType

	New, Old                     interface{} // *Component, string, int, etc
	NodeParent, NodeNew, NodeOld interface{} // *dom.Element, etc

	Data map[string]interface{} // using only for Type == DataType
}

// RenderType renderNode type
type RenderType int

const (
	// ReplaceType type for replace node
	ReplaceType RenderType = iota

	// CreateType type for create nodes
	CreateType

	// DeleteType type for delete node
	DeleteType

	// RecreateType type for ReCreate
	RecreateType
)

// Add push render nodes to render queue and trying to execute all queue
func (rc *RenderCore) Add(node *RenderNode) {
	rc.queue = append(rc.queue, node)
}

// GetAll return render nodes from queue
func (rc *RenderCore) GetAll() []*RenderNode {
	return rc.queue
}

// Exec run all render nodes in render core
func (rc *RenderCore) Exec() error {
	for _, node := range rc.queue {
		err := rc.BE.ExecNode(node)
		if err != nil {
			return err
		}
	}

	rc.queue = []*RenderNode{}

	return nil
}

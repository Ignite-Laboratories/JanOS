package common

type Node struct {
	MaxWidth    int
	Value       string
	InputStream chan string
	MatchMaker  MatchMaker
	Children    map[string]*Node
}

func NewNode(value string, maxWidth int, matchMaker MatchMaker) *Node {
	n := &Node{
		MaxWidth:    maxWidth,
		Value:       value,
		InputStream: make(chan string),
		MatchMaker:  matchMaker,
		Children:    make(map[string]*Node),
	}

	go func(node *Node) {
		for data := range node.InputStream {
			if data[:len(node.Value)] == node.Value {
				// Check if we are at full perceptive width
				if len(data) == len(node.Value) {
					// If so, we identified a pattern and can emit the result to StdOut
					node.MatchMaker.PatternIdentified(data)
					continue
				}

				// No?  Okay, check if our kids know how to handle this...
				d := data[:len(node.Value)+1]
				child, ok := node.Children[d]
				if ok {
					// Yup, we have a child who knows what to do, pass it along
					child.InputStream <- data
				} else {
					// Nope, anyone want to get it on?
					go node.MatchMaker.SpawnRequest(node, d)
				}
			}
		}
	}(n)

	return n
}

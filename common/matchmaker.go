package common

import (
	"log"
	"sync"
	"time"
)

type MatchMaker interface {
	HandleImpulse(data string)
	SpawnRequest(node *Node, perceivedValue string)
	PatternIdentified(data string)
}

type StdMatchMaker struct {
	activeCount        int
	nodeCount          int
	maxPerceptiveWidth int
	spawnPoolSize      int
	hasRequest         bool
	requestTime        time.Time
	lastNode           *Node
	lastValue          string
	mutex              sync.Mutex
	children           map[string]chan string
	InputStream        chan string
	lastIdentified     string
}

func NewStdMatchMaker(spawnPoolSize int, maxPerceptiveWidth int) *StdMatchMaker {
	mm := &StdMatchMaker{
		spawnPoolSize:      spawnPoolSize,
		maxPerceptiveWidth: maxPerceptiveWidth,
		mutex:              sync.Mutex{},
		children:           make(map[string]chan string),
		InputStream:        make(chan string),
	}

	go func() {
		for data := range mm.InputStream {
			mm.HandleImpulse(data)
		}
	}()

	return mm
}

func (mm *StdMatchMaker) HandleImpulse(data string) {
	// If the data is empty, ignore
	if IsNullOrWhitespace(data) {
		return
	}

	// We only store the primal nodes off the first character
	firstChar := data[:1]

	// Have we seen this character yet?
	channel, ok := mm.children[firstChar]
	if ok {
		// Great, send the full data down that channel (sends it to the next node in the sequence)
		channel <- data

		// TODO: Consider waitgrouping this to synchronize multiple neural pathways at once
	} else {
		// No?  Okay, let's build a channel to filter this data into
		mm.children[firstChar] = make(chan string)
		// Now let's spawn off the primal nodes for this channel
		for i := 0; i < mm.spawnPoolSize; i++ {
			log.Printf("[%d Nodes] - Spawn - %s", mm.nodeCount, firstChar)
			n := NewNode(firstChar, mm.maxPerceptiveWidth, mm)
			mm.nodeCount++
			// Finally, launch off a routine to forward on the data off this channel to each primal node
			go func(c chan string) {
				for d := range c {
					n.InputStream <- d
				}
			}(mm.children[firstChar])
		}
	}
}

func (mm *StdMatchMaker) SpawnRequest(node *Node, perceivedValue string) {
	mm.mutex.Lock()
	now := time.Now()

	// Stale out the old request, if necessary
	if mm.hasRequest && now.Sub(mm.requestTime).Milliseconds() > 50 {
		mm.hasRequest = false
	}

	// Then, check if we have a spawn request and if the last node had the same perceived value...
	if mm.hasRequest && mm.lastValue == perceivedValue {
		log.Printf("[%d Nodes] - Spawn - %s", mm.nodeCount, perceivedValue)
		// Great - make it happen!

		mm.hasRequest = false
		n := NewNode(perceivedValue, node.MaxWidth, mm)
		mm.nodeCount++
		node.Children[perceivedValue] = n        // Assign to parent A
		mm.lastNode.Children[perceivedValue] = n // Assign to parent B
	} else {
		// Otherwise, no match - degrade out the old match request
		mm.lastValue = perceivedValue
		mm.lastNode = node
		mm.requestTime = now
		mm.hasRequest = true
	}
	mm.mutex.Unlock()
}

func (mm *StdMatchMaker) PatternIdentified(data string) {
	// Don't repeat duplicate output
	if mm.lastIdentified == data {
		return
	}

	mm.lastIdentified = data
	log.Printf("[%d Nodes] - Match - %s", mm.nodeCount, mm.lastIdentified)
}

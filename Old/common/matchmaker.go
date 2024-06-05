package common

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type MatchMaker interface {
	HandleImpulse(data string)
	SpawnRequest(node *Node, perceivedValue string)
	PatternIdentified(data string)
	Log(data string)
	LogWithStats(data string, logType string)
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
	children           map[string]*Broker[string]
	InputStream        chan string
	lastIdentified     string
}

func NewStdMatchMaker(spawnPoolSize int, maxPerceptiveWidth int) *StdMatchMaker {
	mm := &StdMatchMaker{
		spawnPoolSize:      spawnPoolSize,
		maxPerceptiveWidth: maxPerceptiveWidth,
		mutex:              sync.Mutex{},
		children:           make(map[string]*Broker[string]),
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

	mm.mutex.Lock()
	// We only store the primal nodes off the first character
	firstChar := data[:1]

	// Have we seen this character yet?
	broker, ok := mm.children[firstChar]
	if ok {
		// Great, send the full data down that channel (sends it to all the primal nodes that match this data)
		broker.Publish(data)
	} else {
		// No?  Okay, let's build a channel to filter this data into
		b := NewBroker[string]()
		mm.children[firstChar] = b
		go b.Start()
		// Now let's spawn off the primal nodes for this channel
		for i := 0; i < mm.spawnPoolSize; i++ {
			mm.LogWithStats(firstChar, "Primal Spawn")
			n := NewNode(firstChar, mm.maxPerceptiveWidth, mm)
			go n.ConnectBroker(b)
			mm.nodeCount++
		}
	}
	mm.mutex.Unlock()
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
		mm.LogWithStats(perceivedValue, "Spawn")
		// Great - make it happen!

		mm.hasRequest = false
		n := NewNode(perceivedValue, mm.maxPerceptiveWidth, mm)
		go n.ConnectChannel(n.InputChannel)
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
	mm.LogWithStats(mm.lastIdentified, "Match")
}

func (mm *StdMatchMaker) Log(data string) {
	log.Printf("%s", fmt.Sprintf("%-36s", data))
}

func (mm *StdMatchMaker) LogWithStats(data string, logType string) {
	log.Printf("%s - [%d Nodes] - %s", fmt.Sprintf("%-36s", data), mm.nodeCount, logType)
}

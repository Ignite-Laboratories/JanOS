package Spark

import "github.com/Ignite-Laboratories/JanOS/Logic"

type Inbox struct {
	Subjects map[string][]any
}

type Nexus struct {
	subscribers map[Logic.Entity][]string
	messages    map[string][]any
	queued      map[string][]any
}

func NewNexus() Nexus {
	return Nexus{
		subscribers: make(map[Logic.Entity][]string),
		messages:    make(map[string][]any),
	}
}

// Subscribe takes in a subject and a subscribing entity to register for messages of the provided subject
func (w *Logic.World) Subscribe(subscriber Logic.Entity, subject string) {
	if w.Messaging.subscribers[subscriber] == nil {
		w.Messaging.subscribers[subscriber] = make([]string, 0)
	}
	w.Messaging.subscribers[subscriber] = append(w.Messaging.subscribers[subscriber], subject)
}

// Unsubscribe takes in a subject and a subscribing entity to remove from messages of the provided subject
func (w *Logic.World) Unsubscribe(subscriber Logic.Entity, subject string) {
	var newSubjects []string
	for _, sub := range w.Messaging.subscribers[subscriber] {
		if sub != subject {
			newSubjects = append(newSubjects, sub)
		}
	}
	w.Messaging.subscribers[subscriber] = newSubjects
}

// Publish takes in a subject and a message
func (w *Logic.World) Publish(subject string, message any) {
	w.Messaging.queued[subject] = append(w.Messaging.queued[subject], message)
}

// GetMessages takes in the subscribing entity to retrieve messages for and returns its pending messages
func (w *Logic.World) GetMessages(subscriber Logic.Entity) Inbox {
	inbox := Inbox{make(map[string][]any)}

	for _, subscription := range w.Messaging.subscribers[subscriber] {
		messages := w.Messaging.messages[subscription]
		if len(messages) > 0 {
			inbox.Subjects[subscription] = messages
		}
	}
	return inbox
}

// Cycle empties the message queue for the next iteration of the loop
func (n *Nexus) Cycle() {
	n.messages = n.queued
	n.queued = make(map[string][]any)
}

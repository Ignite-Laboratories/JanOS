package Logic

type Inbox struct {
	Subjects map[string][]any
}

type Nexus struct {
	subscribers map[Entity][]string
	messages    map[string][]any
	queued      map[string][]any
}

func NewNexus() Nexus {
	return Nexus{
		subscribers: make(map[Entity][]string),
		messages:    make(map[string][]any),
	}
}

// Subscribe takes in a subject and a subscribing entity to register for messages of the provided subject
func (w *World) Subscribe(subscriber Entity, subject string) {
	if w.Nexus.subscribers[subscriber] == nil {
		w.Nexus.subscribers[subscriber] = make([]string, 0)
	}
	w.Nexus.subscribers[subscriber] = append(w.Nexus.subscribers[subscriber], subject)
}

// Unsubscribe takes in a subject and a subscribing entity to remove from messages of the provided subject
func (w *World) Unsubscribe(subscriber Entity, subject string) {
	var newSubjects []string
	for _, sub := range w.Nexus.subscribers[subscriber] {
		if sub != subject {
			newSubjects = append(newSubjects, sub)
		}
	}
	w.Nexus.subscribers[subscriber] = newSubjects
}

// Publish takes in a subject and a message
func (w *World) Publish(subject string, message any) {
	w.Nexus.queued[subject] = append(w.Nexus.queued[subject], message)
}

// GetMessages takes in the subscribing entity to retrieve messages for and returns its pending messages
func (w *World) GetMessages(subscriber Entity) Inbox {
	inbox := Inbox{make(map[string][]any)}

	for _, subscription := range w.Nexus.subscribers[subscriber] {
		messages := w.Nexus.messages[subscription]
		if len(messages) > 0 {
			inbox.Subjects[subscription] = messages
		}
	}
	return inbox
}

// Clear empties the message queue for the next iteration of the loop
func (n *Nexus) Clear() {
	n.messages = n.queued
	n.queued = make(map[string][]any)
}

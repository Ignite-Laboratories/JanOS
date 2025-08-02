package std

// Synchro represents a way to synchronize execution across threads.
//
// To send execution using a synchro, first create one using make.  Then, Engage the synchro
// from the thread you wish to execute on.  The calling thread can then Send actions to the
// other thread for intermittent execution.
//
// NOTE: The sending thread will block and wait for execution to complete on the other thread.
//
//		 global -
//	   var synchro = make(std.Synchro)
//
//		 main loop -
//		  for ... {
//	    ...
//		   synchro.Engage()
//		   ...
//		  }
//
//		 sender -
//	   synchro.Send(func() { ... })
type Synchro chan *SyncAction

// Send sends the provided action over the synchro channel and waits for it to be executed.
func (s Synchro) Send(action func()) {
	syn := &SyncAction{Action: action}
	syn.Add(1)
	s <- syn
	syn.Wait()
}

// Engage asynchronously handles the currently incoming messages on the Synchro channel before returning control.
func (s Synchro) Engage() {
	for {
		select {
		case syn := <-s:
			syn.Action()
			syn.Done()
		default:
			return
		}
	}
}

// EngageOnce synchronously reads a single message on the Synchro channel before returning control.
func (s Synchro) EngageOnce() {
	syn := <-s
	syn.Action()
	syn.Done()
}

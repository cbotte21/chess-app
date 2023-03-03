package adt_queue

type AdtQueue [2]int

// Join adds a player to the adt_queue. Registering connection to listener.
func (queue *AdtQueue) Join(_id int) {
	queue[len(queue)-1] = _id
	if len(queue) == 2 {
		//TODO: Ping Chess server
		queue[0] = 0
		queue[1] = 0
	}
}

// Leave removes a player from the adt_queue
func (queue *AdtQueue) Leave(_id int) {
	for i := 0; i < len(queue); i++ {
		if queue[i] == _id {
			queue[i] = 0 //NULL
			queue.Format()
			break
		}
	}
}

// Format moves second element to first iff first is null
func (queue *AdtQueue) Format() {
	if queue[0] == 0 {
		queue[0] = queue[1]
	}
}

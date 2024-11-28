package queues

type Queue []*Member

type Member struct {
	Priority int
	Content  interface{}
	Name     string
}

func (m *Member) New(priority int, content interface{}, name string) {
	m.Name = name
	m.Content = content
	m.Priority = priority
}

func (queue *Queue) Add(member *Member) {
	if len(*queue) == 0 {
		*queue = append(*queue, member)
	} else {
		var appended bool
		appended = false

		var i int
		var addedMember *Member

		for i, addedMember = range *queue {
			if member.Priority > addedMember.Priority {
				*queue = append((*queue)[:i], append(Queue{member}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, member)
		}
	}
}

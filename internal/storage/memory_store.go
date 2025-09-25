package storage

type MemoryStore struct {
	contacts map[int]Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]Contact),
		nextID:   1,
	}
}

func (m *MemoryStore) Ajouter(c Contact) Contact {
	c.ID = m.nextID
	m.contacts[m.nextID] = c
	m.nextID++
	return c
}

func (m *MemoryStore) Lister() []Contact {
	list := []Contact{}
	for _, c := range m.contacts {
		list = append(list, c)
	}
	return list
}

func (m *MemoryStore) Supprimer(ID int) bool {
	if _, ok := m.contacts[ID]; ok {
		delete(m.contacts, ID)
		return true
	}
	return false
}

func (m *MemoryStore) Recuperer(ID int) (Contact, bool) {
	c, ok := m.contacts[ID]
	return c, ok
}

func (m *MemoryStore) MettreAJour(c Contact) (Contact, bool) {
	if _, ok := m.contacts[c.ID]; ok {
		m.contacts[c.ID] = c
		return c, true
	}
	return Contact{}, false
}

func (m *MemoryStore) NextID() int {
	return m.nextID
}

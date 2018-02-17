package memento

// 要保存的info type
type Command interface {
	GetValue() interface{}
}

// 实现1
type Value byte

func (v Value) GetValue() interface{} {
	return v
}

// 实现2
type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

// 抽象层
type Memento struct {
	memento Command
}

// processor
type Originnator struct {
	Command Command
}

func (o *Originnator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *Originnator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

// snapshot container
type CareTracker struct {
	MementoList []Memento
}

func (c *CareTracker) Add(m Memento) {
	c.MementoList = append(c.MementoList, m)
}

func (c *CareTracker) Pop() Memento {
	length := len(c.MementoList)
	if length > 0 {
		tempMemento := c.MementoList[length-1]
		c.MementoList = c.MementoList[0 : length-1]
		return tempMemento
	}
	return Memento{}
}

// 层层包裹, 只暴露set 和 restore
type MementoFacade struct {
	originator Originnator
	carTracker CareTracker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.carTracker.Add(m.originator.NewMemento())
}

func (m *MementoFacade) StoreSettings(index int) Command {
	m.originator.ExtractAndStoreCommand(m.carTracker.Pop())
	return m.originator.Command
}

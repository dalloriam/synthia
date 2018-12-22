package modular

type Trigger struct {
	ShouldTrigger bool
}

func NewTrigger() *Trigger {
	return &Trigger{ShouldTrigger: false}
}

func (t *Trigger) Stream() float64 {
	if t.ShouldTrigger {
		t.ShouldTrigger = false
		return 1
	}
	return 0
}

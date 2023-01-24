package helper

type Observer interface {
	Update()
}

type Observable interface {
	AddObserver(o Observer)
	RemoveObserver(o Observer)
	Notify()
}

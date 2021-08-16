package downloadsuite

type URL string

type Watcher interface {
	Watch(URLs []URL)
}

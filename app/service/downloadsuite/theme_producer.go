package downloadsuite

import (
	"github.com/nsqio/go-nsq"
)

type Producer interface {
	Produce(producer *nsq.Producer, topic string) error
}

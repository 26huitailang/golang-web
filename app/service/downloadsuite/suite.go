/*
Package downloadsuite to download one downloadsuite images.

	Now support meituri.
*/
package downloadsuite

import (
	"github.com/nsqio/go-nsq"
)

type Downloader interface {
	Download()
}
type ISuiteOperator interface {
	Downloader
	Producer
}

// todo: 没有任何关于这个struct的方法
// 是否应该把suite方法抽象到这里，接口，mtr等去实现这个接口
type Suite struct {
	Operator ISuiteOperator
}

func NewSuite(operator ISuiteOperator) *Suite {
	return &Suite{
		Operator: operator,
	}
}

func (s *Suite) Download() {
	s.Operator.Download()
}

func (s *Suite) Produce(producer *nsq.Producer, topic string) error {
	err := s.Operator.Produce(producer, topic)
	return err
}

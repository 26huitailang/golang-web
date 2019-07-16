/*
Package downloadsuite to download one downloadsuite images.

	Now support meituri.
*/
package downloadsuite

// ISuiteOperator interface 定义了Suite的操作集合
type ISuiteOperator interface {
	//GetPageURLs(chan string)
	//GetImgURLs(chPage <-chan string, chFailedImg <-chan string) <-chan string
	Download(isTheme bool)
	//GetOrgURL() string
}

type Suite struct {
	Operator ISuiteOperator
}

func NewSuite(operator ISuiteOperator) *Suite {
	return &Suite{
		Operator: operator,
	}
}

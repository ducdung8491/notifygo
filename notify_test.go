package notifygo

import "fmt"

type sms struct{}

func (s *sms) Notify(msg string) error {
	fmt.Println("send sms")
	return nil
}

type email struct{}

func (e *email) Notify(msg string) error {
	fmt.Println("send email")
	return nil
}

func Example() {
	nt := Pipe(
		&sms{},
		&email{},
		NotifyFunc(func(msg string) error {
			fmt.Println("send slack")
			return nil
		}),
	)

	nt.Notify("Xin Chao dung")
	//Output:
	//send sms
	//send email
	//send slack
}

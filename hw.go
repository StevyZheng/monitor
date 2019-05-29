package main

type IOError struct {
}

type Ipmi struct {
	CpuTemp       []int
	MemTemp       []int
	PchTemp       int
	GpuTemp       []int
	FanSpeed      []int
	ErrorEventLog []string
}

func (e *Ipmi) Fill() {

}

type ErrorMsg struct {
	Time    string
	MsgBody string
	Details string
}

func NewErrorMsg() *ErrorMsg {
	return new(ErrorMsg)
}

func NewErrorMsgInit(time, msgBody, details string) *ErrorMsg {
	return &ErrorMsg{Time: time, MsgBody: msgBody, Details: details}
}

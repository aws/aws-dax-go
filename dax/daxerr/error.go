package daxerr

import "github.com/aws/aws-sdk-go/aws/awserr"

type DaxRequestFailure struct {
	awserr.RequestFailure
	Codes []int
}

type DaxTransactionCanceledFailure struct {
	DaxRequestFailure
	CancellationReasonCodes []string
	CancellationReasonMsgs  []string
	CancellationReasonItems []byte
}

func (f *DaxRequestFailure) Recoverable() bool {
	return len(f.Codes) > 0 && f.Codes[0] == 2
}

func (f *DaxRequestFailure) AuthError() bool {
	return len(f.Codes) > 3 && (f.Codes[1] == 23 && f.Codes[2] == 31 &&
		(f.Codes[3] == 32 || f.Codes[3] == 33 || f.Codes[3] == 34))
}

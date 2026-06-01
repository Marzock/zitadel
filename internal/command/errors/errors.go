package errors

type WrongPasswordError struct {
	FailedAttempts int32
}

func (wpe *WrongPasswordError) Error() string {
	return ""
}

type SuspensionNotExceededError struct {
	RemainingTime int32
}

func (sne *SuspensionNotExceededError) Error() string {
	return ""
}

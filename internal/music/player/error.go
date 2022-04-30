package player

type ErrorCode int

const (
	NotConnected ErrorCode = iota + 1
	SongNotFound
	StorageQueryFailed
	ConnectFailed
	NotImplemented
)

var (
	ErrNotConnected       = &Error{code: NotConnected}
	ErrSongNotFound       = &Error{code: SongNotFound}
	ErrStorageQueryFailed = &Error{code: StorageQueryFailed}
	ErrConnectFailed      = &Error{code: ConnectFailed}
	ErrNotImplemented     = &Error{code: NotImplemented}
)

type Error struct {
	code ErrorCode
	msg  string
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Wrap(msg string) *Error {
	e.msg = msg
	return e
}
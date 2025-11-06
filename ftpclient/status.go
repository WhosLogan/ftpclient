package ftpclient

import "errors"

const (
	statusReady               = 220
	statusNeedPassword        = 331
	statusServiceNotAvailable = 421
	statusUnknownCommand      = 500
	statusSyntaxError         = 501
	statusNotLoggedIn         = 530
	statusEnteringPassive     = 227
)

func translateErrorStatusCode(code int) error {
	switch code {
	case statusNeedPassword:
		return errors.New("ftp server needs password")
	case statusServiceNotAvailable:
		return errors.New("ftp service not available")
	case statusUnknownCommand:
		return errors.New("ftp command unknown")
	case statusSyntaxError:
		return errors.New("ftp command syntax error")
	case statusNotLoggedIn:
		return errors.New("ftp server reports invalid username or password")
	}

	return errors.New("unknown error")
}

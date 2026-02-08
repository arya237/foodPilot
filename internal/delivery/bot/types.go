package bot


type state int

const (
	startOfStates state = iota
	idel
	startResturantLogin
	waitingForUsername
	waitingForPassword
	endOfState
)

type userState struct {
	userID   int64
	state    state
	username string
}

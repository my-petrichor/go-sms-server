package connector

//Dispatcher is a interface that could send verification code to a mobile
type Dispatcher interface {
	Send(mobile string, length int) error
}

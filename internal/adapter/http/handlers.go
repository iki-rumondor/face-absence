package customHTTP

type Handlers struct {
	StudentHandler *StudentHandlers
	AuthHandler    *AuthHandlers
	TeacherHandler *TeacherHandlers
	ClassHandler   *ClassHandler
}

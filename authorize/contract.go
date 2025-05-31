package authorize

type AccessManager interface {
	Check(subject, resource, action string) (bool, error)
}

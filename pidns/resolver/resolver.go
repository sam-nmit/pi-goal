package resolver

type INameResolver interface {
	AddRule(name, host string)
	Resolve(name string) *[]string

	Hosts() map[string]string
}

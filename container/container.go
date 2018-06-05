package aos

var containers map[string]interface{}

func ContainerGet(name string)interface{} {
	if nil == containers {
		containers = make(map[string]interface{})
		return nil
	}
	return containers[name]
}

func ContainerSet(name string, object interface{}) {
	if nil == containers {
		containers = make(map[string]interface{})
	}
	containers[name] = object
}

func ContainerHas(name string)bool {
	if _, ok := containers[name]; ok {
		return ok
	} else {
		return false
	}
}
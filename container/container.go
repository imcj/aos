package aos

import "github.com/apex/log"

var containers map[string]interface{}

func ContainerGet(name string) interface{} {
	if nil == containers {
		containers = make(map[string]interface{})
		return nil
	}
	return containers[name]
}

func GetLogger(name string) *log.Entry {
	if nil == containers {
		containers = make(map[string]interface{})
		return nil
	}
	logger := containers[name].(*log.Entry)
	return logger
}

func ContainerSet(name string, object interface{}) {
	if nil == containers {
		containers = make(map[string]interface{})
	}
	containers[name] = object
}

func ContainerHas(name string) bool {
	if _, ok := containers[name]; ok {
		return ok
	} else {
		return false
	}
}

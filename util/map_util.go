package util

func DragStringKeysFromMap(ssm map[string]string) []string {
	if nil == ssm {
		return nil
	}
	rl := make([]string, 0)
	for k := range ssm {
		rl = append(rl, k)
	}
	return rl
}

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

func DragStringValuesFromMap(ssm map[string]string) []string {
	if nil == ssm {
		return nil
	}
	tm := make(map[string]bool)
	rl := make([]string, 0)
	for _, v := range ssm {
		_, ok := tm[v]
		if !ok {
			rl = append(rl, v)
			tm[v] = true
		}
	}
	return rl
}

//
//  Vector2Map
//  @Description: 数组转map
//  @param arr
//  @return map[string]any
//
func Vector2Map(arr []string) map[string]any {
	rm := make(map[string]any)
	if nil == arr {
		return rm
	}
	for _, key := range arr {
		rm[key] = true
	}
	return rm
}

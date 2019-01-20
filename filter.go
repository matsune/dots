package dots

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func removeDupsString(arr []string) []string {
	encountered := map[string]bool{}
	res := []string{}
	for _, v := range arr {
		if !encountered[v] {
			encountered[v] = true
			res = append(res, v)
		}
	}
	return res
}

func removeDupsInt(arr []int) []int {
	encountered := map[int]bool{}
	res := []int{}
	for _, v := range arr {
		if !encountered[v] {
			encountered[v] = true
			res = append(res, v)
		}
	}
	return res
}

// filter targets by names
func filterNames(all []Target, names []string) []Target {
	if len(names) == 0 {
		return all
	}
	names = removeDupsString(names)

	indices := []int{}
	for _, n := range names {
		for idx, t := range all {
			if t.Name == n {
				indices = append(indices, idx)
			}
		}
	}
	indices = removeDupsInt(indices)

	res := make([]Target, len(indices))
	for i, v := range indices {
		res[i] = all[v]
	}
	return res
}

// filter targets by tags
func filterTags(all []Target, tags []string) []Target {
	if len(tags) == 0 {
		var res []Target
		for _, t := range all {
			if len(t.Tags) == 0 {
				res = append(res, t)
			}
		}
		return res
	}

	tags = removeDupsString(tags)

	indices := []int{}
	for _, tag := range tags {
		for idx, t := range all {
		tagsLoop:
			for _, tt := range t.Tags {
				if tt == tag {
					indices = append(indices, idx)
					break tagsLoop
				}
			}
		}
	}
	indices = removeDupsInt(indices)

	res := make([]Target, len(indices))
	for i, v := range indices {
		res[i] = all[v]
	}
	return res
}

func filter(all []Target, targets, tags []string) []Target {
	all = filterNames(all, targets)
	return filterTags(all, tags)
}

package version

import "github.com/blang/semver"

type FinalBump struct{
	Pre string
}

func (bump FinalBump) Apply(v semver.Version) semver.Version {
	if v.Pre == nil || bump.Pre != "" {
		v = PatchBump{bump.Pre}.Apply(v)
	} else {
		v.Pre = nil
	}
	return v
}

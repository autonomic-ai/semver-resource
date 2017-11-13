package version

import "github.com/blang/semver"

type PatchBump struct{
	Pre string
}

func (bump PatchBump) Apply(v semver.Version) semver.Version {
	v.Patch++
	v.Pre = nil
	if bump.Pre != "" {
		v = PreBump{bump.Pre}.init(v)
	}
	return v
}

package version

import "github.com/blang/semver"

type MinorBump struct{
	Pre string
}

func (bump MinorBump) Apply(v semver.Version) semver.Version {
	v.Minor++
	v.Patch = 0
	v.Pre = nil
	if bump.Pre != "" {
		v = PreBump{bump.Pre}.init(v)
	}
	return v
}

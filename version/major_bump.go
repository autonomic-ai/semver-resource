package version

import "github.com/blang/semver"

type MajorBump struct{
	Pre string
}

func (bump MajorBump) Apply(v semver.Version) semver.Version {
	v.Major++
	v.Minor = 0
	v.Patch = 0
	v.Pre = nil
	if bump.Pre != "" {
		v = PreBump{bump.Pre}.init(v)
	}
	return v
}

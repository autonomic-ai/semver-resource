package version

import "github.com/blang/semver"

type PreBump struct {
	Pre string
}

func (bump PreBump) Apply(v semver.Version) semver.Version {
	if v.Pre == nil {
		v = PatchBump{bump.Pre}.Apply(v)
	} else if v.Pre[0].VersionStr != bump.Pre {
		v = bump.init(v)
	} else {
		v.Pre = []semver.PRVersion{
			{VersionStr: bump.Pre},
			{VersionNum: v.Pre[1].VersionNum + 1, IsNum: true},
		}
	}

	return v
}

func (bump PreBump) init(v semver.Version) semver.Version {
	v.Pre = []semver.PRVersion{
		{VersionStr: bump.Pre},
		{VersionNum: 1, IsNum: true},
	}
	return v
}

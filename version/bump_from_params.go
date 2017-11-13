package version

func BumpFromParams(bumpStr string, preStr string) Bump {
	var bump Bump

	switch bumpStr {
	case "major":
		bump = MajorBump{preStr}
	case "minor":
		bump = MinorBump{preStr}
	case "patch":
		bump = PatchBump{preStr}
	case "final":
		bump = FinalBump{preStr}
	default:
		if preStr != "" {
			bump = PreBump{preStr}
		} else {
			bump = IdentityBump{}
		}
	}

	return bump
}

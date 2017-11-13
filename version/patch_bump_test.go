package version_test

import (
	"github.com/blang/semver"
	"github.com/concourse/semver-resource/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PatchBump", func() {
	var inputVersion semver.Version
	var bump version.PatchBump
	var outputVersion semver.Version

	BeforeEach(func() {
		inputVersion = semver.Version{
			Major: 1,
			Minor: 2,
			Patch: 3,
			Pre: []semver.PRVersion{
				{VersionStr: "beta"},
				{VersionNum: 1, IsNum: true},
			},
		}

		bump = version.PatchBump{}
	})

	JustBeforeEach(func() {
		outputVersion = bump.Apply(inputVersion)
	})

	Context("when the bump does not contain a prerelease", func() {

		It("bumps patch and zeroes out the subsequent segments", func() {
			Expect(outputVersion).To(Equal(semver.Version{
				Major: 1,
				Minor: 2,
				Patch: 4,
			}))
		})
	})

	Context("when the bump contains a prerelease", func() {
		BeforeEach(func() {
			bump.Pre = "beta"
		})

		It("bumps patch and bumps to version 1 of the prerelease type", func() {
			Expect(outputVersion).To(Equal(semver.Version{
				Major: 1,
				Minor: 2,
				Patch: 4,
				Pre: []semver.PRVersion{
					{VersionStr: "beta"},
					{VersionNum: 1, IsNum: true},
				},
			}))
		})
	})
})

package version_test

import (
	"github.com/blang/semver"
	"github.com/concourse/semver-resource/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MajorBump", func() {
	var inputVersion semver.Version
	var bump version.MajorBump
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

		bump = version.MajorBump{}
	})

	JustBeforeEach(func() {
		outputVersion = bump.Apply(inputVersion)
	})

	Context("when the bump does not contain a prerelease", func() {

		It("bumps major and zeroes out the subsequent segments", func() {
			Expect(outputVersion).To(Equal(semver.Version{
				Major: 2,
				Minor: 0,
				Patch: 0,
			}))
		})
	})

	Context("when the bump contains a prerelease", func() {
		BeforeEach(func() {
			bump.Pre = "beta"
		})

		It("bumps major and zeroes out the subsequent segments, and bumps to version 1 of the prerelease type", func() {
			Expect(outputVersion).To(Equal(semver.Version{
				Major: 2,
				Minor: 0,
				Patch: 0,
				Pre: []semver.PRVersion{
					{VersionStr: "beta"},
					{VersionNum: 1, IsNum: true},
				},
			}))
		})
	})
})

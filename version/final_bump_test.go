package version_test

import (
	"github.com/blang/semver"
	"github.com/concourse/semver-resource/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FinalBump", func() {
	var inputVersion semver.Version
	var bump version.FinalBump
	var outputVersion semver.Version

	BeforeEach(func() {
		inputVersion = semver.Version{
			Major: 1,
			Minor: 2,
			Patch: 3,
		}

		bump = version.FinalBump{}
	})

	JustBeforeEach(func() {
		outputVersion = bump.Apply(inputVersion)
	})

	Context("when the version is a prerelease", func() {
		BeforeEach(func() {
			inputVersion.Pre = []semver.PRVersion{
				{VersionStr: "beta"},
				{VersionNum: 1, IsNum: true},
			}
		})

		Context("when the bump does not contain a prerelease", func() {

			It("lops off the pre segment", func() {
				Expect(outputVersion).To(Equal(semver.Version{
					Major: 1,
					Minor: 2,
					Patch: 3,
				}))
			})
		})

		Context("when the bump contains a prerelease", func() {
			BeforeEach(func() {
				bump.Pre = "beta"
			})

			It("bumps to the next patch version, and to version 1 of the prerelease type", func() {
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

	Context("when the version is not a prerelease", func() {
		BeforeEach(func() {
			inputVersion.Pre = nil
		})

		Context("when the bump does not contain a prerelease", func() {

			It("bumps to the next patch version", func() {
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

			It("bumps to the next patch version, and to version 1 of the prerelease type", func() {
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
})

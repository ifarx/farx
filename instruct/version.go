/**
 * @create  2019/12/28 13:33
 * @since 1.0.0.0
 * @author  ant
 */
package instruct

import "strconv"

type Version struct {
	major    uint16
	minor    uint16
	revision uint16
	build    uint16
}

func NewVersion(major, minor, revision, build uint16) *Version {
	return &Version{major, minor, revision, build}
}

func NewSimpleVersion(major, minor, revision uint16) *Version {
	return &Version{major, minor, revision, 0}
}

func (v *Version) Major() uint16 {
	return v.major
}

func (v *Version) Minor() uint16 {
	return v.minor
}

func (v *Version) Revision() uint16 {
	return v.revision
}

func (v *Version) Build() uint16 {
	return v.build
}

func (v *Version) String() string {
	if v.build > 0 {
		return strconv.FormatUint(uint64(v.major), 10) + "." + strconv.FormatUint(uint64(v.minor), 10) + "." + strconv.FormatUint(uint64(v.revision), 10) + "." + strconv.FormatUint(uint64(v.build), 10)
	}
	return strconv.FormatUint(uint64(v.major), 10) + "." + strconv.FormatUint(uint64(v.minor), 10) + "." + strconv.FormatUint(uint64(v.revision), 10)
}

func (v *Version) Equal(v1 *Version) bool {
	return (v.major == v1.major) &&
		(v.minor == v1.minor) &&
		(v.revision == v1.revision) &&
		(v.build == v1.build)
}

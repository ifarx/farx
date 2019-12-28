/**
 * @create  2019/12/28 11:31
 * @since 1.0.0.0
 * @author  ant
 */
package instruct

/**
 * | 8bit | 8bit | 8bit | 8bit |
 */

type Magic struct {
	bytes [4]byte
}

func NewMagic(b1, b2, b3, b4 byte) *Magic {
	return &Magic{bytes: [4]byte{b1, b2, b3, b4}}
}

func (m *Magic) Equals(m1 *Magic) bool {
	return m.bytes[0] == m1.bytes[0] &&
		m.bytes[1] == m1.bytes[1] &&
		m.bytes[2] == m1.bytes[2] &&
		m.bytes[3] == m1.bytes[3]
}

func (m *Magic) String() string {
	return string(m.bytes[:])
}

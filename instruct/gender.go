/**
 * @create  2019/12/28 17:18
 * @since 1.0.0.0
 * @author  ant
 */
package instruct

type Gender byte

//未知
const GENDER_UNKNOWN = Gender(0)

//男性
const GENDER_MALE = Gender(1)

//女性
const GENDER_FEMALE = Gender(2)

func (g Gender) Byte() byte {
	return byte(g)
}

func (g Gender) IsMale() bool {
	return g == GENDER_MALE
}

func (g Gender) IsFemale() bool {
	return g == GENDER_FEMALE
}

func ToGender(b byte) Gender {
	return Gender(b)
}

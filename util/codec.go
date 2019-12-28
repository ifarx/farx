/**
 * @create  2019/12/28 16:15
 * @since 1.0.0.0
 * @author  ant
 */
package util

type Codec interface {
	Unmarshal(bys []byte, in interface{}) error
	Marshal(in interface{}) ([]byte, error)
}

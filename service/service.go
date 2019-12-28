/**
 * @create  2019/12/28 11:28
 * @since 1.0.0.0
 * @author  ant
 */
package service

import "src/instruct"

type Service interface {
	Component
	Magic() *instruct.Magic
	Version() *instruct.Version
}

/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    ip_test
 * @Date:    2022/5/25 23:58
 * @package: extip
 * @Version: x.x.x
 *
 * @Description: xxx
 *
 */

package extip

import (
	"fmt"
	"testing"
)

func TestGetMyIp(t *testing.T) {
	fmt.Println(GetMyIp("en8"))
}

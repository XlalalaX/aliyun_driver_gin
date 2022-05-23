package heartbeat

import (
	"main/Control_center/global"
	"net/http"
)

func Hearbeat() {
	c := http.Client{}

	for s, _ := range global.M_ip {
		r, err := http.NewRequest("POST", "http://"+s+"/hearbeat", nil)
		if err != nil {
			global.M_ip[s]++
		}
		r.Header.Set("token", global.User.RefreshToken)
		r.Header.Set("root_path", global.Root_path)
		re, err := c.Do(r)
		if err != nil || re.StatusCode != http.StatusOK {
			global.M_ip[s]++
		}
		if global.M_ip[s] >= 6 {
			delete(global.M_ip, s)
		}
	}
}

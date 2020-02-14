package service

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceStat struct{}

// GetRateLimit is sending a neutral request that make possible to read the current RateLimit.
func (ServiceStat) GetRateLimit() (*model.RateLimit, error) {

	r, err := http.NewRequest(http.MethodGet, BaseURL+"clients/", nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", BasicAuth)
	r.Header.Set("Content-Type", "application/json")

	client := DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dump, _ := httputil.DumpResponse(resp, false)
	res := fmt.Sprintln(string(dump))
	lines := strings.Split(res, "\n")

	var rateLimit model.RateLimit

	for k := range lines {
		rateLimitFound, rateLimitRemainingFound, rateLimitResetFound := 0, 0, 0
		if keyVal := strings.Split(lines[k], ":"); len(keyVal) == 2 {
			res, _ := strconv.Atoi(keyVal[1])
			switch keyVal[0] {
			case "X-RateLimit":
				rateLimitFound++
				switch rateLimitFound {
				case 1:
					rateLimit.Limit1 = res
				case 2:
					rateLimit.Limit2 = res
				case 3:
					rateLimit.Limit3 = res
				case 4:
					rateLimit.Limit4 = res
				default:
					continue
				}
			case "X-RateLimit-Remaining":
				rateLimitRemainingFound++
				switch rateLimitRemainingFound {
				case 1:
					rateLimit.LimitRemaining1 = res
				case 2:
					rateLimit.LimitRemaining2 = res
				case 3:
					rateLimit.LimitRemaining3 = res
				case 4:
					rateLimit.LimitRemaining4 = res
				default:
					continue
				}
			case "X-RateLimit-Reset":
				rateLimitResetFound++
				switch rateLimitResetFound {
				case 1:
					rateLimit.LimitReset1 = res
				case 2:
					rateLimit.LimitReset2 = res
				case 3:
					rateLimit.LimitReset3 = res
				case 4:
					rateLimit.LimitReset4 = res
				default:
					continue
				}
			default:
				continue
			}
		}
	}
	return &rateLimit, nil
}

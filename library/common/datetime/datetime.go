/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-07
 * File: datetime.go
 * Desc: parse json time.Time -> datetime
 */

package datetime

import "time"

type Datetime time.Time

// MarshalJSON datetime
func (d Datetime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.DateTime)+len(`""`))
	b = append(b, '"')
	b = time.Time(d).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON datetime
func (d *Datetime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.Parse(`"`+time.DateTime+`"`, string(data))
	*d = Datetime(now)
	return
}

// String 格式化为string
func (d Datetime) String() string {
	return time.Time(d).Format(time.DateTime)
}

package traffic

import (
	"fmt"
	"strconv"
	"strings"
)

type Option struct {
	Green int64
	All   int64
}

func ParseOption(s string) (opt Option, err error) {
	split := strings.Split(s, "/")
	if len(split) != 2 {
		err = fmt.Errorf("traffic: invalid string s %s", s)
		return
	}
	green, _err := strconv.ParseInt(split[0], 10, 64)
	if _err != nil {
		err = _err
		return
	}
	all, _err := strconv.ParseInt(split[1], 10, 64)
	if _err != nil {
		err = _err
		return
	}
	opt.Green, opt.All = green, all
	return
}

func NewTrafficLight(option *Option) (*Light, error) {
	if option.All < option.Green || option.Green < 0 {
		return nil, fmt.Errorf("invalid param: all %d is le green %d or green is le zero", option.All, option.Green)
	}
	return &Light{
		all:   option.All,
		green: option.Green,
		cur:   1,
	}, nil
}

type Light struct {
	cur, green, all int64
}

func (d *Light) IsGreen() (b bool) {
	if d.green == 0 {
		return false
	}
	if d.green == d.all {
		return true
	}
	if d.cur == d.all {
		d.cur = 1
		return false
	}
	if d.cur <= d.green {
		b = true
	}
	d.cur++
	return
}

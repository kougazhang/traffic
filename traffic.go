package traffic

import "fmt"

type TrafficLightOption struct {
	Green int
	All   int
}

func NewTrafficLight(option *TrafficLightOption) (*TrafficLight, error) {
	if option.All < option.Green || option.Green < 0 {
		return nil, fmt.Errorf("invalid param: all %d is le green %d or green is le zero", option.All, option.Green)
	}
	return &TrafficLight{
		all:   option.All,
		green: option.Green,
		cur:   1,
	}, nil
}

type TrafficLight struct {
	cur, green, all int
}

func (d *TrafficLight) IsGreen() (b bool) {
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

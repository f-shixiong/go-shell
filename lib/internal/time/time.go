package time

import "time"

var StuMap = map[string]interface{}{
	"Ticker":     time.Ticker{},
	"ParseError": time.ParseError{},
	"Time":       time.Time{},
	"Timer":      time.Timer{},
	"Location":   time.Location{},
}

var FucMap = map[string]interface{}{
	"FixedZone":              time.FixedZone,
	"NewTicker":              time.NewTicker,
	"Since":                  time.Since,
	"Until":                  time.Until,
	"LoadLocationFromTZData": time.LoadLocationFromTZData,
	"LoadLocation":           time.LoadLocation,
	"ParseDuration":          time.ParseDuration,
	"Sleep":                  time.Sleep,
	"Tick":                   time.Tick,
	"Unix":                   time.Unix,
	"Parse":                  time.Parse,
	"AfterFunc":              time.AfterFunc,
	"ParseInLocation":        time.ParseInLocation,
	"Now":                    time.Now,
	"Date":                   time.Date,
	"NewTimer":               time.NewTimer,
}

package companies

import "time"

type Companies struct {
	Id         int
	Name       string
	Active     int
	Created_at time.Time
	Created_by string
	Updated_at time.Time
	Updated_by string
}

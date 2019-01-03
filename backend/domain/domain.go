package domain

import (
	"encoding/json"
	"time"
)

// User is user
type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Time represent custom time type for custom marshaling
type Time time.Time

// MarshalJSON converts custom time to json string
func (t *Time) MarshalJSON() ([]byte, error) {
	time := time.Time(*t)
	strTime := time.Format("02/01/2006")
	return json.Marshal(strTime)
}

// UnmarshalJSON converts json string to custom time
func (t *Time) UnmarshalJSON(data []byte) error {
	var strTime string
	err := json.Unmarshal(data, &strTime)
	if err != nil {
		return err
	}
	time, err := time.Parse("02/01/2006", strTime)
	if err != nil {
		return err
	}
	*t = Time(time)
	return nil
}

// Debt is debt
type Debt struct {
	ID       int     `json:"id,omitempty"`
	Creditor User    `json:"creditor,omitempty"`
	Debtor   User    `json:"debtor,omitempty"`
	Sum      float64 `json:"sum,omitempty"`
	Reason   string  `json:"reason,omitempty"`
	Date     *Time   `json:"date,omitempty"`
}

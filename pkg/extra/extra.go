package extra

import (
	"encoding/json"
	"strings"
	"time"
)

type ChargeStatus string

var (
	ChargeStatusSubmittedForSettlement ChargeStatus = "submitted_for_settlement"
	ChargeStatusSettled                ChargeStatus = "settled"
	ChargeStatusVoid                   ChargeStatus = "void"
)

type Boolean uint8

var (
	False Boolean = 0
	True  Boolean = 1
)

func (b Boolean) FromBool(input bool) Boolean {
	if input {
		return True
	}
	return False
}

func (b Boolean) AsBool() bool {
	return b == True
}

type YesOrNo string

var (
	Yes YesOrNo = "yes"
	No  YesOrNo = "no"
)

func (y YesOrNo) AsBool() bool {
	return y == Yes
}

func (y YesOrNo) String() string {
	return string(y)
}

type ChargeCardLevel string

var (
	ChargeCardLevel1 ChargeCardLevel = "LEVEL1"
	ChargeCardLevel2 ChargeCardLevel = "LEVEL2"
	ChargeCardLevel3 ChargeCardLevel = "LEVEL3"
)

type Currency string

var CurrencyUSD Currency = "usd"

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(time.RFC3339))
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	format := time.DateOnly
	if strings.Contains(str, "T") {
		format = time.RFC3339
	} else if strings.Contains(str, " ") {
		format = time.DateTime
	}

	parsed, err := time.Parse(format, str)
	if err != nil {
		return err
	}

	d.Time = parsed

	return nil
}

func (d DateTime) String() string {
	return d.Format(time.DateTime)
}

func (d *DateTime) UnmarshalText(data []byte) error {
	str := string(data)

	format := time.DateOnly
	if strings.Contains(str, "T") {
		format = time.RFC3339
	} else if strings.Contains(str, " ") {
		format = time.DateTime
	}

	parsed, err := time.Parse(format, str)
	if err != nil {
		return err
	}

	d.Time = parsed

	return nil
}

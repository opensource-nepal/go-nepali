// contains models/structs related to the package
package nepalitime

import (
	"time"
)

type NepaliTime struct {
	// note that these members represent nepali values
	year        int8
	month       int8
	day         int8
	hour        int8
	minute      int8
	second      int8
	millisecond int8
	nanosecond  int8

	englishTime time.Time
}

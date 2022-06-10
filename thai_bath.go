package bahttext

import "fmt"

var numToText = map[uint8]string{
	0: "ศูนย์",
	1: "หนึ่ง",
	2: "สอง",
	3: "สาม",
	4: "สี่",
	5: "ห้า",
	6: "หก",
	7: "เจ็ด",
	8: "แปด",
	9: "เก้า",
}

func THBText(v float64) string {
	text := ""
	if v == 0 {
		text = fmt.Sprintf("%sบาท", numToText[0])
	}
	return text
}

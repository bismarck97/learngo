package main

//练习2：写一个Ticket类,有一个距离属性,
//不能为负数,有一个价格属性 ,
//并且根据距离distance计算价格Price (1元/公里):
//0-100公里     票价不打折
//101-200公里    总额打9.5折
//201-300公里    总额打9折
//300公里以上    总额打8折

type Ticket struct {
	Distance float64
	Price    float64
}

func (t *Ticket) GetValue(distance, price float64) {
	if distance < 0 {
		distance = 0
	}
	t.Distance = distance
	t.Price = price
	if t.Distance > 0 && t.Distance <= 100 {
		t.Price = t.Price * 1.0
	} else if t.Distance > 101 && t.Distance <= 200 {
		t.Price = t.Price * 0.95
	} else if t.Distance > 201 && t.Distance <= 300 {
		t.Price = t.Price * 0.9
	} else {
		t.Price = t.Price * 0.8
	}
}
func main09() {
	var t Ticket
	t.GetValue(300, 200)
}

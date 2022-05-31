package visitor_module

import "math"

type Device interface {
	Accept(v visitor)
	GetType()
}

func GetDeviceType(d Device){
	d.GetType()
}

//Структура SmartLamp и ее методы
type SmartLamp struct{
	resistanceOne float64
	resistanceTwo float64
	current float64
}

func (s* SmartLamp)SetParams(r1,r2,c float64){
	s.resistanceOne=r1
	s.resistanceTwo=r2
	s.current=c
}

func (s *SmartLamp) Accept(v visitor){
	v.visitForLamp(s)
}

func (s *SmartLamp) GetType() string{
	return "Лампочка"
}

//Структура SmartTeapot и ее методы
type SmartTeapot struct{
	innerResistance float64
	outerResistance float64
	additionalResistance float64
	current float64
}

func (s* SmartTeapot)SetParams(ri,ro,ra,c float64){
	s.innerResistance=ri
	s.outerResistance=ro
	s.additionalResistance=ra
	s.current=c
}

func (s *SmartTeapot) Accept(v visitor){
	v.visitForTeapot(s)
}

func (s *SmartTeapot) GetType()string{
	return "Чайничек"
}

//интерфейс посетителя
type visitor interface {
	visitForLamp(*SmartLamp)
	visitForTeapot(*SmartTeapot)
}

//Методы, работающие через паттерн посетитель
type PowerCalculator struct {
	power float64
}
func (p *PowerCalculator)GetPower()float64{
	return p.power
}

func (p *PowerCalculator) visitForLamp(l *SmartLamp){
	p.power=l.current*math.Pow(l.resistanceOne+l.resistanceTwo,2)
}

func (p *PowerCalculator) visitForTeapot(t *SmartTeapot){
	p.power=t.current*math.Pow(t.innerResistance+t.outerResistance+t.additionalResistance,2)
}

type ResistanceCalculator struct {
	totalResistance float64
}

func (r *ResistanceCalculator) visitForLamp(l *SmartLamp){
	r.totalResistance=l.resistanceOne+l.resistanceTwo
}

func (r *ResistanceCalculator) visitForTeapot(l *SmartLamp){
	r.totalResistance=l.resistanceOne+l.resistanceTwo
}

func (r *ResistanceCalculator) GetTotal() float64{
	return r.totalResistance
}
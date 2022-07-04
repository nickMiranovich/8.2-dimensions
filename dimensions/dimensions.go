package dimensions

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

// размер + единица измерения
type Unit struct {
	Value float64
	T     UnitType
}

//конвертация inch\cm и обратно
func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		switch u.T {
		case "inch":
			value = value * 2.54
		case "cm":
			value = value / 2.54
		}
	}
	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type dimensions struct {
	length Unit
	width  Unit
	height Unit
}

func (d dimensions) Length() Unit {
	return d.length
}

func (d dimensions) Width() Unit {
	return d.width
}

func (d dimensions) Height() Unit {
	return d.height
}

//конструкторы
func NewDimensions(l, w, h Unit, u UnitType) dimensions {
	return dimensions{
		length: Unit{Value: l.Get(u), T: u},
		width:  Unit{Value: w.Get(u), T: u},
		height: Unit{Value: h.Get(u), T: u},
	}
}

func NewUnit(v float64, t UnitType) Unit {
	return Unit{
		Value: v,
		T:     t,
	}
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

//Auto реализация
type auto struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

func (d auto) Brand() string {
	return d.brand
}

func (d auto) Model() string {
	return d.model
}

func (d auto) MaxSpeed() int {
	return d.maxSpeed
}

func (d auto) EnginePower() int {
	return d.enginePower
}

func (d auto) Dimensions() Dimensions {
	return d.dimensions
}

//конструктор
func NewAuto(b, m string, d Dimensions, ms, ep int) auto {
	return auto{
		brand:       b,
		model:       m,
		dimensions:  d,
		maxSpeed:    ms,
		enginePower: ep,
	}
}

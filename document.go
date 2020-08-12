package segment

type Document interface {
	Analyze()
	EachField(vf VisitField)
}

type VisitField func(Field)

type Field interface {
	Name() string
	Length() int
	EachTerm(vt VisitTerm)
	Value() []byte

	Index() bool
	Store() bool
	IndexDocValues() bool
}

type VisitTerm func(FieldTerm)

type FieldTerm interface {
	Term() []byte
	Frequency() int
	EachLocation(vl VisitLocation)
}

type VisitLocation func(Location)

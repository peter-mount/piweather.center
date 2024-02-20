package units

type UnitsVisitor interface {
	Unit(*Unit) error
}

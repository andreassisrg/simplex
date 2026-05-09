package dto

type PLInput struct {
	NumVariables               uint64
	NumRestrictions            uint64
	VariableRestrictions       []VarRestriction
	TargetFunctionCoefficients []int64
	Restrictions               []Restriction
}

type VarRestriction int64

const (
	NonNegativeVar VarRestriction = 1
	NonPositiveVar VarRestriction = -1
	FreeVar        VarRestriction = 0
)

type Restriction struct {
	Coefficients      []int64
	RestrictionSign   Sign
	DependentVariable int64
}

type Sign string

const (
	LessThanOrEqualTo    Sign = "<="
	GreaterThanOrEqualTo Sign = ">="
	EqualTo              Sign = "=="
)

package filter

import (
	"github.com/bobtfish/sensuplugingo"
)

func filter(e sensuplugingo.Event) (bool, error) {
	r, err := FilterDisabled(e)
	if r == true || err != nil {
		return r, err
	}
	r, err = FilterRepeated(e)
        if r == true || err != nil {
                return r, err
        }
	r, err = FilterSilenced(e)
        if r == true || err != nil {
                return r, err
        }
	r, err = FilterDependencies(e)
         return r, err
}

func FilterDisabled(e sensuplugingo.Event) (bool, error) {
	return false, nil
}

func FilterRepeated(e sensuplugingo.Event) (bool, error) {
        return false, nil
}

func FilterSilenced(e sensuplugingo.Event) (bool, error) {
        return false, nil
}

func FilterDependencies(e sensuplugingo.Event) (bool, error) {
        return false, nil
}


package jd

type jsonNull struct{}

var _ JsonNode = jsonNull{}

func (n jsonNull) Json() string {
	return renderJson(nil)
}

func (n jsonNull) Equals(node JsonNode, metadata ...Metadata) bool {
	switch node.(type) {
	case jsonNull:
		return true
	default:
		return false
	}
}

func (n jsonNull) hashCode() [8]byte {
	return hash([]byte{0xFE, 0x73, 0xAB, 0xCC, 0xE6, 0x32, 0xE0, 0x88}) // random bytes
}

func (n jsonNull) Diff(node JsonNode, metadata ...Metadata) Diff {
	return n.diff(node, Path{}, metadata)
}

func (n jsonNull) diff(node JsonNode, path Path, metadata []Metadata) Diff {
	d := make(Diff, 0)
	if n.Equals(node) {
		return d
	}
	e := DiffElement{
		Path:      path.clone(),
		OldValues: nodeList(n),
		NewValues: nodeList(node),
	}
	return append(d, e)
}

func (n jsonNull) Patch(d Diff, metadata ...Metadata) (JsonNode, error) {
	return patchAll(n, d, metadata)
}

func (n jsonNull) patch(pathBehind, pathAhead Path, oldValues, newValues []JsonNode, metadata []Metadata) (JsonNode, error) {
	if len(pathAhead) != 0 {
		return patchErrExpectColl(n, pathAhead[0])
	}
	if len(oldValues) > 1 || len(newValues) > 1 {
		return patchErrNonSetDiff(oldValues, newValues, pathBehind)
	}
	oldValue := singleValue(oldValues)
	newValue := singleValue(newValues)
	if !n.Equals(oldValue) {
		return patchErrExpectValue(oldValue, n, pathBehind)
	}
	return newValue, nil
}

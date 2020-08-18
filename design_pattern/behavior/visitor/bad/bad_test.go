package bad

import "testing"

func TestVisitor(t *testing.T) {
	persons := make([]Person, 0)

	man1 := &Man{action: "成功"}
	persons = append(persons, man1)
	woman1 := &Woman{action: "成功"}
	persons = append(persons, woman1)
	man2 := &Man{action: "失败"}
	persons = append(persons, man2)
	woman2 := &Woman{action: "失败"}
	persons = append(persons, woman2)

	for _, p := range persons {
		p.GetConclusion()
	}
}

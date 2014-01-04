package graph

import (
	"testing"
)

func TestVertexId(t *testing.T) {
    v := NewBaseVertex();
    if v.Id() == nil {
        t.Errorf("Identifier on vertex was nil");
    }
}

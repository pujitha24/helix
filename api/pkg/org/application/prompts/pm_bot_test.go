package prompts_test

import (
	"context"
	"strings"
	"testing"

	"github.com/helixml/helix/api/pkg/org/application/helixevents"
	"github.com/helixml/helix/api/pkg/org/application/prompts"
)

func TestPMBotTemplateMatchesCurrentEventAndAccessModel(t *testing.T) {
	t.Parallel()
	msgs, err := prompts.PMBot{}.Render(context.Background(), nil)
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	text := msgs[0].Text
	for _, stale := range []string{"Spec tasks:", "<projectId>"} {
		if strings.Contains(text, stale) {
			t.Errorf("template still references removed per-project trigger %q", stale)
		}
	}
	for _, want := range []string{helixevents.TriggerID, "Project access"} {
		if !strings.Contains(text, want) {
			t.Errorf("template missing %q", want)
		}
	}
}

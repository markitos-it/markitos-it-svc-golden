package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestGolden_ShouldKeepAssignedDomainState(t *testing.T) {
	now := time.Date(2026, 2, 22, 10, 30, 0, 0, time.UTC)
	prefix := HelperRandomAlphaPrefix(t, 8)
	expectedID := prefix + "-golden-001"
	expectedTitle := prefix + "-Servicio de Documentos"
	expectedDescription := prefix + "-Plantilla base"
	expectedCategory := prefix + "-backend"
	expectedContentB64 := prefix + "-Y29udGVudA=="
	expectedCoverImage := "https://example.com/" + prefix + "/cover.png"
	expectedTags := []string{prefix + "-go", prefix + "-grpc", prefix + "-postgres"}

	golden := Golden{
		ID:          expectedID,
		Title:       expectedTitle,
		Description: expectedDescription,
		Category:    expectedCategory,
		Tags:        expectedTags,
		UpdatedAt:   now,
		ContentB64:  expectedContentB64,
		CoverImage:  expectedCoverImage,
	}

	if golden.ID != expectedID {
		t.Fatalf("expected ID %s, got %s", expectedID, golden.ID)
	}
	if golden.Title != expectedTitle {
		t.Fatalf("expected Title %s, got %s", expectedTitle, golden.Title)
	}
	if golden.Description != expectedDescription {
		t.Fatalf("expected Description %s, got %s", expectedDescription, golden.Description)
	}
	if golden.Category != expectedCategory {
		t.Fatalf("expected Category %s, got %s", expectedCategory, golden.Category)
	}
	if !reflect.DeepEqual(golden.Tags, expectedTags) {
		t.Fatalf("expected Tags %v, got %v", expectedTags, golden.Tags)
	}
	if !golden.UpdatedAt.Equal(now) {
		t.Fatalf("expected UpdatedAt %v, got %v", now, golden.UpdatedAt)
	}
	if golden.ContentB64 != expectedContentB64 {
		t.Fatalf("expected ContentB64 %s, got %s", expectedContentB64, golden.ContentB64)
	}
	if golden.CoverImage != expectedCoverImage {
		t.Fatalf("expected CoverImage %s, got %s", expectedCoverImage, golden.CoverImage)
	}
}

func TestGolden_ShouldExposeZeroValueAsEmptyDomainState(t *testing.T) {
	var golden Golden

	if golden.ID != "" {
		t.Fatalf("expected empty ID, got %s", golden.ID)
	}
	if golden.Title != "" {
		t.Fatalf("expected empty Title, got %s", golden.Title)
	}
	if golden.Description != "" {
		t.Fatalf("expected empty Description, got %s", golden.Description)
	}
	if golden.Category != "" {
		t.Fatalf("expected empty Category, got %s", golden.Category)
	}
	if len(golden.Tags) != 0 {
		t.Fatalf("expected empty Tags, got %v", golden.Tags)
	}
	if !golden.UpdatedAt.IsZero() {
		t.Fatalf("expected zero UpdatedAt, got %v", golden.UpdatedAt)
	}
	if golden.ContentB64 != "" {
		t.Fatalf("expected empty ContentB64, got %s", golden.ContentB64)
	}
	if golden.CoverImage != "" {
		t.Fatalf("expected empty CoverImage, got %s", golden.CoverImage)
	}
}

package inventory

import (
	"errors"
	"testing"
	"time"
)

func TestBug2_StockInOverflowRejected(t *testing.T) {
	s := New()
	now := time.Date(2026, 8, 14, 12, 0, 0, 0, time.UTC)
	maxStock := int64(9223372036854775800)
	if _, err := s.Create(CreateInput{SKU: "OVF", Name: "overflow", Stock: maxStock}, now); err != nil {
		t.Fatal(err)
	}
	_, err := s.StockIn("OVF", AmountInput{Amount: 100}, now)
	if err == nil {
		t.Fatal("expected error for overflow stock-in")
	}
	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected ErrInvalidAmount, got %v", err)
	}
	p, _ := s.Get("OVF")
	if p.Stock != maxStock {
		t.Fatalf("stock should remain %d, got %d", maxStock, p.Stock)
	}
}

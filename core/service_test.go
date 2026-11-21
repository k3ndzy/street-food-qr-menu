package qrmenu

import "testing"

func TestQRMenuService_OrderFromTable(t *testing.T) {
	svc := NewQRMenuService()
	ticket, err := svc.OrderFromTable("BAN-05", []int{1, 3})
	if err != nil {
		t.Fatalf("lỗi gọi món: %v", err)
	}

	if ticket.TableCode != "BAN-05" || len(ticket.ItemNames) != 2 {
		t.Errorf("phiếu bếp không đúng: %+v", ticket)
	}
}

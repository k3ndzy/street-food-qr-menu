package qrmenu

import (
	"errors"
	"sync"
	"time"
)

type QRMenuService struct {
	mu      sync.Mutex
	items   map[int]*QRItem
	tickets map[int]*KitchenTicket
	lastID  int
}

func NewQRMenuService() *QRMenuService {
	s := &QRMenuService{
		items:   make(map[int]*QRItem),
		tickets: make(map[int]*KitchenTicket),
	}
	s.items[1] = &QRItem{ID: 1, Name: "Bánh Tráng Trộn Đặc Biệt", Price: 25000}
	s.items[2] = &QRItem{ID: 2, Name: "Trứng Cút Lộn Xào Me", Price: 35000}
	s.items[3] = &QRItem{ID: 3, Name: "Trà Tắc Khổng Lồ", Price: 15000}
	return s
}

func (s *QRMenuService) OrderFromTable(tableCode string, itemIDs []int) (*KitchenTicket, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var names []string
	for _, id := range itemIDs {
		it, ok := s.items[id]
		if !ok || it.IsSoldOut {
			return nil, errors.New("món ăn không khả dụng")
		}
		names = append(names, it.Name)
	}

	s.lastID++
	t := &KitchenTicket{
		TicketID:  s.lastID,
		TableCode: tableCode,
		ItemNames: names,
		OrderedAt: time.Now(),
		Status:    "COOKING",
	}
	s.tickets[s.lastID] = t
	return t, nil
}

package linkedlist

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateList(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want interface{}
	}{
		{name: "Int", args: 55, want: 55},
		{name: "Nil", args: nil, want: nil},
		{name: "String", args: "ABC", want: "ABC"},
		{name: "Struct", args: struct{}{}, want: struct{}{}},
		{name: "Array", args: [2]int{1, 2}, want: [2]int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateList(tt.args)
			diff := cmp.Diff(got.head.value, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestLinkedList_Error(t *testing.T) {

	tests := []struct {
		name string
		lle  *LinkedListError
		want string
	}{
		{name: "Remove lone head value", lle: &LinkedListError{"Invalid index"}, want: "Invalid index"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lle.Error(); got != tt.want {
				t.Errorf("LinkedList.RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNode(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want interface{}
	}{
		{name: "Int", args: 55, want: 55},
		{name: "Nil", args: nil, want: nil},
		{name: "String", args: "ABC", want: "ABC"},
		{name: "Struct", args: struct{}{}, want: struct{}{}},
		{name: "Array", args: [2]int{1, 2}, want: [2]int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateNode(tt.args)
			diff := cmp.Diff(got.value, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	tests := []struct {
		name string
		ll   *LinkedList
		args interface{}
		want interface{}
	}{
		{name: "Add string", ll: CreateList(55), args: "77", want: "77"},
		{name: "Add int", ll: CreateList(55), args: 77, want: 77},
		{name: "Add struct", ll: CreateList(55), args: struct{}{}, want: struct{}{}},
		{name: "Add Nil", ll: CreateList(55), args: nil, want: 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.Add(tt.args)
			diff := cmp.Diff(tt.ll.tail.value, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)
	ll1.Add(99)
	ll1.Add(33)

	tests := []struct {
		name    string
		ll      *LinkedList
		args    int
		want    interface{}
		wantErr bool
	}{
		{name: "Get head value", ll: CreateList(55), args: 0, want: 55, wantErr: false},
		{name: "Get valid value", ll: ll1, args: 3, want: 33, wantErr: false},
		{name: "Get invalid index", ll: CreateList(55), args: 1, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ll.Get(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestLinkedList_RemoveIndex(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)
	ll1.Add(99)
	ll1.Add(33)

	tests := []struct {
		name    string
		ll      *LinkedList
		args    int
		want    bool
		wantErr bool
	}{
		{name: "Remove head value", ll: CreateList(55), args: 0, want: false, wantErr: true},
		{name: "Remove valid index", ll: ll1, args: 3, want: true, wantErr: false},
		{name: "Remove invalid index", ll: ll1, args: 33, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ll.RemoveIndex(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.RemoveIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LinkedList.RemoveIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveElement(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)
	ll1.Add(99)
	ll1.Add(33)

	tests := []struct {
		name string
		ll   *LinkedList
		args int
		want bool
	}{
		{name: "Remove lone head value", ll: CreateList(55), args: 55, want: false},
		{name: "Remove valid value", ll: ll1, args: 99, want: true},
		{name: "Remove valid value", ll: ll1, args: 33, want: true},
		{name: "Remove invalid value", ll: ll1, args: 88, want: false},
		{name: "Remove valid head value", ll: ll1, args: 55, want: true},
		{name: "Remove lone head value", ll: ll1, args: 77, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.RemoveElement(tt.args); got != tt.want {
				t.Errorf("LinkedList.RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)

	tests := []struct {
		name string
		ll   *LinkedList
		args int
		want bool
	}{
		{name: "Contains head", ll: CreateList(55), args: 55, want: true},
		{name: "Contains valid value", ll: ll1, args: 55, want: true},
		{name: "Contains invalid value", ll: ll1, args: 88, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Contains(tt.args); got != tt.want {
				t.Errorf("LinkedList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)

	tests := []struct {
		name string
		ll   *LinkedList
		want int
	}{
		{name: "Size of 1", ll: CreateList(55), want: 1},
		{name: "Size of 2", ll: ll1, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Size(); got != tt.want {
				t.Errorf("LinkedList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Reverse()
	diff := cmp.Diff(ll1.String(), "List -> 55")
	if diff != "" {
		t.Fatalf(diff)
	}

	ll1.Add(77)

	ll2 := CreateList(77)
	ll2.Add(55)

	ll1.Reverse()
	diff = cmp.Diff(ll1.String(), ll2.String())
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestLinkedList_ReverseRecurse(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)

	ll2 := CreateList(77)
	ll2.Add(55)

	ll1.ReverseRecurse()
	diff := cmp.Diff(ll1.String(), ll2.String())
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestLinkedList_String(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)

	ll1.Reverse()
	diff := cmp.Diff(ll1.String(), "List -> 77 -> 55")
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestLinkedList_HasCycle(t *testing.T) {
	ll1 := CreateList(55)
	if got := ll1.HasCycle(); !reflect.DeepEqual(got, false) {
		t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, false)
	}

	ll1.Add(77)
	ll1.Add(99)
	ll1.Add(33)

	if got := ll1.HasCycle(); !reflect.DeepEqual(got, false) {
		t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, false)
	}

	newNode := CreateNode(22)
	ll1.tail.next, ll1.tail = newNode, newNode
	newNode.next = ll1.head.next.next

	if got := ll1.HasCycle(); !reflect.DeepEqual(got, true) {
		t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, true)
	}
}

func TestLinkedList_MiddleNode(t *testing.T) {
	ll1 := CreateList(55)
	ll1.Add(77)
	ll1.Add(99)

	ll2 := CreateList(77)
	ll2.Add(55)

	tests := []struct {
		name string
		ll   *LinkedList
		want interface{}
	}{
		{name: "Middle value", ll: CreateList(55), want: 55},
		{name: "Middle value", ll: ll1, want: 77},
		{name: "Middle value", ll: ll2, want: 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.MiddleNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

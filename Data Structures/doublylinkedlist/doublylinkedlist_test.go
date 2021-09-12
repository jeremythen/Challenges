package doublylinkedlist

import (
	"fmt"
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

func TestDoublyLinkedListError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *DoublyLinkedListError
		want string
	}{
		{name: "Remove lone head value", e: &DoublyLinkedListError{"Invalid index"}, want: "Invalid index"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("DoublyLinkedListError.Error() = %v, want %v", got, tt.want)
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

func TestDoublyLinkedList_Add(t *testing.T) {
	tests := []struct {
		name string
		dll  *DoublyLinkedList
		args interface{}
		want interface{}
	}{
		{name: "Add string", dll: CreateList(55), args: "77", want: "77"},
		{name: "Add int", dll: CreateList(55), args: 77, want: 77},
		{name: "Add struct", dll: CreateList(55), args: struct{}{}, want: struct{}{}},
		{name: "Add Nil", dll: CreateList(55), args: nil, want: 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dll.Add(tt.args)
			diff := cmp.Diff(tt.dll.tail.value, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestDoublyLinkedList_GetIndex(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Add(77)
	dll1.Add(99)
	dll1.Add(33)

	tests := []struct {
		name    string
		dll     *DoublyLinkedList
		args    int
		want    interface{}
		wantErr bool
	}{
		{name: "Get head value", dll: CreateList(55), args: 0, want: 55, wantErr: false},
		{name: "Get valid value", dll: dll1, args: 3, want: 33, wantErr: false},
		{name: "Get invalid index", dll: CreateList(55), args: 1, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dll.GetIndex(tt.args)
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

func TestDoublyLinkedList_RemoveIndex(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Add(77)
	dll1.Add(99)
	dll1.Add(33)

	tests := []struct {
		name    string
		dll     *DoublyLinkedList
		args    int
		want    bool
		wantErr bool
	}{
		{name: "Remove head value", dll: CreateList(55), args: 0, want: false, wantErr: true},
		{name: "Remove valid index", dll: dll1, args: 3, want: true, wantErr: false},
		{name: "Remove invalid index", dll: dll1, args: 33, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dll.RemoveIndex(tt.args)
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

func TestDoublyLinkedList_RemoveElement(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Add(77)
	dll1.Add(99)
	dll1.Add(33)

	tests := []struct {
		name string
		dll  *DoublyLinkedList
		args int
		want bool
	}{
		{name: "Remove lone head value", dll: CreateList(55), args: 55, want: false},
		{name: "Remove valid value", dll: dll1, args: 99, want: true},
		{name: "Remove valid value", dll: dll1, args: 33, want: true},
		{name: "Remove invalid value", dll: dll1, args: 88, want: false},
		{name: "Remove valid head value", dll: dll1, args: 55, want: true},
		{name: "Remove lone head value", dll: dll1, args: 77, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.dll)
			if got := tt.dll.RemoveElement(tt.args); got != tt.want {
				t.Errorf("LinkedList.RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Add(77)
	dll1.Add(99)
	dll1.Add(33)

	tests := []struct {
		name string
		dll  *DoublyLinkedList
		args int
		want bool
	}{
		{name: "Contains head", dll: CreateList(55), args: 55, want: true},
		{name: "Contains valid value", dll: dll1, args: 99, want: true},
		{name: "Contains invalid value", dll: dll1, args: 88, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dll.Contains(tt.args); got != tt.want {
				t.Errorf("LinkedList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_Size(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Add(77)

	tests := []struct {
		name string
		dll  *DoublyLinkedList
		want int
	}{
		{name: "Size of 1", dll: CreateList(55), want: 1},
		{name: "Size of 2", dll: dll1, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dll.Size(); got != tt.want {
				t.Errorf("LinkedList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_Reverse(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Reverse()
	diff := cmp.Diff(dll1.String(), "List <-> 55")
	if diff != "" {
		t.Fatalf(diff)
	}

	dll1.Add(77)

	dll2 := CreateList(77)
	dll2.Add(55)

	dll1.Reverse()
	diff = cmp.Diff(dll1.String(), dll2.String())
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestDoublyLinkedList_ReverseRecurse(t *testing.T) {
	dll1 := CreateList(55)
	dll1.ReverseRecurse()
	diff := cmp.Diff(dll1.String(), "List <-> 55")
	if diff != "" {
		t.Fatalf(diff)
	}

	dll1.Add(77)

	dll2 := CreateList(77)
	dll2.Add(55)

	dll1.ReverseRecurse()
	diff = cmp.Diff(dll1.String(), dll2.String())
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestDoublyLinkedList_String(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Add(77)

	dll1.Reverse()
	diff := cmp.Diff(dll1.String(), "List <-> 77 <-> 55")
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestDoublyLinkedList_HasCycle(t *testing.T) {
	dll1 := CreateList(55)
	if got := dll1.HasCycle(); !reflect.DeepEqual(got, false) {
		t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, false)
	}

	dll1.Add(77)
	dll1.Add(99)
	dll1.Add(33)

	if got := dll1.HasCycle(); !reflect.DeepEqual(got, false) {
		t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, false)
	}

	newNode := CreateNode(22)
	newNode.previous, dll1.tail.next, dll1.tail = dll1.tail, newNode, newNode
	newNode.next = dll1.head.next.next

	if got := dll1.HasCycle(); !reflect.DeepEqual(got, true) {
		t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, true)
	}
}

func TestDoublyLinkedList_MiddleNode(t *testing.T) {
	dll1 := CreateList(55)
	dll1.Add(77)
	dll1.Add(99)

	dll2 := CreateList(77)
	dll2.Add(55)

	tests := []struct {
		name string
		dll  *DoublyLinkedList
		want interface{}
	}{
		{name: "Middle value", dll: CreateList(55), want: 55},
		{name: "Middle value", dll: dll1, want: 77},
		{name: "Middle value", dll: dll2, want: 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dll.MiddleNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_AddFirst(t *testing.T) {
	tests := []struct {
		name string
		dll  *DoublyLinkedList
		args interface{}
		want interface{}
	}{
		{name: "Add string", dll: CreateList(55), args: "77", want: "77"},
		{name: "Add int", dll: CreateList(55), args: 77, want: 77},
		{name: "Add struct", dll: CreateList(55), args: struct{}{}, want: struct{}{}},
		{name: "Add Nil", dll: CreateList(55), args: nil, want: 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dll.AddFirst(tt.args)
			diff := cmp.Diff(tt.dll.head.value, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestDoublyLinkedList_GetFirst(t *testing.T) {
	dll1 := CreateList(33)
	dll1.Add(77)
	dll1.Add(99)

	tests := []struct {
		name string
		dll  *DoublyLinkedList
		want interface{}
	}{
		{name: "Get first", dll: CreateList(55), want: 55},
		{name: "Get first", dll: dll1, want: 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dll.GetFirst()
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestDoublyLinkedList_GetLast(t *testing.T) {
	dll1 := CreateList(33)
	dll1.Add(77)
	dll1.Add(99)

	tests := []struct {
		name string
		dll  *DoublyLinkedList
		want interface{}
	}{
		{name: "Get last", dll: CreateList(55), want: 55},
		{name: "Get last", dll: dll1, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dll.GetLast()
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

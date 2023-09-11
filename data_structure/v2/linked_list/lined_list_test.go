package linkedlist

import (
	"reflect"
	"testing"
)

func TestCreateLinkerList(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLinkerList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLinkerList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	type fields struct {
		node *Node[any]
	}
	type args struct {
		data any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				node: tt.fields.node,
			}
			if err := l.Add(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	type fields struct {
		node *Node[any]
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				node: tt.fields.node,
			}
			if err := l.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Edit(t *testing.T) {
	type fields struct {
		node *Node[any]
	}
	type args struct {
		key     string
		newData any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				node: tt.fields.node,
			}
			if err := l.Edit(tt.args.key, tt.args.newData); (err != nil) != tt.wantErr {
				t.Errorf("Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Find(t *testing.T) {
	type fields struct {
		node *Node[any]
	}
	tests := []struct {
		name         string
		fields       fields
		wantNodeList []*Node[any]
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				node: tt.fields.node,
			}
			gotNodeList, err := l.Find()
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNodeList, tt.wantNodeList) {
				t.Errorf("Find() gotNodeList = %v, want %v", gotNodeList, tt.wantNodeList)
			}
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	type fields struct {
		node *Node[any]
	}
	type args struct {
		data any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantKey string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				node: tt.fields.node,
			}
			gotKey, err := l.Get(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("Get() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

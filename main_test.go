package main

import (
	"reflect"
	"testing"
)

func TestBT(t *testing.T) {
	cases := []struct {
		operations []struct {
			method string
			data   int
		}
		expected []any
	}{
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 3},
				{"insert", 7},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 5},
				{"min", 0},
				{"max", 0},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				7,
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 3},
				{"insert", 7},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 5},
				{"min", 0},
				{"max", 0},
				{"delete", 1},
				{"delete", 7},
				{"min", 2},
				{"max", 6},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				7,
				nil,
				nil,
				2,
				6,
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 3},
				{"insert", 5},
				{"insert", 7},
				{"min", 0},
				{"max", 0},
				{"preOrder", 0},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				7,
				[]int{4, 2, 1, 3, 6, 5, 7},
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 3},
				{"insert", 5},
				{"insert", 7},
				{"min", 0},
				{"max", 0},
				{"postOrder", 0},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				7,
				[]int{1, 3, 2, 5, 7, 6, 4},
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 3},
				{"insert", 5},
				{"insert", 7},
				{"min", 0},
				{"max", 0},
				{"inOrder", 0},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				7,
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 3},
				{"insert", 5},
				{"insert", 7},
				{"min", 0},
				{"max", 0},
				{"exist", 3},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				7,
				true,
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 3},
				{"insert", 5},
				{"insert", 7},
				{"min", 0},
				{"max", 0},
				{"exist", 10},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				7,
				false,
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"insert", 4},
				{"insert", 2},
				{"insert", 6},
				{"insert", 1},
				{"insert", 3},
				{"min", 0},
				{"max", 0},
				{"height", 0},
			},
			expected: []any{
				nil,
				nil,
				nil,
				nil,
				nil,
				1,
				6,
				2,
			},
		},
		{
			operations: []struct {
				method string
				data   int
			}{
				{"height", 0},
			},
			expected: []any{
				0,
			},
		},
	}

	for _, c := range cases {
		bt := bTNode{}

		for i, o := range c.operations {
			switch o.method {
			case "insert":
				bt.insert(o.data)
			case "delete":
				bt = *bt.delete(o.data)
			case "preOrder":
				got, expected := bt.preOrder(), c.expected[i]
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("IDX: %v, got: %v != expected: %v", i, got, expected)
				}
			case "postOrder":
				got, expected := bt.postOrder(), c.expected[i]
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("IDX: %v, got: %v != expected: %v", i, got, expected)
				}
			case "inOrder":
				got, expected := bt.inOrder(), c.expected[i]
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("IDX: %v, got: %v != expected: %v", i, got, expected)
				}
			case "exist":
				got, expected := bt.exist(o.data), c.expected[i]
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("IDX: %v, got: %v != expected: %v", i, got, expected)
				}
			case "height":
				got, expected := bt.height(), c.expected[i]
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("IDX: %v, got: %v != expected: %v", i, got, expected)
				}
			case "min":
				got, expected := bt.min(), c.expected[i]
				if got != expected {
					t.Errorf("IDX: %v, got: %v != expected: %v", i, got, expected)
				}
			case "max":
				got, expected := bt.max(), c.expected[i]
				if got != expected {
					t.Errorf("IDX: %v, got: %v != expected: %v", i, got, expected)
				}
			}
		}
	}
}

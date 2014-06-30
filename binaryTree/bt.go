/*
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) Ivan Anfilatov aka t0pep0 (t0pep0.gentoo@gmail.com), 2014
 */

package binaryTree

//Binary tree node struct
type BinaryTree struct {
	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
	Index  string
	Value  interface{}
}

//Add node by index to binary tree
func (bt *BinaryTree) Set(index string, value interface{}) {
	if bt == nil {
		bt.Index = index
		bt.Value = value
		return
	}
	if bt.Index == "" {
		bt.Index = index
		bt.Value = value
		return
	}
	if bt.Index == index {
		bt.Value = value
		return
	}
	if bt.Index > index {
		if bt.Left == nil {
			bt.Left = new(BinaryTree)
			bt.Left.Parent = bt
		}
		bt.Left.Set(index, value)
		return
	}
	if bt.Index < index {
		if bt.Right == nil {
			bt.Right = new(BinaryTree)
			bt.Right.Parent = bt
		}
		bt.Right.Set(index, value)
		return
	}
}

//move node
func (bt *BinaryTree) move(node *BinaryTree) {
	if node == nil {
		return
	}
	if (bt == nil) || (bt.Index == "") {
		*bt = *node
		return
	}
	if bt.Index > node.Index {
		if bt.Left == nil {
			bt.Left.Parent = bt
		}
		bt.Left.move(node)
		return
	}
	if bt.Index < node.Index {
		if bt.Right == nil {
			bt.Right.Parent = bt
		}
		bt.Right.move(node)
		return
	}
}

// Get value by node index
func (bt *BinaryTree) Get(index string) (value interface{}, found bool) {
	if bt == nil {
		return value, false
	}
	if bt.Index == index {
		return bt.Value, true
	}
	if bt.Index > index {
		return bt.Left.Get(index)
	}
	if bt.Index < index {
		return bt.Right.Get(index)
	}
	return
}

//Delete node from tree by index
func (bt *BinaryTree) Delete(index string) {
	if bt == nil {
		return
	}
	if bt.Index == index {
		if bt.Parent == nil {
			newNode := new(BinaryTree)
			newNode.move(bt.Right)
			newNode.move(bt.Left)
			*bt = *newNode
			return
		}
		if bt.Parent != nil {
			if bt.Parent.Left == bt {
				bt.Parent.Left = nil
			} else {
				bt.Parent.Right = nil
			}
			bt.Parent.move(bt.Right)
			bt.Parent.move(bt.Left)
			bt = new(BinaryTree)
		}
	}
	if bt.Index > index {
		bt.Left.Delete(index)
		return
	}
	if bt.Index < index {
		bt.Right.Delete(index)
		return
	}
}

//Get curent tree length
func (bt *BinaryTree) Length() (size int64) {
	if bt != nil {
		if bt.Index != "" {
			size = bt.Left.Length()
			size++
			size += bt.Right.Length()
		}
	}
	return size
}

//Range tree
func (bt *BinaryTree) Range(rangeFunc func(*BinaryTree)) {
	if bt != nil {
		bt.Left.Range(rangeFunc)
		rangeFunc(bt)
		bt.Right.Range(rangeFunc)
	}
}

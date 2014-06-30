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

package main

import (
	. "./binaryTree"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"text/tabwriter"
	"time"
)

func timing(timFunc func()) (nanoSec int64) {
	timerStart := time.Now().UnixNano()
	timFunc()
	nanoSec = time.Now().UnixNano() - timerStart
	return nanoSec
}

func percent(btTime, mpTime int64) (percent int) {
	if btTime > mpTime {
		percent = 100 - int((float64(mpTime)/float64(btTime))*100.0)
	} else {
		percent = 100 - int((float64(btTime)/float64(mpTime))*100.0)
	}
	return percent
}

func winner(btTime, mpTime int64) (winner string) {
	if btTime > mpTime {
		winner = "HashMap"
	}
	if mpTime > btTime {
		winner = "BinaryTree"
	}
	if mpTime == btTime {
		winner = "none"
	}
	return winner
}

func chars(i int) (c string) {
	h := md5.New()
	iStr := strconv.Itoa(i)
	io.WriteString(h, iStr)
	c = fmt.Sprintf("%x", h.Sum(nil))
	return c
}

func main() {
	for i := 1; i < 1000000; i *= 10 {
		cicle(i)
		runtime.GC()
	}

}

func cicle(LOOP_COUNT int) {
	bt := new(BinaryTree)
	mp := make(map[string]interface{})

	btFillTime := timing(func() {
		i := 0
		for i < LOOP_COUNT {
			bt.Set(chars(i), i)
			i++
		}
	})

	btRangeTime := timing(func() {
		bt.Range(func(node *BinaryTree) {
			node.Value = node.Value
		})
	})

	btGetTime := timing(func() {
		for i := LOOP_COUNT; i > 0; i-- {
			_, _ = bt.Get(chars(i))
		}
	})

	btDeleteTime := timing(func() {
		for i := LOOP_COUNT; i > 0; i-- {
			bt.Delete(chars(i))
		}
	})

	mapFillTime := timing(func() {
		for i := 0; i < LOOP_COUNT; i++ {
			mp[chars(i)] = i
		}
	})

	mapRangeTime := timing(func() {
		for index, value := range mp {
			mp[index] = value
		}
	})

	mapGetTime := timing(func() {
		for i := LOOP_COUNT; i > 0; i-- {
			_, _ = mp[chars(i)]
		}
	})

	mapDeleteTime := timing(func() {
		for i := LOOP_COUNT; i > 0; i-- {
			delete(mp, chars(i))
		}
	})
	fmt.Println("Element count:", LOOP_COUNT)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Name\t| BinaryTree\t| Map\t| Percent\t| Winner")
	fmt.Fprintln(w, "Fill\t| ", btFillTime, "\t| ", mapFillTime, "\t| ", percent(btFillTime, mapFillTime), "\t| ", winner(btFillTime, mapFillTime))
	fmt.Fprintln(w, "Range\t| ", btRangeTime, "\t| ", mapRangeTime, "\t| ", percent(btRangeTime, mapRangeTime), "\t| ", winner(btRangeTime, mapRangeTime))
	fmt.Fprintln(w, "Get\t| ", btGetTime, "\t| ", mapGetTime, "\t| ", percent(btGetTime, mapGetTime), "\t| ", winner(btGetTime, mapGetTime))
	fmt.Fprintln(w, "Delete\t| ", btDeleteTime, "\t| ", mapDeleteTime, "\t| ", percent(btDeleteTime, mapDeleteTime), "\t| ", winner(btDeleteTime, mapDeleteTime))
	fmt.Fprintln(w)
	w.Flush()
}

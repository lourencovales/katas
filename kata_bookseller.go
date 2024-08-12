/**
A bookseller has lots of books classified in 26 categories labeled A, B, ... Z. Each book has a code c of 3, 4, 5 or more characters. The 1st character of a code is a capital letter which defines the book category.

In the bookseller's stocklist each code c is followed by a space and by a positive integer n (int n >= 0) which indicates the quantity of books of this code in stock.

For example an extract of a stocklist could be:

L = {"ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"}.
or
L = ["ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"] or ....

You will be given a stocklist (e.g. : L) and a list of categories in capital letters e.g :

M = {"A", "B", "C", "W"}
or
M = ["A", "B", "C", "W"] or ...

and your task is to find all the books of L with codes belonging to each category of M and to sum their quantity according to each category.

For the lists L and M of example you have to return the string (in Haskell/Clojure/Racket/Prolog a list of pairs):

(A : 20) - (B : 114) - (C : 50) - (W : 0)

where A, B, C, W are the categories, 20 is the sum of the unique book of category A, 114 the sum corresponding to "BKWRK" and "BTSQZ", 50 corresponding to "CDXEF" and 0 to category 'W' since there are no code beginning with W.
**/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Example variables

var k = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
var l = []string{"A", "B", "C", "D"}

func main() {
	a := StockList(k, l)
	fmt.Printf("%s", a)
}

// We take the two variables, and run the extractValues function, formatting the output to the required form

func StockList(listArt []string, listCat []string) string {
	var x []string
	for j, a := range extractValues(listCat, listArt) {
		o := "(" + listCat[j] + " : " + a + ")"
		x = append(x, o)
	}
	z := strings.Join(x, " - ")
	return z
}

// We take the same variables, pass the list of articles to the extractList function so we can get an output we can work on with the objective of adding the stocks of the different categories

func extractValues(listCat, listArt []string) []string {

	var q []string
	var sum int
	for _, s := range listCat {
		var p []string
		for _, i := range extractList(listArt, s) {
			if strings.HasPrefix(i, s) {
				sum = 0
				_, i, _ = strings.Cut(i, " ")
				p = append(p, i)
				for _, w := range p {
					f, _ := strconv.Atoi(w)
					sum += f
				}
			}
		}
		q = append(q, strconv.Itoa(sum))
	}
	return q
}

// This function serves to create a slice of strings that has the required information, i.e., the category and the stock of each entry in the list of articles

func extractList(listArt []string, listCat string) []string {
	var c []string
	for _, b := range listArt {
		if strings.HasPrefix(b, listCat) {
			_, b, _ := strings.Cut(b, " ")
			a := listCat + " " + b
			c = append(c, a)
		} else {
			a := listCat + " 0"
			c = append(c, a)
		}
	}
	return c
}

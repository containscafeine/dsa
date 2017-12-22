//The Stock Span Problem
//The stock span problem is a financial problem where we have a series of n daily price quotes for a stock and we need to calculate span of stock’s price for all n days.
//The span Si of the stock’s price on a given day i is defined as the maximum number of consecutive days just before the given day, for which the price of the stock on the current day is less than or equal to its price on the given day.
//For example, if an array of 7 days prices is given as {100, 80, 60, 70, 60, 75, 85}, then the span values for corresponding 7 days are {1, 1, 1, 2, 1, 4, 6}

package main

import "fmt"

func main() {
    stocks := []int{100, 80, 60, 70, 60, 75, 85}
    span := 0
    currentPos := 1
    for i, stock := range stocks {
        if stocks[currentPos] <= stock {
            fmt.Println("The current stock is ", stock)
            fmt.Println("The current position is ", i)
            span++
            fmt.Println("The span is ", span)
        }
        currentPos++
    }

}
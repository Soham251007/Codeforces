package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func solve() {
	var n, q int
	fmt.Fscan(reader, &n, &q)

	var s string
	fmt.Fscan(reader, &s)

	// Precompute indices of 'B'
	bIndices := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == 'B' {
			bIndices = append(bIndices, i)
		}
	}

	// Process queries
	for i := 0; i < q; i++ {
		var val int
		fmt.Fscan(reader, &val)

		// Case 1: If there are no 'B's, value decreases by 1 every step.
		// Answer is simply the value itself.
		if len(bIndices) == 0 {
			fmt.Fprintln(writer, val)
			continue
		}

		// Case 2: Jump between 'B's
		ans := 0
		currentPos := 0 // Position in string s
		bPtr := 0       // Pointer to the current target in bIndices

		// Since we always start at index 0, the first B we look for is bIndices[0]
		// We don't need to search for the starting B pointer.

		for val > 0 {
			// Get location of the next B
			targetB := bIndices[bPtr]

			// Calculate distance (number of A's) to the next B
			dist := 0
			if targetB >= currentPos {
				dist = targetB - currentPos
			} else {
				// We have to wrap around the string
				dist = (n - currentPos) + targetB
			}

			if val <= dist {
				// We run out of value just walking through the A's
				ans += val
				val = 0
			} else {
				// We survive the A's and reach the B
				ans += dist      // Walked past 'dist' A's
				val -= dist      // Decreased value by 'dist'
				
				// Now process the B
				ans++            // Takes 1 second to process B
				val = val / 2    // Value halves
				
				// Update position to the character immediately after the B
				currentPos = targetB + 1
				if currentPos == n {
					currentPos = 0
				}
				
				// Move the B pointer cyclically
				bPtr++
				if bPtr == len(bIndices) {
					bPtr = 0
				}
			}
		}
		fmt.Fprintln(writer, ans)
	}
}

func main() {
	defer writer.Flush()
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
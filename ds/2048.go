package main

import (
	"fmt"
	"math/rand"
)

type Direction int8

const N = 4
const (
	Up Direction = iota
	Down
	Left
	Rigth
)

type Grid [N][N]int

func rot(g Grid, numRot int) Grid {
	if numRot == 0 {
		return g
	}
	gR := g
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			gR[j][N-i-1] = g[i][j]
		}
	}
	return rot(gR, numRot-1)
}

func rotInv(g Grid, numRot int) Grid {
	return rot(g, N-numRot)
}

func collapseRow(row [N]int) [N]int {
	cur := 0
	for j := 0; j < N; j++ {
		if j != cur && row[j] != 0 {
			if row[cur] == 0 || row[cur] == row[j] {
				row[cur] += row[j]
				row[j] = 0
				cur++
			} else {
				cur++
				if cur >= j {
					continue
				}
				row[cur] = row[j]
				row[j] = 0
			}
		}
	}
	return row
}

func collapse(d Direction, g Grid) (gC Grid) {
	rotForDir := make(map[Direction]int)
	rotForDir[Left] = 0
	rotForDir[Up] = 3
	rotForDir[Rigth] = 2
	rotForDir[Down] = 1

	gC = rot(g, rotForDir[d])
	defer func() { gC = rotInv(gC, rotForDir[d]) }()

	for i := 0; i < N; i++ {
		gC[i] = collapseRow(gC[i])
	}
	return gC
}

func randomInsert(g Grid, r *rand.Rand) Grid {
	for _, i := range r.Perm(N) {
		for _, j := range r.Perm(N) {
			if g[i][j] == 0 {
				g[i][j] = 2
				return g
			}
		}
	}
	panic("Game Over!")
}

func printBoard(g Grid) {
	for _, row := range g {
		fmt.Println(row)
	}
}

func main() {
	r := rand.New(rand.NewSource(99))

	var g Grid

	keyToDir := make(map[string]Direction)
	keyToDir["w"] = Up
	keyToDir["a"] = Left
	keyToDir["s"] = Down
	keyToDir["d"] = Rigth

	for {
		fmt.Println("enter w/a/s/d:")
		dir := ""
		fmt.Scanln(&dir)
		// todo: take input
		g = collapse(keyToDir[dir], g)
		g = randomInsert(g, r)
		printBoard(g)
	}
}

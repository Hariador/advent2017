package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	argCount := len(os.Args)
	if argCount < 2 {
		panic("./day20 <filename>")
	}

	filename := os.Args[1]
	particles := getSourceData(filename)
	fmt.Printf("Starting Particles: %v\n", len(particles))
	for x := 0; x < 1000; x++ {
		for i, p := range particles {
			for j, p1 := range particles {
				if i != j {
					//particles[i].remove = p.collide(p1)
					result := p.collide(p1)
					if result {
						particles[i].remove = result

					}

				}
			}

		}

		particles = prune(particles)

		for i := range particles {
			particles[i].Step()
		}
		if x%10 == 0 {
			fmt.Printf("Remaining Particles: %v\n", len(particles))
		}
	}

	fmt.Printf("Remaining Particles: %v\n", len(particles))

}

func getSourceData(filename string) []particle {
	var result []particle
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		line := scanner.Text()
		p := particle{}
		p.Parse(line)
		p.remove = false
		result = append(result, p)

	}

	return result
}

type particle struct {
	Xp, Yp, Zp int
	Xv, Yv, Zv int
	Xa, Ya, Za int
	vMax       int
	aMax       int
	remove     bool
}

func (p *particle) Parse(sourceLine string) {
	commaLines := strings.Split(sourceLine, ">")
	temp := commaLines[0]
	temp = strings.TrimPrefix(temp, "p=<")

	pValues := strings.Split(temp, ",")
	temp = commaLines[1]
	temp = strings.TrimPrefix(temp, ", v=<")
	vValues := strings.Split(temp, ",")
	temp = commaLines[2]
	temp = strings.TrimPrefix(temp, ", a=<")
	aValues := strings.Split(temp, ",")
	p.Xp, _ = strconv.Atoi(pValues[0])
	p.Yp, _ = strconv.Atoi(pValues[1])
	p.Zp, _ = strconv.Atoi(pValues[2])
	p.Xv, _ = strconv.Atoi(vValues[0])
	p.Yv, _ = strconv.Atoi(vValues[1])
	p.Zv, _ = strconv.Atoi(vValues[2])
	p.Xa, _ = strconv.Atoi(aValues[0])
	p.Ya, _ = strconv.Atoi(aValues[1])
	p.Za, _ = strconv.Atoi(aValues[2])
	p.SetAMax()
	p.SetVMax()
}

func (p *particle) Step() {
	p.Xv = p.Xv + p.Xa
	p.Yv = p.Yv + p.Ya
	p.Zv = p.Zv + p.Za
	p.Xp = p.Xp + p.Xv
	p.Yp = p.Yp + p.Yv
	p.Zp = p.Zp + p.Zv
}

func prune(particles []particle) []particle {
	var temp []particle

	for i, p := range particles {

		if !p.remove {
			temp = append(temp, particles[i])
		}
	}

	return temp
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return n * -1
}

func (p *particle) SetAMax() {
	max := abs(p.Xa)
	if abs(p.Ya) > max {
		max = abs(p.Ya)
	}
	if abs(p.Za) > max {
		max = abs(p.Za)
	}
	p.aMax = max
}

func (p *particle) collide(p1 particle) bool {
	// p.Print()
	// p1.Print()
	result := false
	if p.Xp == p1.Xp && p.Yp == p1.Yp && p.Zp == p1.Zp {
		result = true
	}
	// if p.Xp != p1.Xp {
	// 	return false
	// }
	// if p.Yp != p1.Yp {
	// 	return false
	// }
	// if p.Zp != p1.Zp {
	// 	return false
	// }

	return result
}

func (p *particle) SetVMax() {
	max := abs(p.Xv)
	if abs(p.Yv) > max {
		max = abs(p.Yv)
	}
	if abs(p.Zv) > max {
		max = abs(p.Zv)
	}
	p.vMax = max
}

func (p *particle) Print() {
	fmt.Printf("p=<%v,%v,%v>\tv=<%v,%v,%v>\ta=<%v,%v,%v>\t", p.Xp, p.Yp, p.Zp, p.Xv, p.Yv, p.Zv, p.Xa, p.Ya, p.Za)
	fmt.Printf("aMax=%v\tvMax%v\n", p.aMax, p.vMax)
}

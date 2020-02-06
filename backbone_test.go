package main

import (
	//"fmt"
	//data "housify/data_structures"
	//"math"
	"testing"
)

//               3   4.2  5.4
//   4 +---------X----+----+-+ 4
//     |   6X    |  X |  X |X|
//     |         X----X-X--X-X 2.3
//   2 +---------X      |    |
//     |    X    |	4X  | 3X |
//     |   6     |	    |    |
//   0 +---------+------+----+ 0
//     0         3     4.7   6

func Test_Backbone(t *testing.T) {
	//	bounds, areas := TestTree()
	//	house := Squarify(bounds, areas)
	//	backbone := Backbone(bounds, house)
	//	x1 := 10.0 * (3.0)
	//	x2 := 10.0 * (2.0/5.0*3.0 + 3.0)
	//	x3 := 10.0 * (4.0/7.0*3.0 + 3.0)
	//	x4 := 10.0 * (x2 + 2.0/3.0*(6.0-x2))
	//	x5 := 10.0 * (6.0)
	//	y0 := 10.0 * (0.0)
	//	y1 := 10.0 * (2.0)
	//	y2 := 10.0 * (7.0 / 12.0 * 4.0)
	//	y3 := 10.0 * (4.0)
	//	a := data.Pt{X: math.Round(x1), Y: math.Round(y0)}
	//	b := data.Pt{X: math.Round(x1), Y: math.Round(y1)}
	//	c := data.Pt{X: math.Round(x1), Y: math.Round(y2)}
	//	d := data.Pt{X: math.Round(x2), Y: math.Round(y2)}
	//	e := data.Pt{X: math.Round(x2), Y: math.Round(y3)}
	//	f := data.Pt{X: math.Round(x3), Y: math.Round(y2)}
	//	g := data.Pt{X: math.Round(x3), Y: math.Round(y0)}
	//	h := data.Pt{X: math.Round(x4), Y: math.Round(y2)}
	//	i := data.Pt{X: math.Round(x4), Y: math.Round(y3)}
	//	j := data.Pt{X: math.Round(x5), Y: math.Round(y2)}
	//	//               3   4.2  5.4
	//	//   4 +---------+----e----i-+ 4
	//	//     |         |bath| cl |L|
	//	//     | living  c----d-f--h-j 2.3
	//	//   2 +---------b      |    |
	//	//     |         | mbed | bed|
	//	//     | kitchen |	    |    |
	//	//   0 +---------a------g----+ 0
	//	//     0         3     4.7   6
	//	//living, _ := house.Find("Living")
	//	kitchen, _ := house.Find("Kitchen")
	//	master, _ := house.Find("Master Bed")
	//	bed, _ := house.Find("Bed")
	//	bath, _ := house.Find("Bath")
	//	laundry, _ := house.Find("Laundry")
	//	closet, _ := house.Find("Closet")
	//	kc := kitchen.Value.Center()
	//	mbedc := master.Value.Center()
	//	bedc := bed.Value.Center()
	//	bathc := bath.Value.Center()
	//	lc := laundry.Value.Center()
	//	cc := closet.Value.Center()
	//	// The backbone is going to be some subset of want; this graph is directed,
	//	// too, but it works good enough for testing purposes.
	//	want := data.Graph{
	//		a: []data.Pt{b},
	//		b: []data.Pt{c, kc, mbedc},
	//		c: []data.Pt{d, mbedc, bathc},
	//		d: []data.Pt{e, f, mbedc, bathc},
	//		e: []data.Pt{bathc, cc},
	//		f: []data.Pt{g, h, mbedc, bedc},
	//		g: []data.Pt{mbedc, bedc},
	//		h: []data.Pt{i, j, cc, lc, bedc},
	//		i: []data.Pt{cc, lc},
	//		j: []data.Pt{lc, bedc},
	//	}
	//	for k, vs := range backbone {
	//		fmt.Printf("\nKey: %.2f, %.2f\n", k.X, k.Y)
	//
	//		fmt.Println("Got:")
	//		for _, v := range vs {
	//			fmt.Printf("\t%.2f, %.2f\n", v.X, v.Y)
	//			fmt.Println(data.HasPt(want[data.Pt{math.Round(k.X), math.Round(k.Y)}], v))
	//		}
	//	}
}

package main

import (
	"encoding/csv"
	"errors"
	data "housify/dataStructures"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func loadRoomSizes(fname string, size int) (map[string]float64, error) {
	if size < 0 || size > 2 {
		return nil, errors.New("LoadRoomSizes size parameter out of range")
	}
	roomSizes := make(map[string]float64)
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else {
			f, err := strconv.ParseFloat(record[size+1], 64)
			if err != nil {
				return nil, err
			}
			roomSizes[record[0]] = f
		}
	}
	return roomSizes, nil
}

func loadRoomEdges(fname string) (map[string][]string, error) {
	roomEdges := make(map[string][]string)
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else {
			roomEdges[record[0]] = append(roomEdges[record[0]], record[1])
		}
	}
	return roomEdges, nil
}

func genArea(label string, m map[string]float64) float64 {
	// this formula is more or less arbitrary
	return m[label] * (1 + math.Sqrt(rand.Float64()))
}

func _genHouse(edges map[string][]string, sizes map[string]float64) (data.Rect, *data.FTree) {
	// Service Area
	kitchen := data.FTree{
		Value: genArea("Kitchen", sizes),
		Label: "Kitchen",
		Cnx:   nil,
	}
	dining := data.FTree{
		Value: genArea("Dining", sizes),
		Label: "Dining",
		Cnx:   nil,
	}
	service := data.FTree{
		Value: kitchen.Value + dining.Value,
		Label: "Service",
		Cnx:   []*data.FTree{&kitchen, &dining},
	}

	// Social Area
	living := data.FTree{
		Value: genArea("Living", sizes),
		Label: "Living",
		Cnx:   nil,
	}
	social := data.FTree{
		Value: living.Value,
		Label: "Social",
		Cnx:   []*data.FTree{&living},
	}
	// Private Area
	bath := data.FTree{
		Value: genArea("Bath", sizes),
		Label: "Bath",
		Cnx:   nil,
	}
	mbed := data.FTree{
		Value: genArea("Bed", sizes) * 1.5,
		Label: "Master Bed",
		Cnx:   nil,
	}
	bed := data.FTree{
		Value: genArea("Bed", sizes),
		Label: "Bed",
		Cnx:   nil,
	}
	private := data.FTree{
		Value: bath.Value + bed.Value + mbed.Value,
		Label: "Private",
		Cnx:   []*data.FTree{&bed, &bath, &mbed},
	}

	tree := data.FTree{
		Value: service.Value + social.Value + private.Value,
		Label: "house",
		Cnx:   []*data.FTree{&service, &social, &private},
	}
	sort.SliceStable(tree.Cnx, func(i, j int) bool {
		return tree.Cnx[i].Value < tree.Cnx[j].Value
	})

	// Determine the dimensions of the entire house.
	side := math.Sqrt(tree.Value)
	w := math.Round(side * (1 + rand.Float64()))
	h := math.Round(side * (1 + rand.Float64()))
	bounds := data.Rect{X0: 0, Y0: 0, X1: w, Y1: h, Label: ""}

	return bounds, &tree
}

func GenHouse(edgesFile, sizesFile string) (data.Rect, *data.FTree) {
	edges, err := loadRoomEdges(edgesFile)
	if err != nil {
		panic(err)
	}
	sizes, err := loadRoomSizes(sizesFile, 0)
	if err != nil {
		panic(err)
	}
	return _genHouse(edges, sizes)
}

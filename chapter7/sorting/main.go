package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Delilah", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type trackSlice struct {
	tracks []*Track
	less func(i, j int) bool
}

func (ts *trackSlice) Len() int {
	return len(ts.tracks)
}

func (ts *trackSlice) Less(i, j int) bool {
	return ts.less(i, j)
}

func (ts *trackSlice) Swap(i, j int) {
	ts.tracks[i], ts.tracks[j] = ts.tracks[j], ts.tracks[i]
}

func main() {
	sortSeq := []string{"Title", "Artist", "Year"}

	equal := func(i, j int, field string) bool {
		switch field {
		case "Title":
			return tracks[i].Title == tracks[j].Title
		case "Artist":
			return tracks[i].Artist == tracks[j].Artist
		case "Album":
			return tracks[i].Album == tracks[j].Album
		case "Year":
			return tracks[i].Year == tracks[j].Year
		case "Length":
			return tracks[i].Length == tracks[j].Length
		}
		return false
	}

	differentSort := func(i, j int, field string) bool {
		switch field {
		case "Title":
			return tracks[i].Title < tracks[j].Title
		case "Artist":
			return tracks[i].Artist < tracks[j].Artist
		case "Album":
			return tracks[i].Album < tracks[j].Album
		case "Year":
			return tracks[i].Year < tracks[j].Year
		case "Length":
			return tracks[i].Length < tracks[j].Length
		}
		return false
	}

	less := func(i, j int) bool {
		idx := 0
		for idx < len(sortSeq) && equal(i, j, sortSeq[idx]) {
			idx++
		}
		return idx < len(sortSeq) && differentSort(i, j, sortSeq[idx])
	}

	ts := trackSlice{
		tracks,
		less,
	}
	sort.Sort(&ts)
	printTracks(ts.tracks)
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

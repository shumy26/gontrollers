package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var plotDataFile string = "plot_data.txt"

func plotting(ch chan float64) {
	file, err := os.Create(plotDataFile)
	if err != nil {
		log.Fatal("failure to create plotting data file")
	}
	defer file.Close()

	fmt.Fprintln(file, "# X Y")

	points := make([]float64, 0)

	for {
		value, ok := <-ch
		if !ok {
			break
		}
		points = append(points, value)
	}

	for i := range len(points) {
		fmt.Fprintf(file, "%d %v\n", i, points[i])
	}

	gnuplotArgs := "set terminal pngcairo size 800,600 enhanced font 'Verdana,12'; set output 'pid_plot.png'; set grid lw 1; set xlabel 'Time (s)'; set ylabel 'Controlled Variable'; set style line 1 lc 1 lw 2 ; plot 'plot_data.txt' using 1:2 with lines ls 1 title 'PID controller test';"
	cmd := exec.Command("gnuplot", "-e", gnuplotArgs)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

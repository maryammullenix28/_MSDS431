# Command-Line Applications

This assignment asked students to create a Go program that computed summary statistics for data in the file housesInput.csv that are comparable to that of using R or Python. In addition to assessing the ease of developing such applications, I was also tasked with measuring benchmarks such as memory and processing requirements. The results of my study are:
- Rscript runHouses.R  2.46s user 0.12s system 62% cpu 4.107 total
- python3 runHouses.py  3.28s user 0.96s system 99% cpu 4.260 total
- ( for i in {1..100}; do; go run runHouses.go; done; )  0.13s user 0.13s system 1% cpu 23.409 total

These results show that despite the longer time to compute the Go summary statistics than R or Python, this was done at a fraction of computing power with only 1% CPU power being used. Additonally, it is likely that these measurements can be improved upon with better Go programming.

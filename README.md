# linear-stats

A Go command-line program that reads numbers from a file and calculates:

- The Linear Regression Line.
- The Pearson Correlation Coefficient.

## Requirements

- Go installed on your computer.
- A text file containing one number per line.

## Input File

Example `data.txt`:

```text
2
4
6
```

Each number represents a `y` value.

The `x` value is the line position, starting from `0`:

| x | y |
|---|---|
| 0 | 2 |
| 1 | 4 |
| 2 | 6 |

## Usage

Run the program with:

```sh
go run . data.txt
```

## Output

For the example above, the program prints:

```text
Linear Regression Line: y = 2.000000x + 2.000000
Pearson Correlation Coefficient: 1.0000000000
```

The Linear Regression Line is displayed with 6 decimal places.

The Pearson Correlation Coefficient is displayed with 10 decimal places.

## Linear Regression

The program calculates a line in the form:

```text
y = ax + b
```

Where:

- `a` is the slope of the line.
- `b` is the y-intercept.

## Pearson Correlation Coefficient

The Pearson Correlation Coefficient describes the direction and strength of the linear relationship between the `x` and `y` values.

Its value is between `-1` and `1`:

- `1` means a perfect positive relationship.
- `-1` means a perfect negative relationship.
- `0` means no linear relationship.

## Error Handling

The program reports an error when:

- The file path is missing.
- More than one argument is provided.
- The file cannot be opened.
- A line contains an invalid number.
- A line is empty.
- The file contains `NaN` or infinity.
- The file contains fewer than two numbers.
- All numbers are equal, because the Pearson coefficient is undefined.

Error messages are printed to `stderr`.

## Formatting

To format the Go files, run:

```sh
gofmt -w *.go
```

## Project Structure

```text
linear-stats/
├── data.txt
├── go.mod
├── main.go
├── readData.go
├── linearRegression.go
├── pearson.go
└── README.md
```

## Project Goal

This project was created for the Zone01 `linear-stats` exercise.

It practices:

- File reading.
- Input validation.
- Linear regression.
- Pearson correlation.
- Go slices and loops.
- Error handling.
- Floating-point calculations.


## Testing with the Official Audit Tool

Download and extract the official audit tool. In this example, its folder is named `stat-bin` and is located alongside the `linear-stats` project.

Docker must be installed and running.

### 1. Run the Official Tool

From the `linear-stats` directory, enter the tester directory:

```sh
cd ../stat-bin
```

Give the script executable permissions:

```sh
chmod +x run.sh
```

Generate a dataset and display the reference results:

```sh
./run.sh linear-stats
```

The generated dataset is saved as `data.txt` in the `stat-bin` directory.

### 2. Run the Student Program

Return to the project directory and use the same generated dataset:

```sh
cd ../linear-stats
go run . ../stat-bin/data.txt
```

### 3. Compare the Results

Verify that:

- Both programs use the same output format.
- The Linear Regression Line coefficients have 6 decimal places.
- The Pearson Correlation Coefficient has 10 decimal places.
- All calculated values match.

Repeat the generation and comparison process four times in total, using a new dataset each time.

### Verification

The program was compared with the official tool on four generated datasets. All displayed values matched in all four comparisons.
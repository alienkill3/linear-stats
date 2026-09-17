# PRD — linear-stats

Version: 0.1 — Initial requirements draft

## 1. Purpose

Build a Go command-line program that reads numbers from a file and calculates:

- **Linear Regression Line:** the straight line that best fits the data using the least-squares method.
- **Pearson Correlation Coefficient:** a coefficient describing the direction and strength of the linear relationship between each number's position and its value.

The program is intended for the Zone01 `linear-stats` project and evaluation by an auditor.

## 2. Input and Usage

The user provides a file path as a command-line argument.

Proposed command for this project:

```sh
go run . data.txt
```

The file contains one number per line.

For each line:

- `x` is the line's position, starting at `0`.
- `y` is the number on that line.

Example input:

```text
189
113
121
```

Corresponding data points:

| File line | x | y |
| --- | --- | --- |
| First | 0 | 189 |
| Second | 1 | 113 |
| Third | 2 | 121 |

All data in the file is used, in its original order.

## 3. Required Results

### Linear Regression Line

Calculate the line `y = ax + b`:

- `a`: the slope — how much the estimated value of `y` changes when `x` increases by one.
- `b`: the y-intercept — the estimated value of `y` when `x = 0`.

Both coefficients must be displayed with **6 decimal places**.

### Pearson Correlation Coefficient

Calculate the coefficient `r`, which, when defined, lies between `-1` and `1`:

- `1`: a perfect positive linear relationship.
- `-1`: a perfect negative linear relationship.
- Close to `0`: a weak linear relationship; other types of relationships may still exist.

Display the coefficient with **10 decimal places**.

## 4. Output Format

A successful execution prints exactly two lines, using the labels and order specified in the assignment:

```text
Linear Regression Line: y = <a>x + <b>
Pearson Correlation Coefficient: <r>
```

Output rules:

- Use a period as the decimal separator.
- Round to the required number of decimal places only when formatting the output.
- Do not print prompts, explanations, or debugging messages to standard output.
- Keep the `+ <b>` format even when `b` is negative, for example `+ -2.000000`, following the assignment's output template.

## 5. Proposed Decisions for Unspecified Cases

The following are proposals for the initial implementation, not explicit school requirements.

Revisit them if the official tester requires different behavior.

| Case | Proposed behavior |
| --- | --- |
| No arguments or more than one argument | Print a usage message and exit with an error. |
| Missing or unreadable file | Print a clear message and exit with an error. |
| Non-numeric value or blank line within the data | Print an error identifying the line, numbered from 1, and exit with an error. |
| Empty file or only one value | Report insufficient data to calculate both results. |
| All y values are identical | Report an error: Pearson is undefined, although the regression line is horizontal. |
| Negative numbers or decimals using a period | Accept them. |
| Whitespace before or after a number | Ignore it. |
| Linux/macOS or Windows line endings | Support both formats. |
| A final newline at the end of the file | Accept it; do not treat it as an additional blank record. |
| NaN, infinity, or numeric overflow | Report an error without printing misleading numeric results. |

Error handling rules:

- Write error messages to `stderr`.
- On failure, do not print partial results to `stdout`.
- Return a nonzero exit code on failure.

## 6. Scope

The first version includes:

- File reading.
- Input validation.
- Linear regression calculation.
- Pearson correlation calculation.
- Output in the required format.

The first version does not include:

- A graphical interface.
- Chart rendering.
- Interactive number entry.
- Prediction of the next number's range.
- Additional statistics.

## 7. Acceptance Criteria

- The program runs with a file path supplied as an argument.
- It correctly maps the data to `x = 0, 1, 2, ...` and the corresponding `y` values.
- It correctly calculates the regression line and Pearson coefficient using all data points.
- It follows the required output format and decimal precision exactly.
- It handles the input cases described above without panicking.
- Its results are compared with the official tester when available.

The assignment does not specify the tester's numerical tolerance.

### Example Verification Datasets

Each value goes on a separate line.

| y values | Expected regression line | Expected Pearson coefficient |
| --- | --- | --- |
| 2, 4, 6 | y = 2.000000x + 2.000000 | 1.0000000000 |
| 6, 4, 2 | y = -2.000000x + 6.000000 | -1.0000000000 |
| 1, 0, 1 | y = 0.000000x + 0.666667 | 0.0000000000 |

For the first example, the complete input is:

```text
2
4
6
```

The complete output is:

```text
Linear Regression Line: y = 2.000000x + 2.000000
Pearson Correlation Coefficient: 1.0000000000
```

## 8. Implementation Workflow

- Work in small steps: change → run → check → commit.
- Keep `main()` simple and place individual functions in separate files.
- Panos writes the code with guidance; full code is provided when requested.
- File structure and mathematical formulas will be explained before implementing the relevant functionality.

Requirements source: the `linear-stats` assignment supplied in this conversation.
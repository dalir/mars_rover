# Mars Rover
Program takes in commands and moves one or more robots around Mars.
* The world is modelled as a grid with size `m x n`
* The program reads the input, updates the robots, and print out the final states of the robots
* Each robot has a position `(x, y)`, and an orientation `(N, E, S, W)`
* Each robot can move forward one space `(F)`, rotate left by `90` degrees `(L)`, or rotate right by `90` degrees `(R)`
* If a robot moves off the grid, it is marked as `lost` and its last valid grid position and orientation is recorded
* Going from `x -> x + 1` is in the easterly direction, and `y -> y + 1` is in the northerly direction. i.e. `(0, 0)` represents the south-west corner of the grid

# Input

The input takes the form:
```
4 8
(2, 3, E) LFRFF
(0, 2, N) FFLFRFF
```
The first line of the input `4 8` specifies the size of the grid. The subsequent lines each represent the initial state and commands for a single robot. `(0, 2, N)` specifies the initial state 
of the form `(x, y, orientation)`. `FFLFRFF` represents the sequence of movement commands for the robot. 

# Output
The output should take the form:
```
(4, 4, E)
(0, 4, W) LOST
```
Each line represents the final position and orientation of the robots of the form `(x, y, orientation)`
and optionally whether the robot was lost.

Another example for the input:
```
4 8
(2, 3, N) FLLFR
(1, 0, S) FFRLF
```

The output would be:
```
(2, 3, W)
(1, 0, S) LOST
```

# How to Build and Run 

You can run the program in two ways: directly with the Go compiler or using Docker. If Go is not installed on your machine, the Docker option may be preferable.

## Using Go compiler
1. Build the executable:
    ```
    go build .
    ```
2. Run the program with the help option:
    ```
    ./mars_rover -h
    ```
3. Run the program with an input file:
    ```
    ./mars_rover -i input/test1.txt
    ```

## Using Docker
1. Build the Docker image:
    ```
    make docker-build
    ```
2. Run the program with an input file mounted to the container:
    ```
    docker run -v ${PWD}/input:/input mars_rover --input-file=test1.txt
    ```

Enjoy exploring Mars!
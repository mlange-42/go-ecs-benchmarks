# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche), an archetype-based ECS for Go.

## Benchmark candidates

| ECS | Version |
|-----|---------|
| [Arche](https://github.com/mlange-42/arche) | v0.15.3 |
| [Donburi](https://github.com/yohamta/donburi) | v1.15.7 |
| [go-gameengine-ecs](https://github.com/marioolofo/go-gameengine-ecs) | v0.9.0 |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | v0.0.3 |
| [Volt](https://github.com/akmonengine/volt) | v1.2.0 |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.

In case you are a developer or user of an implementation included here,
feel free to check the benchmarked code for any possible improvements.
Open an issue if you want a version update.

## Benchmarks

Last run: Fri, 14 Feb 2025 16:32:22 UTC  
CPU: AMD EPYC 7763 64-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

Note that the Y axis has logarithmic scale in all plots.
So doubled bar or line height is not doubled time!

All components used in the benchmarks have two `float64` fields.
The initial capacity of the world is set to 1024 where this is supported.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/d71d0056-e919-4b1d-b67d-bf6e495e4376)

| N | Arche | Arche (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.89ns | 45.72ns | 73.79ns | 43.41ns | 16.84ns | 136.26ns |
| 4 | 15.49ns | 13.60ns | 33.06ns | 13.73ns | 6.99ns | 37.58ns |
| 16 | 5.84ns | 5.60ns | 25.59ns | 7.01ns | 4.27ns | 13.21ns |
| 64 | 3.40ns | 4.48ns | 23.15ns | 5.81ns | 3.69ns | 6.95ns |
| 256 | 2.80ns | 2.78ns | 22.88ns | 5.46ns | 3.51ns | 5.27ns |
| 1k | 2.81ns | 2.83ns | 22.77ns | 5.33ns | 3.46ns | 4.87ns |
| 16k | 2.84ns | 2.84ns | 23.93ns | 5.34ns | 3.47ns | 4.78ns |
| 256k | 2.82ns | 2.83ns | 25.70ns | 5.34ns | 3.48ns | 4.75ns |
| 1M | 2.84ns | 2.87ns | 30.53ns | 5.36ns | 3.50ns | 4.76ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

> Volt is left out here, as it is not flexible enough for the required automated setup.

![query32arch](https://github.com/user-attachments/assets/149a53c0-02b6-4591-bd13-47d8a0644f6d)

| N | Arche | Arche (cached) | Donburi | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 51.80ns | 48.58ns | 69.10ns | 42.97ns | 16.83ns |
| 4 | 23.12ns | 17.70ns | 38.70ns | 18.05ns | 14.74ns |
| 16 | 16.26ns | 9.77ns | 34.29ns | 12.04ns | 19.80ns |
| 64 | 9.64ns | 6.15ns | 27.25ns | 8.24ns | 11.70ns |
| 256 | 4.64ns | 3.79ns | 24.37ns | 5.84ns | 5.68ns |
| 1k | 3.34ns | 3.12ns | 25.74ns | 5.46ns | 4.73ns |
| 16k | 2.99ns | 2.85ns | 31.26ns | 5.39ns | 3.61ns |
| 256k | 2.84ns | 2.77ns | 69.29ns | 5.35ns | 3.52ns |
| 1M | 2.89ns | 2.80ns | 102.14ns | 5.37ns | 3.53ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

> Volt is left out here, as it is not flexible enough for the required automated setup.

![query256arch](https://github.com/user-attachments/assets/33d238ea-3aae-4e3e-a184-7562f6edabee)

| N | Arche | Arche (cached) | Donburi | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 61.54ns | 46.59ns | 72.15ns | 51.97ns | 16.85ns |
| 4 | 27.36ns | 13.53ns | 33.16ns | 24.59ns | 7.37ns |
| 16 | 18.58ns | 5.71ns | 24.96ns | 18.70ns | 4.41ns |
| 64 | 16.11ns | 3.37ns | 22.92ns | 18.06ns | 3.79ns |
| 256 | 6.21ns | 2.80ns | 22.69ns | 8.47ns | 3.51ns |
| 1k | 3.64ns | 2.80ns | 24.24ns | 6.11ns | 3.47ns |
| 16k | 2.80ns | 2.75ns | 24.35ns | 5.42ns | 3.47ns |
| 256k | 2.73ns | 2.75ns | 26.40ns | 5.35ns | 3.48ns |
| 1M | 2.77ns | 2.80ns | 34.85ns | 5.40ns | 3.50ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/f8864357-3ed7-4893-a360-31f8507a6a54)

| N | Arche | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 2.49ns | 8.24ns | 9.21ns | 33.64ns | 33.10ns |
| 4 | 2.31ns | 8.11ns | 9.04ns | 33.16ns | 32.64ns |
| 16 | 2.28ns | 8.06ns | 12.03ns | 33.89ns | 33.74ns |
| 64 | 2.24ns | 8.23ns | 11.93ns | 39.57ns | 33.25ns |
| 256 | 2.26ns | 8.48ns | 12.19ns | 40.57ns | 39.60ns |
| 1k | 2.57ns | 11.98ns | 13.34ns | 39.82ns | 44.72ns |
| 16k | 5.97ns | 40.48ns | 35.86ns | 58.34ns | 66.94ns |
| 256k | 7.90ns | 66.97ns | 42.42ns | 77.90ns | 165.95ns |
| 1M | 35.95ns | 227.92ns | 129.37ns | 208.20ns | 269.40ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/a335d5f9-aef0-4b95-8d9d-a6fa277deb04)

| N | Arche | Arche (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 251.66ns | 286.92ns | 1.12us | 364.34ns | 446.94ns | 1.11us |
| 4 | 82.31ns | 71.68ns | 530.10ns | 178.07ns | 202.98ns | 505.28ns |
| 16 | 51.49ns | 27.38ns | 347.84ns | 139.83ns | 124.44ns | 346.71ns |
| 64 | 39.64ns | 14.12ns | 259.41ns | 130.84ns | 105.97ns | 304.90ns |
| 256 | 41.15ns | 11.00ns | 232.87ns | 106.22ns | 100.16ns | 269.76ns |
| 1k | 31.53ns | 10.41ns | 206.65ns | 94.84ns | 270.25ns | 259.95ns |
| 16k | 26.16ns | 8.23ns | 223.88ns | 109.65ns | 277.03ns | 261.11ns |
| 256k | 26.25ns | 8.27ns | 216.36ns | 109.35ns | 367.20ns | 347.32ns |
| 1M | 26.83ns | 8.55ns | 228.71ns | 198.76ns | 460.93ns | 389.23ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/56ea3d87-ccce-4178-b808-a8cea96c856a)

| N | Arche | Arche (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 8.01us | 8.29us | 4.36us | 35.57us | 3.26us | 2.57us |
| 4 | 2.04us | 2.05us | 1.44us | 8.12us | 1.08us | 1.18us |
| 16 | 551.69ns | 498.79ns | 695.23ns | 2.44us | 374.96ns | 635.53ns |
| 64 | 165.18ns | 136.79ns | 459.49ns | 972.79ns | 227.41ns | 503.79ns |
| 256 | 69.65ns | 44.79ns | 344.20ns | 413.84ns | 170.71ns | 365.60ns |
| 1k | 46.30ns | 25.48ns | 308.43ns | 277.26ns | 138.78ns | 294.38ns |
| 16k | 65.64ns | 41.43ns | 402.71ns | 254.35ns | 181.32ns | 450.51ns |
| 256k | 138.85ns | 28.87ns | 479.57ns | 706.13ns | 294.90ns | 675.93ns |
| 1M | 107.98ns | 24.90ns | 481.39ns | 1.47us | 339.72ns | 635.39ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/8121f428-e45f-4f4c-811a-d8891c4301ec)

| N | Arche | Arche (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 364.29ns | 396.37ns | 2.07us | 493.67ns | 633.77ns | 2.43us |
| 4 | 165.38ns | 104.36ns | 1.31us | 276.46ns | 399.02ns | 1.68us |
| 16 | 111.21ns | 33.92ns | 1.11us | 229.44ns | 294.81ns | 1.34us |
| 64 | 95.30ns | 16.20ns | 834.84ns | 182.92ns | 274.81ns | 1.29us |
| 256 | 86.36ns | 11.72ns | 761.22ns | 146.69ns | 240.09ns | 1.11us |
| 1k | 77.12ns | 10.29ns | 743.02ns | 151.59ns | 540.25ns | 1.12us |
| 16k | 76.33ns | 8.26ns | 787.78ns | 158.05ns | 489.09ns | 1.11us |
| 256k | 77.09ns | 8.28ns | 828.67ns | 158.55ns | 601.22ns | 1.14us |
| 1M | 76.16ns | 8.25ns | 790.98ns | 227.39ns | 626.47ns | 1.11us |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/c56052c0-43a6-4f28-bbfe-465700b2542a)

| N | Arche | Arche (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 211.38ns | 549.20ns | 545.78ns | 567.59ns | 382.64ns | 839.83ns |
| 4 | 140.44ns | 151.85ns | 470.35ns | 519.17ns | 353.85ns | 542.16ns |
| 16 | 124.33ns | 50.29ns | 464.84ns | 532.49ns | 347.51ns | 468.10ns |
| 64 | 118.01ns | 24.32ns | 449.25ns | 537.31ns | 366.46ns | 451.92ns |
| 256 | 114.08ns | 10.59ns | 441.02ns | 548.94ns | 377.78ns | 467.65ns |
| 1k | 113.04ns | 7.17ns | 441.54ns | 557.38ns | 773.11ns | 501.97ns |
| 16k | 113.50ns | 7.61ns | 491.34ns | 585.90ns | 823.52ns | 541.68ns |
| 256k | 114.59ns | 8.31ns | 495.43ns | 638.77ns | 996.35ns | 775.17ns |
| 1M | 114.13ns | 9.90ns | 514.30ns | 831.38ns | 1.42us | 962.93ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/8e24c09b-acad-4141-ae71-b6a1e99d1fdc)

| N | Arche | Arche (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 528.40ns | 884.53ns | 1.20us | 994.34ns | 889.21ns | 568.09ns |
| 4 | 490.66ns | 282.53ns | 1.14us | 994.45ns | 855.07ns | 347.71ns |
| 16 | 467.76ns | 138.81ns | 1.14us | 1.02us | 846.79ns | 294.70ns |
| 64 | 456.66ns | 99.24ns | 1.09us | 991.65ns | 882.33ns | 286.10ns |
| 256 | 438.29ns | 37.10ns | 1.11us | 992.67ns | 911.29ns | 289.26ns |
| 1k | 464.01ns | 19.37ns | 1.18us | 1.00us | 1.49us | 299.83ns |
| 16k | 469.56ns | 25.58ns | 1.29us | 1.08us | 1.54us | 301.10ns |
| 256k | 636.49ns | 39.92ns | 1.42us | 1.38us | 1.89us | 396.06ns |
| 1M | 582.74ns | 43.70ns | 1.29us | 1.55us | 2.10us | 422.27ns |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/0f9c0aa6-6b2e-46f1-8bf0-0e0385de7653)

| N | Arche | Arche (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 160.82ns | 591.58ns | 299.84ns | 342.50ns | 160.96ns | 775.76ns |
| 4 | 78.49ns | 159.26ns | 107.93ns | 151.78ns | 69.95ns | 495.80ns |
| 16 | 51.63ns | 55.40ns | 68.57ns | 150.15ns | 47.99ns | 430.35ns |
| 64 | 45.05ns | 26.72ns | 57.12ns | 119.82ns | 40.05ns | 365.03ns |
| 256 | 44.04ns | 11.60ns | 54.14ns | 110.19ns | 41.86ns | 277.00ns |
| 1k | 37.93ns | 7.05ns | 46.27ns | 104.16ns | 35.27ns | 305.43ns |
| 16k | 33.10ns | 6.91ns | 45.66ns | 111.50ns | 49.52ns | 311.82ns |
| 256k | 33.13ns | 6.57ns | 55.77ns | 123.69ns | 70.89ns | 399.52ns |
| 1M | 33.09ns | 7.80ns | 56.63ns | 211.10ns | 165.83ns | 481.55ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/7edcba68-368d-4be7-9a0d-743bc7435de2)

| N | Arche | Arche (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 240.04ns | 570.52ns | 434.42ns | 469.46ns | 170.11ns | 871.22ns |
| 4 | 161.08ns | 168.68ns | 216.37ns | 278.27ns | 66.64ns | 646.10ns |
| 16 | 150.56ns | 56.23ns | 164.05ns | 265.02ns | 48.37ns | 501.01ns |
| 64 | 131.22ns | 27.25ns | 154.58ns | 236.63ns | 41.68ns | 370.87ns |
| 256 | 115.12ns | 12.88ns | 132.90ns | 217.91ns | 40.91ns | 339.62ns |
| 1k | 111.52ns | 7.19ns | 195.37ns | 191.69ns | 35.32ns | 335.20ns |
| 16k | 106.37ns | 6.50ns | 137.83ns | 183.53ns | 49.61ns | 349.35ns |
| 256k | 116.09ns | 6.64ns | 162.57ns | 246.04ns | 87.35ns | 449.90ns |
| 1M | 112.44ns | 8.04ns | 173.68ns | 323.82ns | 164.73ns | 527.63ns |


### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 7.77us | 6.45us | 117.42us | 2.30us | 45.63us |


### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/arche&yohamta/donburi&marioolofo/go-gameengine-ecs&unitoftime/ecs&akmonengine/volt&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/arche,yohamta/donburi,marioolofo/go-gameengine-ecs,unitoftime/ecs,akmonengine/volt&type=Date" alt="Star History Chart" width="600"/>
</a>
</p>

## Running the benchmarks

Run the benchmarks using the following command:

```shell
go run . -test.benchtime=0.25s
```

> On PowerShell use this instead:  
> `go run . --% -test.benchtime=0.25s`

The `benchtime` limit is required for some of the benchmarks that have a high
setup cost which is not timed. They would take forever otherwise.
The benchmarks can take up to one hour to complete.

To run a selection of benchmarks, add their names as arguments:

```shell
go run . query2comp query1in10 query32arch
```

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```

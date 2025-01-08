# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the developer of
> [Arche](https://github.com/mlange-42/arche), an archetype-based ECS for Go.

## Benchmark candidates

| ECS | Type | Version |
|-----|------|---------|
| [Arche](https://github.com/mlange-42/arche) | Archetype | v0.14.5 |
| [Donburi](https://github.com/yohamta/donburi) | Archetype | v1.15.6 |
| [ento](https://github.com/wwfranczyk/ento) | Sparse Set | v0.1.0 |
| [go-gameengine-ecs](https://github.com/marioolofo/go-gameengine-ecs) | Archetype | v0.9.0 |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | Archetype | v0.0.3 |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.

## Benchmarks

Last run: Wed, 08 Jan 2025 00:42:22 UTC  
CPU: AMD EPYC 7763 64-Core Processor

For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

All components used in the benchmarks have two `float64` fields.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/1f486d47-2522-4f40-ba97-e28fe5af3830)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 51.40ns | 48.17ns | 66.82ns | 67.23ns | 43.62ns | 16.76ns |
| 4 | 15.07ns | 13.76ns | 36.04ns | 73.64ns | 14.83ns | 7.13ns |
| 16 | 6.52ns | 5.60ns | 24.83ns | 67.20ns | 6.91ns | 4.27ns |
| 64 | 4.51ns | 4.46ns | 22.80ns | 66.95ns | 5.49ns | 3.62ns |
| 256 | 3.22ns | 3.19ns | 22.73ns | 73.70ns | 5.13ns | 3.65ns |
| 1k | 2.75ns | 2.64ns | 22.81ns | 84.64ns | 5.00ns | 3.46ns |
| 16k | 2.82ns | 2.75ns | 23.47ns | 114.74ns | 4.99ns | 3.46ns |
| 256k | 2.82ns | 2.73ns | 25.09ns | 170.40ns | 5.01ns | 3.45ns |
| 1M | 2.83ns | 2.79ns | 29.54ns | 293.44ns | 5.02ns | 3.47ns |

### Query sparse

`N` entities with components `Position` and `Velocity`.
Additionally, there are `9*N` entities with only `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query1in10](https://github.com/user-attachments/assets/6fbaeb5f-9016-46b3-bbac-b67bbb045a26)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.98ns | 46.72ns | 70.51ns | 96.51ns | 46.16ns | 17.29ns |
| 4 | 17.45ns | 13.50ns | 33.30ns | 95.02ns | 14.79ns | 6.82ns |
| 16 | 6.62ns | 5.37ns | 24.94ns | 96.84ns | 7.06ns | 4.26ns |
| 64 | 3.59ns | 4.50ns | 23.49ns | 99.86ns | 5.56ns | 3.62ns |
| 256 | 3.25ns | 2.79ns | 22.63ns | 112.40ns | 5.14ns | 3.49ns |
| 1k | 2.75ns | 2.80ns | 22.77ns | 131.23ns | 5.05ns | 3.44ns |
| 16k | 2.77ns | 2.77ns | 24.81ns | 176.81ns | 5.00ns | 3.47ns |
| 256k | 2.75ns | 2.75ns | 35.32ns | 329.19ns | 5.10ns | 3.47ns |

### Query fragmented

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/7a5cd2e7-1421-45fe-91d1-d198d6ca43a7)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.97ns | 45.92ns | 67.22ns | 66.90ns | 43.42ns | 16.73ns |
| 4 | 23.37ns | 16.94ns | 38.15ns | 63.75ns | 18.27ns | 14.87ns |
| 16 | 16.69ns | 9.91ns | 33.68ns | 63.81ns | 12.29ns | 20.01ns |
| 64 | 9.88ns | 6.11ns | 29.08ns | 66.47ns | 8.22ns | 11.66ns |
| 256 | 4.76ns | 3.74ns | 24.87ns | 71.80ns | 5.66ns | 5.53ns |
| 1k | 3.38ns | 3.09ns | 25.73ns | 81.51ns | 5.16ns | 4.00ns |
| 16k | 2.91ns | 2.86ns | 32.68ns | 116.16ns | 5.11ns | 3.61ns |
| 256k | 2.80ns | 2.81ns | 72.15ns | 199.60ns | 5.00ns | 3.47ns |
| 1M | 2.84ns | 2.84ns | 97.75ns | 277.05ns | 5.04ns | 3.50ns |

### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/0a8dea5d-e029-4589-8491-07d7ff29a8f5)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 226.92ns | 243.57ns | 1.30us | 1.21us | 615.45ns | 1.06us |
| 4 | 88.21ns | 67.41ns | 638.74ns | 585.94ns | 294.50ns | 500.29ns |
| 16 | 53.79ns | 28.44ns | 434.97ns | 402.38ns | 203.74ns | 349.51ns |
| 64 | 39.12ns | 14.58ns | 342.05ns | 361.87ns | 157.90ns | 284.56ns |
| 256 | 41.68ns | 10.96ns | 276.85ns | 293.00ns | 133.02ns | 238.80ns |
| 1k | 30.27ns | 9.37ns | 262.66ns | 294.56ns | 126.66ns | 405.29ns |
| 16k | 26.34ns | 8.35ns | 290.43ns | 327.67ns | 138.93ns | 421.91ns |
| 256k | 26.19ns | 8.20ns | 280.87ns | 409.07ns | 148.21ns | 486.14ns |
| 1M | 26.20ns | 8.24ns | 258.38ns | 516.94ns | 232.45ns | 594.70ns |

### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/18654532-d303-46e3-aef4-6dd9356a9933)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 9.69us | 9.62us | 4.51us | 1.61us | 43.24us | 3.46us |
| 4 | 2.46us | 2.44us | 1.52us | 690.02ns | 11.35us | 1.29us |
| 16 | 645.49ns | 626.34ns | 762.98ns | 480.13ns | 3.37us | 615.28ns |
| 64 | 197.17ns | 166.46ns | 504.80ns | 393.99ns | 973.06ns | 415.54ns |
| 256 | 76.49ns | 51.07ns | 380.64ns | 327.36ns | 490.03ns | 333.56ns |
| 1k | 52.69ns | 27.15ns | 340.72ns | 316.65ns | 293.96ns | 289.89ns |
| 16k | 97.68ns | 43.09ns | 450.29ns | 531.73ns | 268.26ns | 326.81ns |
| 256k | 170.46ns | 34.70ns | 548.41ns | 648.49ns | 718.86ns | 399.19ns |
| 1M | 290.91ns | 22.13ns | 526.40ns | 759.79ns | 1.62us | 457.09ns |

### Add/remove component

`N` entities with components `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/71ff0e20-76e1-41c6-85f4-0f21e56ec63c)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 207.23ns | 1.16us | 559.42ns | 114.16ns | 553.92ns | 366.42ns |
| 4 | 140.28ns | 310.15ns | 502.31ns | 110.32ns | 518.44ns | 348.36ns |
| 16 | 123.16ns | 99.15ns | 493.00ns | 108.87ns | 532.21ns | 352.35ns |
| 64 | 121.47ns | 45.77ns | 479.88ns | 111.61ns | 541.43ns | 360.77ns |
| 256 | 114.36ns | 32.08ns | 475.13ns | 117.55ns | 559.13ns | 371.88ns |
| 1k | 113.32ns | 28.35ns | 474.61ns | 126.66ns | 553.67ns | 746.18ns |
| 16k | 114.12ns | 28.33ns | 525.61ns | 142.66ns | 580.35ns | 807.25ns |
| 256k | 114.51ns | 28.63ns | 540.10ns | 167.53ns | 628.91ns | 1.04us |
| 1M | 121.12ns | 30.13ns | 513.54ns | 224.95ns | 841.25ns | 1.38us |

### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 9.36us | 6.85us | 67.88us | 522.28us | 5.55us |

## Running the benchmarks

Run the benchmarks using the following command:

```shell
go run . -test.benchtime=0.25s
```

> On PowerShell use this instead:  
> `go run . --% -test.benchtime=0.25s`

The `benchtime` limit is required for some of the benchmarks that have a high
setup cost which is not timed. They would take forever otherwise.

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```

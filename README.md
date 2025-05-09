# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche) and [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

| ECS | Version |
|-----|---------|
| [Arche](https://github.com/mlange-42/arche) | v0.15.3 |
| [Ark](https://github.com/mlange-42/ark) | v0.4.2 |
| [Donburi](https://github.com/yottahmd/donburi) | v1.15.7 |
| [go-gameengine-ecs](https://github.com/marioolofo/go-gameengine-ecs) | v0.9.0 |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | v0.0.3 |
| [Volt](https://github.com/akmonengine/volt) | v1.6.0 |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.
See the section on [Contributing](#contributing) for details.

In case you are a developer or user of an implementation included here,
feel free to check the benchmarked code for any possible improvements.
Open an issue if you want a version update.

## Benchmarks

Last run: Tue, 29 Apr 2025 09:47:51 UTC  
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

![query2comp](https://github.com/user-attachments/assets/3e20bc10-5f23-4f14-af03-6000871e34aa)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 54.46ns | 46.14ns | 51.49ns | 49.70ns | 65.25ns | 44.00ns | 16.21ns | 200.49ns |
| 4 | 18.02ns | 16.85ns | 15.64ns | 15.03ns | 31.97ns | 16.58ns | 6.40ns | 53.90ns |
| 16 | 9.02ns | 8.50ns | 6.29ns | 6.46ns | 23.58ns | 11.26ns | 3.93ns | 16.95ns |
| 64 | 7.08ns | 6.90ns | 4.24ns | 4.51ns | 21.67ns | 9.90ns | 3.45ns | 8.08ns |
| 256 | 6.43ns | 6.39ns | 3.63ns | 3.95ns | 22.32ns | 9.51ns | 3.28ns | 5.55ns |
| 1k | 6.30ns | 6.29ns | 3.56ns | 3.82ns | 22.39ns | 9.38ns | 3.15ns | 4.92ns |
| 16k | 6.30ns | 6.33ns | 3.49ns | 3.79ns | 22.95ns | 9.38ns | 3.18ns | 4.81ns |
| 256k | 6.30ns | 6.29ns | 3.48ns | 3.80ns | 22.06ns | 9.38ns | 3.18ns | 4.76ns |
| 1M | 6.31ns | 6.29ns | 3.48ns | 3.78ns | 26.82ns | 9.41ns | 3.20ns | 4.77ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/89dfda41-36a4-422c-984e-6ba004df4327)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 52.68ns | 48.45ns | 51.53ns | 49.66ns | 62.37ns | 43.98ns | 16.29ns | 217.97ns |
| 4 | 24.77ns | 18.36ns | 26.15ns | 22.31ns | 32.76ns | 21.33ns | 14.82ns | 106.30ns |
| 16 | 18.90ns | 12.27ns | 19.36ns | 14.31ns | 25.77ns | 14.81ns | 25.28ns | 74.01ns |
| 64 | 11.91ns | 8.93ns | 11.48ns | 8.90ns | 22.97ns | 11.53ns | 15.70ns | 38.22ns |
| 256 | 7.68ns | 6.87ns | 5.15ns | 4.67ns | 21.25ns | 9.99ns | 6.26ns | 12.99ns |
| 1k | 6.62ns | 6.43ns | 3.88ns | 4.02ns | 21.88ns | 9.68ns | 4.05ns | 7.08ns |
| 16k | 6.37ns | 6.35ns | 3.60ns | 3.85ns | 30.36ns | 9.44ns | 3.26ns | 5.01ns |
| 256k | 6.33ns | 6.30ns | 3.51ns | 3.81ns | 66.34ns | 9.40ns | 3.20ns | 4.78ns |
| 1M | 6.30ns | 6.31ns | 3.49ns | 3.81ns | 92.70ns | 9.41ns | 3.19ns | 4.77ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/e1dd0db3-3da3-4313-96c7-196324705c07)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 63.59ns | 46.57ns | 63.98ns | 49.38ns | 62.35ns | 55.96ns | 16.48ns | 232.22ns |
| 4 | 29.17ns | 16.88ns | 28.14ns | 15.25ns | 29.69ns | 27.31ns | 9.60ns | 88.85ns |
| 16 | 21.40ns | 8.37ns | 19.25ns | 6.45ns | 21.09ns | 23.02ns | 4.64ns | 57.34ns |
| 64 | 19.60ns | 6.89ns | 16.85ns | 4.81ns | 19.13ns | 21.84ns | 3.58ns | 55.82ns |
| 256 | 9.59ns | 6.39ns | 6.81ns | 3.91ns | 20.02ns | 12.62ns | 3.24ns | 17.46ns |
| 1k | 7.11ns | 6.32ns | 4.36ns | 3.86ns | 43.24ns | 10.17ns | 3.17ns | 7.99ns |
| 16k | 6.34ns | 6.28ns | 3.62ns | 3.79ns | 21.12ns | 9.49ns | 3.21ns | 5.02ns |
| 256k | 6.29ns | 6.29ns | 3.49ns | 3.80ns | 22.06ns | 9.39ns | 3.18ns | 4.77ns |
| 1M | 6.31ns | 6.44ns | 3.52ns | 3.80ns | 25.71ns | 9.39ns | 3.20ns | 4.77ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/2948302b-5bd0-465f-a703-1c1270dce783)

| N | Arche | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 4.72ns | 4.79ns | 10.37ns | 8.73ns | 36.82ns | 36.45ns |
| 4 | 4.17ns | 4.63ns | 8.94ns | 8.52ns | 36.67ns | 35.94ns |
| 16 | 4.06ns | 4.38ns | 8.89ns | 13.78ns | 37.39ns | 36.42ns |
| 64 | 4.13ns | 4.45ns | 9.88ns | 14.09ns | 38.29ns | 36.10ns |
| 256 | 4.02ns | 4.46ns | 9.52ns | 14.40ns | 39.87ns | 37.10ns |
| 1k | 4.39ns | 4.51ns | 11.71ns | 17.74ns | 40.62ns | 40.04ns |
| 16k | 9.78ns | 8.09ns | 36.30ns | 29.86ns | 59.04ns | 63.02ns |
| 256k | 12.59ns | 13.74ns | 135.82ns | 56.75ns | 141.47ns | 94.72ns |
| 1M | 47.64ns | 38.48ns | 201.54ns | 122.01ns | 189.94ns | 237.34ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/51de492e-febe-4ea7-9fde-2afbb94dd721)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 210.25ns | 569.51ns | 293.69ns | 269.05ns | 1.07us | 380.43ns | 431.26ns | 888.39ns |
| 4 | 94.69ns | 148.80ns | 114.81ns | 73.83ns | 471.37ns | 155.48ns | 208.24ns | 391.56ns |
| 16 | 56.45ns | 42.91ns | 79.08ns | 24.54ns | 326.01ns | 128.78ns | 136.54ns | 257.73ns |
| 64 | 42.12ns | 19.43ns | 63.05ns | 14.00ns | 253.93ns | 113.93ns | 113.81ns | 201.77ns |
| 256 | 39.13ns | 10.99ns | 53.07ns | 10.00ns | 191.51ns | 105.62ns | 98.09ns | 175.46ns |
| 1k | 32.25ns | 9.59ns | 47.32ns | 8.86ns | 194.31ns | 104.99ns | 271.51ns | 183.97ns |
| 16k | 27.84ns | 8.32ns | 41.68ns | 7.77ns | 212.04ns | 110.52ns | 277.69ns | 198.49ns |
| 256k | 27.61ns | 8.23ns | 41.51ns | 7.60ns | 200.55ns | 114.65ns | 302.48ns | 220.17ns |
| 1M | 27.99ns | 8.47ns | 41.82ns | 7.75ns | 187.91ns | 181.89ns | 434.34ns | 339.93ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/5d07c86c-1fc3-49fd-83e9-ebbfc912c67f)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 6.56us | 6.72us | 6.31us | 6.17us | 4.29us | 15.96us | 3.14us | 1.76us |
| 4 | 1.95us | 1.90us | 1.90us | 1.80us | 1.34us | 4.30us | 1.07us | 781.85ns |
| 16 | 468.38ns | 443.91ns | 485.97ns | 453.93ns | 670.05ns | 1.28us | 375.89ns | 407.14ns |
| 64 | 158.23ns | 126.19ns | 169.89ns | 120.99ns | 380.52ns | 567.14ns | 233.28ns | 276.35ns |
| 256 | 66.06ns | 41.89ns | 85.71ns | 43.36ns | 307.33ns | 258.17ns | 160.85ns | 227.67ns |
| 1k | 47.71ns | 26.47ns | 55.69ns | 19.38ns | 287.09ns | 210.61ns | 157.42ns | 214.19ns |
| 16k | 62.16ns | 31.08ns | 68.80ns | 38.75ns | 403.88ns | 235.24ns | 171.55ns | 360.72ns |
| 256k | 111.80ns | 26.93ns | 89.43ns | 37.56ns | 462.99ns | 631.34ns | 271.59ns | 515.97ns |
| 1M | 101.54ns | 19.32ns | 88.47ns | 32.25ns | 424.37ns | 1.46us | 367.81ns | 591.62ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/c5f6cf29-424a-4959-9afb-a9805930d5c2)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 421.20ns | 781.00ns | 409.93ns | 390.66ns | 1.89us | 510.41ns | 688.40ns | 2.95us |
| 4 | 173.78ns | 192.45ns | 192.02ns | 108.28ns | 1.28us | 262.95ns | 398.88ns | 2.06us |
| 16 | 112.27ns | 58.87ns | 129.59ns | 34.90ns | 1.10us | 222.98ns | 311.56ns | 1.52us |
| 64 | 98.13ns | 22.34ns | 105.05ns | 15.74ns | 818.06ns | 171.65ns | 243.23ns | 1.28us |
| 256 | 81.77ns | 12.44ns | 97.84ns | 10.81ns | 719.66ns | 161.36ns | 220.65ns | 1.24us |
| 1k | 79.70ns | 9.67ns | 88.83ns | 9.96ns | 732.77ns | 153.97ns | 460.57ns | 1.27us |
| 16k | 75.86ns | 8.25ns | 85.53ns | 7.94ns | 788.51ns | 172.51ns | 475.62ns | 1.27us |
| 256k | 75.34ns | 8.60ns | 85.42ns | 7.60ns | 695.24ns | 175.36ns | 573.88ns | 1.41us |
| 1M | 76.39ns | 8.25ns | 85.15ns | 7.58ns | 665.61ns | 260.89ns | 676.61ns | 1.46us |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/238f4b07-dd34-47af-8ec6-8a2799ae24d8)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 215.27ns | 523.77ns | 214.98ns | 332.42ns | 506.62ns | 576.51ns | 353.84ns | 965.29ns |
| 4 | 152.24ns | 146.12ns | 155.81ns | 93.29ns | 438.58ns | 522.65ns | 333.55ns | 590.35ns |
| 16 | 134.95ns | 47.77ns | 128.00ns | 34.01ns | 426.50ns | 546.52ns | 333.99ns | 483.94ns |
| 64 | 127.92ns | 23.70ns | 121.70ns | 20.52ns | 422.27ns | 537.34ns | 341.97ns | 467.95ns |
| 256 | 134.90ns | 10.31ns | 118.77ns | 10.70ns | 417.71ns | 538.43ns | 352.24ns | 466.26ns |
| 1k | 127.04ns | 7.12ns | 122.28ns | 9.03ns | 420.90ns | 555.24ns | 726.30ns | 473.15ns |
| 16k | 127.53ns | 7.14ns | 119.62ns | 7.16ns | 472.08ns | 588.38ns | 778.43ns | 523.56ns |
| 256k | 126.95ns | 8.20ns | 119.24ns | 8.40ns | 464.54ns | 636.07ns | 862.56ns | 607.03ns |
| 1M | 127.14ns | 9.96ns | 119.74ns | 10.20ns | 475.48ns | 838.44ns | 1.27us | 882.49ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/c737b90a-2889-4ad6-b22e-397c4189fc1f)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 565.09ns | 859.05ns | 435.00ns | 603.83ns | 1.11us | 1.01us | 817.82ns | 2.41us |
| 4 | 513.92ns | 277.74ns | 410.27ns | 218.36ns | 1.06us | 975.66ns | 789.33ns | 1.88us |
| 16 | 489.73ns | 135.82ns | 393.03ns | 118.82ns | 1.03us | 964.11ns | 799.60ns | 1.63us |
| 64 | 478.16ns | 99.52ns | 394.72ns | 95.71ns | 999.54ns | 962.37ns | 871.56ns | 1.59us |
| 256 | 457.27ns | 35.12ns | 381.62ns | 35.26ns | 1.02us | 953.43ns | 872.35ns | 1.64us |
| 1k | 468.96ns | 19.09ns | 422.05ns | 18.81ns | 1.09us | 975.77ns | 1.39us | 1.64us |
| 16k | 470.84ns | 21.24ns | 426.57ns | 25.49ns | 1.18us | 1.03us | 1.43us | 1.73us |
| 256k | 610.91ns | 37.75ns | 536.47ns | 38.01ns | 1.26us | 1.35us | 1.86us | 2.27us |
| 1M | 576.44ns | 41.24ns | 490.42ns | 41.52ns | 1.39us | 1.63us | 1.90us | 2.61us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/5a1061d0-21aa-4644-bd7f-8aad17102f96)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 189.07ns | 619.78ns | 160.50ns | 742.44ns | 251.74ns | 352.95ns | 204.79ns | 388.32ns |
| 4 | 86.13ns | 173.14ns | 69.35ns | 195.04ns | 102.08ns | 157.74ns | 84.55ns | 220.65ns |
| 16 | 55.54ns | 56.38ns | 40.67ns | 61.11ns | 69.06ns | 138.45ns | 54.40ns | 184.94ns |
| 64 | 45.13ns | 26.74ns | 35.25ns | 28.29ns | 60.16ns | 126.20ns | 42.64ns | 147.94ns |
| 256 | 39.59ns | 11.62ns | 30.43ns | 9.92ns | 56.99ns | 104.76ns | 46.21ns | 131.52ns |
| 1k | 34.38ns | 7.98ns | 24.75ns | 5.61ns | 50.83ns | 103.12ns | 35.72ns | 128.66ns |
| 16k | 34.24ns | 6.86ns | 23.38ns | 5.34ns | 51.95ns | 115.12ns | 51.79ns | 151.76ns |
| 256k | 35.14ns | 6.91ns | 23.63ns | 5.14ns | 49.54ns | 147.73ns | 84.17ns | 224.51ns |
| 1M | 34.13ns | 7.81ns | 23.83ns | 6.49ns | 50.34ns | 278.67ns | 187.72ns | 379.36ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/b9f8cd14-0e23-48d3-94e9-6943401c4d51)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 274.85ns | 642.51ns | 233.68ns | 1.05us | 406.11ns | 485.32ns | 180.42ns | 781.95ns |
| 4 | 175.39ns | 168.68ns | 138.78ns | 314.91ns | 202.45ns | 286.39ns | 78.82ns | 594.19ns |
| 16 | 133.46ns | 57.89ns | 113.90ns | 117.53ns | 160.51ns | 204.15ns | 52.09ns | 454.67ns |
| 64 | 135.42ns | 27.05ns | 103.44ns | 72.31ns | 130.84ns | 178.09ns | 40.25ns | 336.12ns |
| 256 | 116.03ns | 12.08ns | 94.24ns | 19.31ns | 117.15ns | 163.50ns | 43.39ns | 321.08ns |
| 1k | 111.18ns | 7.58ns | 74.23ns | 7.88ns | 128.41ns | 164.30ns | 34.99ns | 317.24ns |
| 16k | 105.44ns | 6.61ns | 72.76ns | 7.78ns | 121.93ns | 176.41ns | 51.03ns | 366.70ns |
| 256k | 134.67ns | 7.47ns | 84.41ns | 14.44ns | 197.40ns | 371.86ns | 184.39ns | 565.67ns |
| 1M | 113.36ns | 7.94ns | 82.34ns | 14.85ns | 126.52ns | 431.86ns | 239.32ns | 648.59ns |


### Create world

- Create a new world

| N | Arche | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 8.84us | 13.51us | 2.38us | 180.96us | 2.22us | 40.20us |


### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/arche&mlange-42/ark&yottahmd/donburi&marioolofo/go-gameengine-ecs&unitoftime/ecs&akmonengine/volt&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/arche,mlange-42/ark,yottahmd/donburi,marioolofo/go-gameengine-ecs,unitoftime/ecs,akmonengine/volt&type=Date" alt="Star History Chart" width="600"/>
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

## Contributing

Developers of ECS frameworks are welcome to add their implementation to the benchmarks.
However, there are a few (quality) criteria that need to be fulfilled for inclusion:

- All benchmarks must be implemented, which means that the ECS must have the required features
- The ECS must be working properly and not exhibit serious flaws; it will undergo a basic review by maintainers
- The ECS must be sufficiently documented so that it can be used without reading the code
- There must be at least basic unit tests
- Unit tests must be run in the CI of the repository
- The ECS *must not* be tightly coupled to a particular game engine, particularly graphics stuff
- There must be tagged release versions; only tagged versions will be included here

Developers of included frameworks are encouraged to review the benchmarks,
and to fix (or point to) misuse or potential optimizations.

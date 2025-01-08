# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche), an archetype-based ECS for Go.

## Benchmark candidates

| ECS | Type | Version |
|-----|------|---------|
| [Arche](https://github.com/mlange-42/arche) | Archetype | v0.15.0 |
| [Donburi](https://github.com/yohamta/donburi) | Archetype | v1.15.6 |
| [ento](https://github.com/wwfranczyk/ento) | Sparse Set | v0.1.0 |
| [go-gameengine-ecs](https://github.com/marioolofo/go-gameengine-ecs) | Archetype | v0.9.0 |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | Archetype | v0.0.3 |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.

In case you are a developer or user of an implementation included here,
feel free to check the benchmarked code for any possible improvements.
Open an issue if you want a version update.

## Benchmarks

${info}

For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

All components used in the benchmarks have two `float64` fields.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](query2comp.svg)

${query2comp}

### Query sparse

`N` entities with components `Position` and `Velocity`.
Additionally, there are `9*N` entities with only `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query1in10](query1in10.svg))

${query1in10}

### Query fragmented

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](query32arch.svg))

${query32arch}

### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](create2comp.svg))

${create2comp}

### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](create2comp_alloc.svg))

${create2comp_alloc}

### Add/remove component

`N` entities with components `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](add_remove.svg))

${add_remove}

### Create world

- Create a new world
- Register two components (if applicable)

${new_world}

### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/arche&yohamta/donburi&wwfranczyk/ento&marioolofo/go-gameengine-ecs&unitoftime/ecs&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/arche,yohamta/donburi,wwfranczyk/ento,marioolofo/go-gameengine-ecs,unitoftime/ecs&type=Date" alt="Star History Chart" width="600"/>
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

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```

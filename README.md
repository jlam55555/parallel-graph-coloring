# Parallel and Distributed Graph Coloring
using Gebremedhin-Manne speculative coloring

Homework assignments for ECE465 Cloud Computing

Jonathan Lam & Henry Son

Prof. Marano

---

We chose to focus on the graph coloring problem.

- [Project 1](./src/proj1/README.md): Graph coloring in a single-node,
  multithreaded environment
- [Project 2](./src/proj2/README.md): Graph coloring in a multi-node,
  multithreaded environment (i.e., networked)

Multithreading offered the expected speedup, but the communications cost
for multithreaded environments was too high and caused major slowdowns.

---

### Build Instructions

Make sure Golang is installed on your system. The test systems for projects
1 and 2 use `go1.15.8 linux/amd64` (Debian and Ubuntu).

See the respective README files for build instructions for a specific project.
All build commands should be run from this top-level directory of the repo.

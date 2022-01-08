# langton
## Langton's Ants using Go
---

Simulataneously an experiment in using 2D graphics libraries for Go and exploring emergent visual structures in the Langton's ant program

Follow the instructions here to install SDL2 for your system: [SDL2](https://github.com/veandco/go-sdl2#requirements)

Run `go build` in the project directory to create executable

### Components
---
- **Grid** : Grid object is an abstraction over the SDL2 canvas object with utility functions for rendering and managing colors in the grid.

- **Ants** : Ant object is an autonomous agent that is configured with a set of initial rules and a pointer to it's parent grid as context of where it exists. It can be Move()d and changes the colors in the grid and it's own orientation

### Results
---
![ant_with_fading](/results/ants.gif?raw=true)



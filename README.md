# findPath

The "Path Finder" program is a project I created using the Go programming language. It explores graph algorithms and graphical rendering. Here's a breakdown of the main points:

Technologies:

    Go (Golang): The language used for the project implementation.
    github.com/veandco/go-sdl2/sdl: A third-party library enabling graphical rendering capabilities.

Program Overview:

    Grid Generation: Dynamically creates a grid for exploration, specifying cell size and window dimensions.
    Random Square Placement: Scatters random squares on the grid to act as obstacles, with adjustable quantity and positions.
    Start and End Points: Designates specific points on the grid for pathfinding purposes.
    Pathfinding Algorithm: Utilizes a breadth-first search algorithm to find the shortest path between the start and end points, maneuvering around obstacles.
    Rendering with SDL: Uses the go-sdl2 library to render and visualize the pathfinding results. The SDL renderer displays the grid with the highlighted shortest path.

Constants:

    Window Dimensions: windowWidth and windowHeight determine the graphical window's size.
    Grid Configuration: gridSize specifies each grid cell's size.
    Font Details: fontPath and fontSize define the font used for rendering.
    Random Squares: numRandomSquares controls the quantity of randomly generated obstacle squares.
    Max Square Positions: maxSquarePositionx and maxSquarePositiony set limits for square placement on the grid.

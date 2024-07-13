Common algorithms used for pathfinding in games include:

1. **A\* (A-Star) Algorithm**: A\* is a popular choice for pathfinding and graph traversal. It efficiently finds the shortest path from a start node to a goal node by using heuristics to guide its search.

2. **Dijkstra's Algorithm**: This algorithm finds the shortest paths from the start node to all other nodes in the graph. It's a special case of A\* where the heuristic is always zero, meaning it does not take into account an estimated distance to the goal.

3. **Breadth-First Search (BFS)**: BFS is used for finding the shortest path on unweighted graphs. It explores the neighbor nodes first, before moving to the next level neighbors. It guarantees the shortest path in an unweighted graph.

4. **Depth-First Search (DFS)**: DFS explores as far as possible along each branch before backtracking. It's not typically used for finding the shortest path because it explores all possible paths without considering the goal until it accidentally finds it or explores all options.

5. **Greedy Best-First Search**: This algorithm chooses which node to visit next based on a heuristic of distance to the goal, ignoring the cost to reach a node. It's faster but does not always find the shortest path.

6. **Bellman-Ford Algorithm**: It calculates the shortest paths from a single source vertex to all of the other vertices in a weighted graph. It's slower than Dijkstra's but can handle graphs with negative weight edges.

7. **Jump Point Search (JPS)**: An optimization of the A\* algorithm for uniform-cost grids. It significantly reduces the number of nodes that are evaluated by "jumping" over unnecessary nodes, making it faster than standard A\* for grid-based games.

8. **Swarm Algorithm**: A variation of A\* that uses multiple agents (or swarms) to explore the graph. It can be more efficient in certain types of searches, especially in dynamic or real-time environments.

9. **Theta\* Algorithm**: An any-angle pathfinding algorithm that is an extension of A\*. Unlike A\*, which restricts movements to the graph structure, Theta\* allows for paths that cut corners and move in any direction.

These algorithms can be chosen based on the specific requirements of the game, such as the type of graph (weighted/unweighted, grid-based, etc.), the need for the shortest path vs. a "good enough" path, performance considerations, and whether the environment is static or dynamic.
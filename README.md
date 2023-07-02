
## Computational Complexity:

### Big O notation
Big O notation is a mathematical notation used to describe the upper bound or worst-case scenario (Big O) of an algorithm's time complexity or space complexity. It provides an asymptotic upper limit on the growth rate of an algorithm as the input size approaches infinity. The notation is denoted as O(f(n)), where 'f(n)' represents a function that describes the growth rate of the algorithm. Similarly, Omega notation (Ω) represents the lower bound or best-case scenario, and Theta notation (Θ) provides a tight bound or average-case scenario for the algorithm's complexity. These notations (O, Ω, Θ) help analyze and compare algorithms based on their efficiency and scalability.


### Time Complexity Analysis
Time complexity is like measuring how fast an algorithm is. It tells us how the running time of an algorithm grows as the input gets bigger. Think of it as counting the number of steps the algorithm needs to perform for a given input. The time complexity is usually represented using Big O notation, which gives us an idea of the worst-case scenario and how the algorithm's performance scales with larger inputs.

For example, if an algorithm has a time complexity of O(n), it means that if we double the input size, the algorithm will take roughly twice as long to run. If it has a time complexity of O(n^2), doubling the input size will make the algorithm take roughly four times as long. So, time complexity helps us understand how efficient an algorithm is and how it handles larger and more complex tasks.

### Space Complexity Analysis
Space complexity is like measuring how much memory an algorithm needs. It tells us how the memory usage of an algorithm grows as the input gets bigger. It's important because computers have limited memory resources, and we want to make sure our algorithms don't use more memory than necessary.

Space complexity is also represented using Big O notation. It helps us understand how the additional memory used by the algorithm grows with larger inputs. For example, if an algorithm has a space complexity of O(n), it means that the amount of additional memory it needs is directly proportional to the input size. If it has a space complexity of O(1), it means it requires a constant amount of additional memory, regardless of the input size.

By considering both time complexity and space complexity, we can choose algorithms that are not only fast but also use memory efficiently. These complexities help us evaluate and compare different algorithms to make informed decisions based on the efficiency and resource usage required for our specific tasks.

## Sorting Algorithms:

### [Quick Sort](./sorting/quick.go)
Divide-and-conquer algorithm that selects a pivot, partitions the array, and recursively sorts sub-arrays. 

Time complexity (average):`O(n log n)`.

### [Merge Sort](./sorting/merge.go)
Divide-and-conquer algorithm that divides the array into two halves,
recursively sorts the sub-arrays, and then merges them.

Time complexity (average):`O(n log n)`.

### [Heap Sort](./sorting/heap.go)
Utilizes a binary heap data structure.

Time complexity (average): `O(n log n)`.

### [Insertion Sort](./sorting/insertion.go)
Simple sorting algorithm that builds the final sorted array
one item at a time.

Time complexity (average): `O(n^2)`.

### [Selection Sort](./sorting/selection.go)
In-place comparison sorting algorithm.

Time complexity (average): `O(n^2)`.

## Searching Algorithms:

### Binary Search

### Depth-First Search (DFS)

### Breadth-First Search (BFS)

## Graph Algorithms:

### Dijkstra's Algorithm (shortest path)
### Bellman-Ford Algorithm (shortest path with negative weights)
### Kruskal's Algorithm (minimum spanning tree)
### Prim's Algorithm (minimum spanning tree)

## Dynamic Programming:

### Knapsack Problem
### Longest Common Subsequence (LCS)
### Fibonacci Sequence


## Tree Algorithms:

### Binary Search Tree (BST) operations
### Tree Traversal (pre-order, in-order, post-order)


## String Algorithms:

### String Matching (KMP Algorithm, Rabin-Karp Algorithm)
### Longest Common Substring
### Levenshtein Distance (Edit Distance)


## Greedy Algorithms:

### Huffman Coding
### Interval Scheduling


## Hashing and Hash Tables:

### Hash Functions
### Collision Resolution Techniques (Chaining, Open Addressing)

## Advanced Data Structures:

### AVL Trees
### Red-Black Trees
### Trie (Prefix Tree)
### B-trees

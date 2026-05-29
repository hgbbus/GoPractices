import java.util.*;

class Edge {
    // from ?
    int to;
    int weight;

    Edge(int to, int weight) {
        this.to = to;
        this.weight = weight;
    }
}

class Pair implements Comparable<Pair> {
    int node;
    int weight;

    Pair(int node, int weight) {
        this.node = node;
        this.weight = weight;
    }

    @Override
    public int compareTo(Pair other) {
        return this.weight - other.weight;
    }
}

public class PrimsAjacencyListPQ {
    static void primMST(List<List<Edge>> graph, int V) {
        boolean[] visited = new boolean[V];

        PriorityQueue<Pair> pq = new PriorityQueue<>();

        // Start from node 0
        pq.offer(new Pair(0, 0));

        int mstWeight = 0;

        System.out.println("Edges in MST:");
        while (!pq.isEmpty()) {
            Pair current = pq.poll();
            int u = current.node;

            // Skip if already visited
            if (visited[u]) continue;
            visited[u] = true;
            mstWeight += current.weight;

            System.out.println(
                "Node: " + u +
                " Weight: " + current.weight
            );

            for (Edge edge : graph.get(u)) {
                int v = edge.to;
                if (!visited[v]) {
                    pq.offer(new Pair(v, edge.weight));
                }
            }
        }

        System.out.println("Total MST Weight = " + mstWeight);
    }

    public static void main(String[] args) {

        int V = 5;

        List<List<Edge>> graph = new ArrayList<>();

        for (int i = 0; i < V; i++) {
            graph.add(new ArrayList<>());
        }

        // Undirected graph
        addEdge(graph, 0, 1, 2);
        addEdge(graph, 0, 3, 6);
        addEdge(graph, 1, 2, 3);
        addEdge(graph, 1, 3, 8);
        addEdge(graph, 1, 4, 5);
        addEdge(graph, 2, 4, 7);
        addEdge(graph, 3, 4, 9);

        primMST(graph, V);
    }

    static void addEdge(List<List<Edge>> graph,
                        int u,
                        int v,
                        int weight) {

        graph.get(u).add(new Edge(v, weight));
        graph.get(v).add(new Edge(u, weight));
    }
}

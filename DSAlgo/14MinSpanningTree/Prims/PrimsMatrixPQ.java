import java.util.*;

class Edge implements Comparable<Edge> {
    int     vertex;     // this vertext to a vertex in the current MST
    int     parent;
    int     weight;

    Edge(int v, int p, int w) {
        vertex = v;
        parent = p;
        weight = w;
    }

    @Override
    public int compareTo(Edge o) {
        return this.weight - o.weight;
    }
}

public class PrimsMatrixPQ {
    static int V = 5;

    static void primMST(int[][] graph) {
        boolean[] visited = new boolean[V];

        PriorityQueue<Edge> pq = new PriorityQueue<>();
        pq.offer(new Edge(0, -1, 0));   // Vertex 0 has weight 0 to beginning with

        int totWeight = 0;
        int totVisited = 0;

        System.out.println("Vertex picked:");
        while (/*!pq.isEmpty()*/ totVisited < V) {
            Edge cur = pq.poll();

            int u = cur.vertex;
            if (visited[u]) continue;

            System.out.println("Vertex: " + u + 
                                ", parent: " + cur.parent + 
                                ", weight: " + cur.weight);
            visited[u] = true;
            totVisited++;
            totWeight += cur.weight;

            // update neighbors
            for (int v = 0; v < V; v++) {
                if (graph[u][v] > 0 && !visited[v]) {
                    pq.offer(new Edge(v, u, graph[u][v]));
                }
            }
        }

        System.out.println("Total weight = " + totWeight);
    }

    public static void main(String[] args) {
        int[][] graph = {
                {0, 2, 0, 6, 0},
                {2, 0, 3, 8, 5},
                {0, 3, 0, 0, 7},
                {6, 8, 0, 0, 9},
                {0, 5, 7, 9, 0}
        };

        primMST(graph);
    }
}

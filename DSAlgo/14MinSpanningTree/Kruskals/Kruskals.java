import java.util.*;

class Edge implements Comparable<Edge> {
    int u, v, w;

    Edge(int u, int v, int w) {
        this.u = u;
        this.v = v;
        this.w = w;
    }

    @Override
    public int compareTo(Edge o) {
        return this.w - o.w;
    }
}

public class Kruskals {
    static void kruskalMST(List<Edge> edges, int V) {
        Collections.sort(edges);

        DisjointSet ds = new DisjointSet(V);
        int mstWeight = 0;

        System.out.println("Edges in MST:");
        for (Edge e : edges) {
            if (ds.find(e.u, true) == ds.find(e.v, true)) {
                continue;
            }
            ds.union(e.u, e.v);
            mstWeight += e.w;
            System.out.println(e.u + " -- " + e.v + " (weight: " + e.w + ")");
        }
        System.out.println("Total weight of MST: " + mstWeight);
    }

    public static void main(String[] args) {
        int V = 5;

        List<Edge> edges = new ArrayList<>();

        edges.add(new Edge(0, 1, 2));
        edges.add(new Edge(0, 3, 6));
        edges.add(new Edge(1, 2, 3));
        edges.add(new Edge(1, 3, 8));
        edges.add(new Edge(1, 4, 5));
        edges.add(new Edge(2, 4, 7));
        edges.add(new Edge(3, 4, 9));

        kruskalMST(edges, V);
    }
}

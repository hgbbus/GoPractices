class DisjointSet {
    // Vertices are from 0 to V-1
    int     V;
    int[]   parent;
    int[]   size;   // prefer size over rank

    public DisjointSet(int V) {
        this.V = V;
        parent = new int[V];
        size = new int[V];
        for (int i = 0; i < V; i++) {
            parent[i] = i;
            size[i] = 1;
        }
    }

    private int _find(int x) {
        assert 0 <= x && x < V;
        while (parent[x] != x) {
            x = parent[x];
        }
        return parent[x];
    }

    private int _find_with_compression(int x) {
        assert 0 <= x && x < V;
        if (parent[x] != x) {
            parent[x] = _find_with_compression(parent[x]);
        }
        return parent[x];
    }

    public int find(int x, boolean compression) {
        if (compression) {
            return _find_with_compression(x);
        } else {
            return _find(x);
        }
    }

    public boolean union(int x, int y) {
        assert 0 <= x && x < V;
        assert 0 <= y && y < V;

        if (find(x, true) == find(y, true))
            return false;

        // reprentatives
        int rx = find(x, true);
        int ry = find(y, true);

        if (size[rx] > size[ry]) {
            parent[ry] = rx;
            size[rx] += size[ry];
        } else {
            parent[rx] = ry;
            size[ry] += size[rx];
        }

        return true;
    }
}


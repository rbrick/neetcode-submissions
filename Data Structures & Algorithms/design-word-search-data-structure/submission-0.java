class WordDictionary {

    class Node {
        Node[] children = new Node[26];
        boolean end;
        char c;


        public boolean hasChild(char ch) {
            return children[ch - 'a'] != null;
        }

        @Override
        public String toString() {
            return "%c".formatted(c);
        }
    }

    private final Node root;

    public WordDictionary() {
        this.root = new Node();
    }


    public void addWord(String word) {
        Node current = this.root;

        for (int i = 0; i < word.length(); i++) {
            char ch = word.charAt(i);
            Node child = current.children[ch - 'a'];

            if (child == null) {
                child = current.children[ch - 'a'] = new Node();
            }
            child.c = ch;
            current = child;

        }

        current.end = true;
    }


    // bfs
    public boolean search(String word) {
        var layer = new HashSet<Node>();
        var ch = word.charAt(0);
        // initialize the layer with either all the children. i.e. ".ay"
        if (ch == '.') layer.addAll(Arrays.stream(this.root.children).filter(Objects::nonNull).toList());
        else if (this.root.hasChild(ch)) {
            // or just one of them.
            // i.e. "b.."
            layer.add(this.root.children[ch - 'a']);
        } else return false;

        for (int i = 1; i < word.length(); i++) {
            ch = word.charAt(i);
            var curr = ch;
            // reduce the current layer to the new layer of nodes
            if (ch == '.') {
                // collect all the nodes children
                var nextLayer = layer.stream().filter(Objects::nonNull).map(n -> n.children).toList();
                layer.clear();
                nextLayer.forEach(a -> layer.addAll(Arrays.stream(a).filter(Objects::nonNull).toList()));
            } else {
                var nextLayer = layer.stream().filter(Objects::nonNull).filter(n -> n.hasChild(curr)).map(n -> n.children[curr - 'a']).toList();
                layer.clear();
                layer.addAll(nextLayer);
            }
        }

        return layer.stream().anyMatch(node -> node.end);
    }
}
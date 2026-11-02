class PrefixTree {

    class Node {
        Node[] children = new Node[26];
        boolean end;


        @Override
        public String toString() {
            return "{end: %b}".formatted(end);
        }
    }

    private Node root;

    public PrefixTree() {
        this.root = new Node();
    }

    public void insert(String word) {
        Node current = this.root;


        for (int i = 0; i < word.length(); i++) {
            char ch = word.charAt(i);
            Node child = current.children[ch - 'a'];

            if (child == null) {
                child = current.children[ch - 'a'] = new Node();
            }
            current = child;
        }

        current.end = true;
    }

    public boolean search(String word) {
        Node current = this.root;
        int depth = 0;
        for (int i = 0; i < word.length(); i++) {
            char ch = word.charAt(i);
            Node child = current.children[ch - 'a'];

            if (child == null) {
                return false;
            }

            current = child;
            depth++;
        }

        return current.end;
    }

    public boolean startsWith(String prefix) {
        Node current = this.root;
        for (int i = 0; i < prefix.length(); i++) {
            char ch = prefix.charAt(i);
            Node child = current.children[ch - 'a'];

            if (child == null) {
                return false;
            }

            current = child;
        }

        return true;
    }
}

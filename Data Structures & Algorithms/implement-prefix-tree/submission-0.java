class PrefixTree {

    class Node {
        Node[] children = new Node[26];
        char c;
        String str;


        @Override
        public String toString() {
            return "{char: %c, str: %s}".formatted(c, str);
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
                child.c = ch;
            }
            current = child;
        }

        current.str = word;
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

        return current.str != null && current.str.equals(word);
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

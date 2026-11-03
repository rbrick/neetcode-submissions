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


    private boolean dfs(String word, int depth, Node current) {
        if (current == null) return false;
        if (depth == word.length()) return current.end;
        
        char ch = word.charAt(depth);

        if (ch == '.') {
            boolean found = false;
            for (Node child : current.children) {
                if (dfs(word, depth + 1, child)) found = true;
            }
            return found;
        } else if (!current.hasChild(word.charAt(depth))) return false;

        return dfs(word, depth + 1, current.children[ch - 'a']);
    }
    public boolean search(String word) {
        return dfs(word, 0, this.root);
    }
}
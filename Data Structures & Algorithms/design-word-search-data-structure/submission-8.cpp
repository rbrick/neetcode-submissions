struct node {
    node* children[26];
    bool end;
};


class WordDictionary {
public:
    node* root;

    WordDictionary() {
        this->root = new node();
        for (int i = 0; i < 26; i++) {
            this->root->children[i] = nullptr;
        }
    }
    
    void addWord(string word) {
       node* curr = this->root;
       for (char c : word) {
           int idx = c - 'a';
           if (!curr->children[idx]) {
               curr->children[idx] = new node();
           }
           curr = curr->children[idx];
       }

       curr->end = true;
    }
    


    bool search(string word) {
        return this->dfs(this->root, word, 0);
    }
private:
    bool dfs(node* curr, string& word, int idx) {
        if (idx == word.size()) return curr->end;
        if (word[idx] == '.') {
            for (int i = 0; i < 26; i++) {
                if (curr->children[i] && dfs(curr->children[i], word, idx + 1)) return true;
            }
            return false;
        } else {
            int i = word[idx] - 'a';
            return curr->children[i] && dfs(curr->children[i], word, idx + 1);
        }
    }
};
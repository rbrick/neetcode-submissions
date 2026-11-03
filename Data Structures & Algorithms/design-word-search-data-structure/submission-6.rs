
struct Node {
    children: [Option<Box<Node>>; 26],
    end: bool,
}

impl Node {
    fn new() -> Self {
        Self {
            children: std::array::from_fn(|_| None),
            end: false,
        }
    }

    fn insert(&mut self, c: char) {
        let i = (c as u8 - b'a') as usize;

        if self.children[i].is_none() {
            self.children[i] = Some(Box::new(Node::new()));
        }
    }
}

struct WordDictionary {
    root: Node,
}

impl WordDictionary {
    pub fn new() -> Self {
        Self {
            root: Node::new(),
        }
    }

    pub fn add_word(&mut self, word: String) {
        let mut curr = &mut self.root;

        for c in word.chars() {
            let i = (c as u8 - b'a') as usize;
            curr.insert(c); // insert the word
            curr = curr.children[i].as_mut().unwrap();
        }

        curr.end = true;
    }

    fn dfs(&self, word: &String, depth: usize, node: Option<&Node>) -> bool {
        match node {
           None => false,
           Some(n) => {
            if depth == word.len() {
                return n.end;
            }

            match word.chars().nth(depth).unwrap() {
                '.' => n.children.iter().filter(|n| n.is_some()).any(|n| self.dfs(word, depth+1, n.as_deref())),
                ch => {
                    let index = (ch as u8 - b'a') as usize;
                    let child = &n.children[index];
                    if child.is_none() {
                        return false;
                    }

                    return self.dfs(word,depth+1, child.as_deref())
                }
            }
           }
        }
    }

    pub fn search(&self, word: String) -> bool {
        // layer of node
        return self.dfs(&word, 0,  Some(&self.root));
    }
}

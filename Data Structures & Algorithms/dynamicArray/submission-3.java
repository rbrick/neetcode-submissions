class DynamicArray {

        private int[] backingArray;


    // cap = total # of available elements
    // size = total # of set elements
    private int cap, size;


    public DynamicArray(int capacity) {
        this.backingArray = new int[capacity];
        this.cap = capacity;
    }

    public int get(int i) {
        // Docs say assume I is valid
        return backingArray[i];
    }

    public void set(int i, int n) {
        // docs say assume I is valid
        backingArray[i] = n;
    }

    public void pushback(int n) {
        if (this.size >= this.backingArray.length) {
            this.resize();
        }

        this.backingArray[this.size] =  n;
        this.size++;
    }

    public int popback() {
        int last = this.backingArray[this.size - 1];
        this.backingArray[this.size - 1] = 0;
        this.size--;
        return last;
    }
    private void resize() {
        int[] resizedArray = new int[this.backingArray.length * 2];
        System.arraycopy(this.backingArray, 0, resizedArray, 0, this.size);
        this.backingArray = resizedArray;
        this.cap *= 2;
    }

    public int getSize() {
        return this.size;
    }

    public int getCapacity() {
        return this.backingArray.length;
    }

}

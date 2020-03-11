//Partition这样实现比较简单
   int partition(int[] nums, int lo, int hi) {
        int v = nums[lo];
        int i = lo, j = lo + 1;
        for (; j <= hi; j ++) {
            if (nums[j] < v) {
                i ++;
                swap(nums, i, j);
            }
        }
        swap(nums, lo, i);
        return i;
    }
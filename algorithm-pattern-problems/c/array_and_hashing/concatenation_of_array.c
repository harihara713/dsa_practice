#include <stdlib.h>

int* getConcatenation(int* nums, int numsSize, int* returnSize) {
    int *ans = malloc(2 * numsSize * sizeof(int));
    *returnSize = 2 * numsSize;

    for (int i = 0; i < numsSize; i++)
    {
        ans[i] = nums[i];
        ans[i + numsSize] = nums[i];
    }

    return ans;
}

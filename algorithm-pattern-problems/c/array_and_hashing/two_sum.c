#include <stdlib.h>

// NOTE: We can solve this using a hashmap optimally

// TODO: Solve it using hashmap

int* twoSum(int* nums, int numsSize, int target, int* returnSize) 
{

  for (int i = 0; i < numsSize - 1; i++)
  {
    for (int j = i+1; j < numsSize; j++)
    {
      if (nums[j] == target - nums[i])
      {
        int *result = malloc(2 * sizeof(int));
        result[0] = i;
        result[1] = j;
        *returnSize = 2;

        return result;
      }
    }
  }

  *returnSize = 2;
  return NULL;
  // Time: O(n^2)
  // Space: O(1)
}

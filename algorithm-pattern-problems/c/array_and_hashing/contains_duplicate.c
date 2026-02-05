#include <stdlib.h>
#include <stdbool.h>

int compare(const void *a, const void *b){
	return *(int*)a - *(int*)b;
}

bool containsDuplicate(int* nums, int numsSize) {
   // sort the array
   qsort(nums, numsSize, sizeof(int), compare);

   // check for adjacent duplicate element
   for (int i = 1; i < numsSize; i++)
   {
	  if (nums[i] == nums[i-1])
	  {
		  return true;
	  }
   }

   return false;
   // Time: O(n logn)
   // space: O(1)
}

#include <stdbool.h>
#include <stdlib.h>

bool isAnagram(char* s, char* t) {
  int hash[27];
  for (int i = 0; i < 27; i++)
  {
    hash[i] = 0;
  }
  
  int i = 0;
  while (s[i] != '\0') 
  {
    hash[s[i] - 'a']++;
    i++;
  }

  i = 0;
  while (t[i] != '\0')
  {
    hash[t[i] - 'a']--;
    i++;
  }

  i = 0;
  for (i = 0; i < 27; i++)
  {
    if (hash[i] != 0)
    {
      return false;
    }
  }

  return true;
  // Time: O(n)
  // Space: O(1)
}

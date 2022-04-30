"""
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.
Input: s = "hello"
Output: "holle"
"""


class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']
        reverse = [0] * len(s)
        j = len(s)
        s = list(s)

        # store the vowels separately
        for i in range(len(s)):
            if s[i] in vowels:
                reverse[j] = s[i]
                j += 1
        for i in range(len(s)):
            if s[i] in vowels:
                print(j)
                j -= 1
                print(j)
                s[i] = reverse[j]
        return ''.join(s)


input = "hello"
result = Solution()
print(result.reverseVowels(input))

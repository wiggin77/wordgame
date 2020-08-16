# wordgame

Find all the dictionary words for any group of letters.

## How

For efficiency, not every permutation of letters is checked. Instead, the dictionary is arranged in a tree where each word's characters are children of the previous character. For example, a dictionary consisting of [bat, cat, sat] would form a tree like:

```plain
b  c  s
\  | /
   a
   |
   t
```

When searching this dictionary for permutations of [b, c, s, a, t] it will not search for permutations such as "atbcb" since 'a' was not found at the tree root, meaning no words begin with 'a'.  "bctas" is not checked because 'c' was not found in the second level.  Once a letter is not found in its corresponding level within the tree, no further checks are done against the remaining letters.

For larger numbers of letters this can dramatically reduce the number of dictionary searches versus simple brute force.

## Why?

I was playing an android game called "Word Collect" where you try to find all the words in a scrambled collection of letters. I set the difficulty at "challenging" and it was nice distraction. When I got stuck I had to watch video ads to get hints which is annoying... hence this app. Figuring out an efficient algorithm became more interesting than playing the game.

Unfortunately this app works too well, and the game is beaten. On to something else...

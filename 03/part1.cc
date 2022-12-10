#include <fstream>
#include <iostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

using namespace std;

typedef unordered_map<char, int> umap;

int get_pr(char ch) {
  if (ch >= 97) {
    return int(ch) - 96;
  } else {
    return int(ch) - 65 + 27;
  }
}

int solve(vector<string> lines) {
  umap hashMap;
  umap cms;
  for (int i = 0; i < lines[0].size(); i++) {
    hashMap[lines[0][i]] = 1;
  }
  for (int j = 0; j < lines[1].size(); j++) {
    if (hashMap[lines[1][j]]) {
      cms[lines[1][j]] = 1;
    }
  }
  for (int i = 0; i < lines[2].size(); i++) {
    char ch = lines[2][i];
    if (cms[ch]) {
      cout << ch << "-" << int(ch) << "\n";
      return get_pr(ch);
    }
  }
  return 0;
}

void debug(int alphas[52]) {
  for (int i = 0; i < 52; i++) {
    cout << alphas[i] << " ";
  }
}

int solve_op2(vector<string> lines) {
  int alphas[52] = {1};
  unordered_set<char> str;
  unordered_map<char, int> res;
  for (int t = 0; t < 3; t++) {
    for (int i = 0; i < lines[t].size(); i++) {
      str.insert(lines[t][i]);
    }
  }
  return 0;
}

int main(void) {
  vector<string> lines = {"vJrwpWtwJgWrhcsFMMfFFhFp",
                          "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
                          "PmmdzqPrVvPwwTWBwg"};

  cout << solve_op2(lines) << endl;

  return 0;
}

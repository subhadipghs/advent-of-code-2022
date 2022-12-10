#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <unordered_map>
#include <vector>

using namespace std;

vector<string> split(string s, char delim)
{
  stringstream ss(s);
  vector<string> out;
  while (!ss.eof()) {
    string token;
    getline(ss, token, delim);
    out.push_back(token);
  }
  return out;
}

bool is_win(string a, string b)
{
  return (a == "A" && b == "Y") || (a == "B" && b == "Z") || (a == "C" && b == "X");
}

bool is_draw(string a, string b)
{
  return (a == "A" && b == "X") || (a == "B" && b == "Y") || (a == "C" && b == "Z");
}

// 2.a
int get_score(string a, string b) 
{
  unordered_map<string, int> pts = {
    {"X", 1},
    {"Y", 2},
    {"Z", 3},
  };
  int score = 0;
  if (is_win(a, b)) {
    score = pts.at(b) + 6;
  } else if (is_draw(a, b)) {
    score = pts.at(b) + 3;
  } else {
    score = pts.at(b);
  }
  return score;
}

int main(int argc, char** argv)
{
  ifstream fin;
  string line;
  fin.open("2.data");
  int total = 0;
  while (getline(fin, line)) {
    vector<string> m = split(line, ' ');
    total += get_score(m[0], m[1]);
  }
  cout << total << endl;
  return 0;
}

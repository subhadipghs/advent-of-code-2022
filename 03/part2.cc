#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <unordered_map>
#include <vector>
#include <chrono>

using namespace std;
using namespace std::chrono;

typedef unordered_map<string, int> umsi;
typedef unordered_map<string, string> umss;

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

// 2.b
// A - X - ROCK
// B - Y - PAPER
// C - Z - SCISSORS
int get_score(string a, string b) 
{
  umsi pts = {
    {"X", 1},
    {"Y", 2},
    {"Z", 3},
  };
  umss draw = {
    {"A", "X"},
    {"B", "Y"},
    {"C", "Z"},
  };
  umss lose = {
    {"A", "Z"},
    {"B", "X"},
    {"C", "Y"},
  };
  umss win = {
    {"A", "Y"},
    {"B", "Z"},
    {"C", "X"},
  };
  int score = 0;
  if (b == "X") {
    score = pts.at(lose.at(a));
  } else if (b == "Y") {
    score = pts.at(draw.at(a)) + 3;
  } else {
    score = pts.at(win.at(a)) + 6;
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
  fin.close();
  return 0;
}

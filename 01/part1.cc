#include <iostream>
#include <fstream>
#include <string>

using namespace std;


int main(int argc, char **argv)
{
  ifstream fio;
  string line;

  fio.open("1.data");
  int max = 0;
  int a = 0;

  while (getline(fio, line)) {
    if (line != "") {
      a += stoi(line);
      if (a > max) max = a;
    } else {
      a = 0;
    }
  }
  cout << "result " << max << endl;
  fio.close();
  return 0;
}

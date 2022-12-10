#include <iostream>
#include <fstream>
#include <queue>
#include <string>

using namespace std;


int main(int argc, char **argv)
{
  ifstream fio;
  string line;

  priority_queue<int> pq;

  fio.open("1.data");
  int a = 0;

  while (getline(fio, line)) {
    if (line != "") {
      a += stoi(line);
    } else {
      pq.push(a);
      a = 0;
    }
  }
  int k = 3, res = 0;
  while (k-- && !pq.empty()) {
    cout << pq.top() << '\n';
    res += pq.top();
    pq.pop();
  }
  cout << "result " << res << endl;
  fio.close();
  return 0;
}


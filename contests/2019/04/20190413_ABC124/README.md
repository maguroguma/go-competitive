# ABC124 感想

全完できたが、問題が易しめだったこともあり、パフォ1300ちょいだった。
1WA込で75分ぐらいだった。

- A問題は Max 関数を使えばスッキリするが、めんどかったので愚直なif文でゴリ推した。
- B問題は N が最大でも100程度だったので、難しいことを考えずに2重ループを回した。
  - 暫定の最大の高さの山を保持しておけば1重ループで済む。
- C問題は `10101...` か `01010...` の2パターンしかないので、それぞれのケースを調べて比較して終わり。
  - `rune` 型で読み込んだことを忘れていたため、サンプルを通すまでに時間がかかったのが反省。
    - このミスは今後なくしたい。。
- D問題は0を埋める場所を連続するようにすればよいが、アルゴリズムの実装が若干難しい。
  - 自分は結局、連続区間の圧縮・両端にダミーデータを置いて状況を正規化・累積和あたりのテクニックを使って解いた。
    - 半開区間風な累積和の扱いにはもっと慣れたい（というかスニペット化してしまうべきでは。。）。
  - 効率的な方法は解説放送で学びたい。

---

## ランレングス圧縮

http://algoful.com/Archive/Algorithm/RLE

今回用いた、連続部分の圧縮アルゴリズムの名前。
単純なものだが、今後もお世話になることもあるかもしれない。

## D問題（解説放送）

[解説放送リンク](https://www.youtube.com/watch?v=FRzpDCx17vw&feature=push-lsb&attr_tag=J0ywr-iRHWSeCmy1%3A6)

累積和を使うほうがシンプルで良さそうだし、自分の書いた方法も大方間違っていなかったので、単に経験値が不足していただけかもしれない。

```c#
#include<iostream>
#include<vector>
#include<algorithm>
#include<string>
using namespace std;

int main() {
	int N, K;
	cin >> N >> K;
	string S;
	cin >> S;

	vector<int> Nums;
	int now = 1; //今見ている数
	int cnt = 0; //nowがいくつ並んでいるか
	for (int i = 0; i < N; i++)
	{
		if (S[i] == (char)('0' + now)) cnt++;
		else {
			Nums.push_back(cnt);
			now = 1 - now; //0と1を切り替える時の計算 now ^= 1;
			cnt = 1; //新しいのをカウントし始める
		}
	}
	if (cnt != 0) Nums.push_back(cnt);

	//1-0-1-0-1-0-1って感じの配列が欲しい
	//1-0-1-0みたいに0で終わってたら、適当に１つ足す
	if (Nums.size() % 2 == 0) Nums.push_back(0);

	int Add = 2 * K + 1;

	int ans = 0;

	int left = 0;
	int right = 0;
	int tmp = 0;// [left, right) のsum

	//1-0-1...の1から始めるので、偶数番目だけ見る
	for (int i = 0; i < Nums.size(); i += 2)
	{
		int Nextleft = i;
		int Nextright = min(i + Add, (int)Nums.size());

		//左端を移動する
		while (Nextleft > left) {
			tmp -= Nums[left];
			left++;
		}
		//右端を移動する
		while (Nextright > right) {
			tmp += Nums[right];
			right++;
		}

		ans = max(tmp, ans);
	}

	cout << ans << endl;
}
```
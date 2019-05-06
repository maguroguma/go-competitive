package ordered_set

import (
	"testing"

	"github.com/wangjia184/sortedset"
	"gotest.tools/assert"
)

func Test更新や削除(t *testing.T) {
	// 新しい set の作成
	set := sortedset.New()

	// 新しいノードの追加
	// 新規ノード追加時はtrueを、更新されたときはfalseを返す
	// key: string, score: int64, value: interface{}
	set.AddOrUpdate("a", 89, "Kelly")
	set.AddOrUpdate("b", 100, "Staley")
	set.AddOrUpdate("c", 100, "Jordon")
	set.AddOrUpdate("d", -321, "Park")
	set.AddOrUpdate("e", 101, "Albert")
	set.AddOrUpdate("f", 99, "Lyman")
	set.AddOrUpdate("g", 99, "Singleton")
	set.AddOrUpdate("h", 70, "Audrey")

	// 既存ノードのスコアのアップデート
	// **O(1)**
	actual := set.GetByKey("e")
	assert.Equal(t, actual.Value, "Albert")
	set.AddOrUpdate("e", 99, "ntrnrt")
	actual = set.GetByKey("e")
	assert.Equal(t, actual.Value, "ntrnrt")

	// keyによってノードを取得
	actual = set.GetByKey("b")
	assert.Equal(t, actual.Value, "Staley")

	// keyによってノードを削除
	// **O(log(N))**
	// 削除時はpopのように取得できる
	actual = set.Remove("b")
	assert.Equal(t, actual.Value, "Staley")
	// 存在しないkeyを取得しようとするとその型のnilが返ってくる
	actual = set.GetByKey("b")
	var expected *sortedset.SortedSetNode
	assert.Equal(t, actual, expected)
	// 存在しないkeyを削除しようとしてもpanicは発生しない
	actual = set.Remove("b")
	assert.Equal(t, actual, expected)

	// 集合中のノード数を取得する
	assert.Equal(t, set.GetCount(), 7)
}

func Testランク1(t *testing.T) {
	set := sortedset.New()

	set.AddOrUpdate("a", 89, "Kelly")
	set.AddOrUpdate("b", 100, "Staley")
	set.AddOrUpdate("c", 100, "Jordon")
	set.AddOrUpdate("d", -321, "Park")
	set.AddOrUpdate("e", 101, "Albert")
	set.AddOrUpdate("f", 99, "Lyman")
	set.AddOrUpdate("g", 99, "Singleton")
	set.AddOrUpdate("h", 70, "Audrey")

	// keyをもつノードのrank(position)を取得する
	// **O(log(N))**
	// スコアが小さいほど小さいランクが 1-based で出力される
	// 存在しないkeyを指定した場合、0が返される
	assert.Equal(t, set.FindRank("d"), 1)
	assert.Equal(t, set.FindRank("e"), 8)
	assert.Equal(t, set.FindRank("aaa"), 0)

	// スコア最小値のノードを取り出す
	// **O(log(N))**
	actual := set.PopMin()
	assert.Equal(t, actual.Key(), "d")
	assert.Equal(t, actual.Score(), sortedset.SCORE(-321))
	assert.Equal(t, int(actual.Score()), -321)
	assert.Equal(t, actual.Value, "Park")
	// ノード数は減少する
	assert.Equal(t, set.GetCount(), 7)

	// スコア最大のノードを取り出す
	// **O(1)**
	actual = set.PeekMax()
	assert.Equal(t, actual.Key(), "e")
	assert.Equal(t, actual.Score(), sortedset.SCORE(101))
	assert.Equal(t, int(actual.Score()), 101)
	assert.Equal(t, actual.Value, "Albert")
	// ノード数は変化しない
	assert.Equal(t, set.GetCount(), 7)
}

func Testランク2(t *testing.T) {
	set := sortedset.New()

	set.AddOrUpdate("a", 89, "Kelly")
	set.AddOrUpdate("b", 100, "Staley")
	set.AddOrUpdate("c", 100, "Jordon")
	set.AddOrUpdate("d", -321, "Park")
	set.AddOrUpdate("e", 101, "Albert")
	set.AddOrUpdate("f", 99, "Lyman")
	set.AddOrUpdate("g", 99, "Singleton")
	set.AddOrUpdate("h", 70, "Audrey")

	// ランク1のノード（スコア最小のノード）を取得する
	// 第2引数は削除フラグ
	// **O(lg(N))**
	actual := set.GetByRank(1, false)
	assert.Equal(t, actual.Key(), "d")
	// ノード数は変化しない
	assert.Equal(t, set.GetCount(), 8)

	// // ランク-1のノード（スコア最大のノード）を取得する
	// actual = set.GetByRank(-1, true)
	// assert.Equal(t, actual.Key(), "e")
	// // ノード数は減少する
	// assert.Equal(t, set.GetCount(), 7)

	// ランク-1のノード（スコア最大のノード）を取得する
	actual = set.GetByRank(-1, false)
	assert.Equal(t, actual.Key(), "e")
	// ノード数は減少するj
	assert.Equal(t, set.GetCount(), 8)

	// 2番目にスコアの大きいノードを取得する
	actual = set.GetByRank(-2, false)
	assert.Equal(t, actual.Key(), "c")
	actual = set.GetByRank(-3, false)
	assert.Equal(t, actual.Key(), "b")

	// [1, -1] のランクのノード（すなわち全ノード）の取得
	// **O(log(N))**
	// start, end, remove
	L := set.GetByRankRange(1, -1, false)
	assert.Equal(t, len(L), 8)
	assert.Equal(t, L[0].Key(), "d")
	assert.Equal(t, L[len(L)-1].Key(), "e")

	// スコアの大きさが2番目と3番目のランクのノードを逆順で取得
	L = set.GetByRankRange(-2, -3, false)
	assert.Equal(t, len(L), 2)
	assert.Equal(t, L[0].Key(), "c")
	assert.Equal(t, L[1].Key(), "b")
}

func Testスコア(t *testing.T) {
	set := sortedset.New()

	set.AddOrUpdate("a", 89, "Kelly")
	set.AddOrUpdate("b", 100, "Staley")
	set.AddOrUpdate("c", 100, "Jordon")
	set.AddOrUpdate("d", -321, "Park")
	set.AddOrUpdate("e", 101, "Albert")
	set.AddOrUpdate("f", 99, "Lyman")
	set.AddOrUpdate("g", 99, "Singleton")
	set.AddOrUpdate("h", 70, "Audrey")

	// スコアが [80, 100] なノードリストを取得
	// **O(log(N))**
	// start, end, options
	// オプションによって開区間の設定や、取得ノード数の制限を加えたりできる
	L := set.GetByScoreRange(80, 100, nil)
	assert.Equal(t, len(L), 5)
	assert.Equal(t, L[0].Key(), "a")
	assert.Equal(t, L[len(L)-1].Key(), "c")
}

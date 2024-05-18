package main

import (
	// 
)

// ２点間の距離は、まあそのまま求めれば良い
// 点と直線の距離は、点から直線へ垂線を引き、点から交点までの距離で算出可能
// 点と線分の距離は、
// ・点が線分の２つの端点それぞれから「90度〜-90度」の範囲に収まっていれば、垂線を引いて交点までの距離で算出可能
// ・収まっていない場合は、点から近い方の端点までの距離で算出可能
// 線分と線分の距離は、４つの端点それぞれから、もう一方の線分への距離を算出し、求まった４つの距離のうち最短のもの、で算出可能
func distance() {
	// 実装は、射影や反射をいい感じに応用するだけのように感じるので、パス
}
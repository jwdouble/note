package benchmark

import "testing"

//对于变量和内容简单的结构体，两者性能五五开
//对于极复杂结构体，性能还是五五开
func BenchmarkClassicForS(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ClassicForS()
	}
}

func BenchmarkRangeForS(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RangeForS()
	}
}

func BenchmarkClassicForB(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ClassicForB()
	}
}

func BenchmarkRangeForB(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RangeForB()
	}
}

func ClassicForS() {
	list := make([]*Small, 10000)
	for i := 0; i < 10000; i++ {
		list[i] = &Small{}
	}
}

func RangeForS() {
	list := make([]*Small, 10000)
	for i := range list {
		list[i] = &Small{}
	}
}

func ClassicForB() {
	list := make([]*Big, 10000)
	for i := 0; i < 10000; i++ {
		list[i] = &Big{
			aa: "",
			ab: "",
			ac: "",
			ad: "",
			ae: "",
			af: "",
			ag: "",
			ah: "",
			ai: "",
			aj: "",
			ak: "",
			al: "",
			am: "",
			an: "",
			ao: "",
			ap: "",
			aq: "",
			ar: "",
			as: "",
			at: "",
			au: "",
			av: "",
			aw: "",
			ax: "",
			ay: "",
			az: "",
			ba: 0,
			bb: 0,
			bc: 0,
			bd: 0,
			be: 0,
			bf: 0,
			bg: 0,
			bh: 0,
			bi: 0,
			bj: 0,
			bk: 0,
			bl: 0,
			bm: 0,
			bn: 0,
			bo: 0,
			bp: 0,
			bq: 0,
			br: 0,
			bs: 0,
			bt: 0,
			bu: 0,
			bv: 0,
			bw: 0,
			bx: 0,
			by: 0,
			bz: 0,
			ca: 0,
			cb: 0,
			cc: 0,
			cd: 0,
			ce: 0,
			cf: 0,
			cg: 0,
			ch: 0,
			ci: 0,
			cj: 0,
			ck: 0,
			cl: 0,
			cm: 0,
			cn: 0,
			co: 0,
			cp: 0,
			cq: 0,
			cr: 0,
			cs: 0,
			ct: 0,
			cu: 0,
			cv: 0,
			cw: 0,
			cx: 0,
			cy: 0,
			cz: 0,
		}
	}
}

func RangeForB() {
	list := make([]*Big, 10000)
	for i := range list {
		list[i] = &Big{
			aa: "",
			ab: "",
			ac: "",
			ad: "",
			ae: "",
			af: "",
			ag: "",
			ah: "",
			ai: "",
			aj: "",
			ak: "",
			al: "",
			am: "",
			an: "",
			ao: "",
			ap: "",
			aq: "",
			ar: "",
			as: "",
			at: "",
			au: "",
			av: "",
			aw: "",
			ax: "",
			ay: "",
			az: "",
			ba: 0,
			bb: 0,
			bc: 0,
			bd: 0,
			be: 0,
			bf: 0,
			bg: 0,
			bh: 0,
			bi: 0,
			bj: 0,
			bk: 0,
			bl: 0,
			bm: 0,
			bn: 0,
			bo: 0,
			bp: 0,
			bq: 0,
			br: 0,
			bs: 0,
			bt: 0,
			bu: 0,
			bv: 0,
			bw: 0,
			bx: 0,
			by: 0,
			bz: 0,
			ca: 0,
			cb: 0,
			cc: 0,
			cd: 0,
			ce: 0,
			cf: 0,
			cg: 0,
			ch: 0,
			ci: 0,
			cj: 0,
			ck: 0,
			cl: 0,
			cm: 0,
			cn: 0,
			co: 0,
			cp: 0,
			cq: 0,
			cr: 0,
			cs: 0,
			ct: 0,
			cu: 0,
			cv: 0,
			cw: 0,
			cx: 0,
			cy: 0,
			cz: 0,
		}
	}
}

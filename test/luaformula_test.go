package test
import (
	. "github.com/onsi/ginkgo"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"fmt"
	"time"
	"github.com/stephenlyu/goformula/formulalibrary"
	"github.com/stephenlyu/tds/period"
	"github.com/stephenlyu/goformula/datalibrary"
)

var _ = Describe("LuaMACD", func() {
	It("test", func () {
		_, data := loadJson("data.json")
		rv := function.RecordVector(data)

		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.LoadLuaFormulas("luas")

		formula := library.NewFormula("MACD", rv)
		defer formula.Destroy()

		fmt.Println("Name:", formula.GetName())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
		}

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("macd.len:", len, "data len:", rv.Len())
		for i := 0; i < len; i++ {
			r := formula.Get(i)
			fmt.Printf("%s\t%.02f\t%.02f\t%.02f\t%.02f\t%.02f\n", rv.Get(i).GetDate(), r[0], r[1], r[2], r[3], r[4])
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})

var _ = Describe("LuaDrawLine", func() {
	It("test", func () {
		_, data := loadJson("300666.SZ.json")
		rv := function.RecordVector(data)

		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.LoadLuaFormulas("luas")

		formula := library.NewFormula("DRAWLINE", rv)
		defer formula.Destroy()

		fmt.Println("Name:", formula.GetName())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
		}

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("formula.len:", len, "data len:", rv.Len())
		for i := 0; i < len; i++ {
			r := formula.Get(i)
			fmt.Printf("%d. %s\t%.02f\n", i, rv.Get(i).GetDate(), r[0])
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})

var _ = Describe("LuaDDGS", func() {
	It("test", func () {
		_, data := loadJson("300666.SZ.json")
		_, p := period.PeriodFromString("D1")
		rv := function.RecordVectorEx("300666", p, data)

		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.LoadLuaFormulas("luas")

		formula := library.NewFormula("DDGS", rv)
		defer formula.Destroy()

		fmt.Println("Name:", formula.GetName())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
		}

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("formula.len:", len, "data len:", rv.Len())
		for i := 0; i < len; i++ {
			r := formula.Get(i)
			fmt.Printf("%d. %s\t%.02f\n", i, rv.Get(i).GetDate(), r[0])
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})

var _ = Describe("LuaEMA513", func() {
	It("test", func () {
		_, data := loadJson("300666.SZ.json")
		_, p := period.PeriodFromString("D1")
		rv := function.RecordVectorEx("300666", p, data)

		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.LoadLuaFormulas("luas")

		formula := library.NewFormula("EMA513", rv)
		defer formula.Destroy()

		fmt.Println("Name:", formula.GetName())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
		}

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("formula.len:", len, "data len:", rv.Len())
		for i := 0; i < len; i++ {
			r := formula.Get(i)
			fmt.Printf("%d. %s\t%.02f\n", i, rv.Get(i).GetDate(), r[0])
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})

var _ = Describe("LuaCross", func() {
	It("test", func () {
		dl := datalibrary.NewDataLibrary("data")
		rv := dl.GetData("000001", "M5")

		fmt.Println("data len:", rv.Len())
		start := time.Now().UnixNano()

		library := formulalibrary.GlobalLibrary
		library.Reset()
		library.SetDebug(true)
		library.SetDataLibrary(dl)
		library.LoadLuaFormulas("luas")

		formula := library.NewFormula("CROSS", rv)
		defer formula.Destroy()

		fmt.Println("Name:", formula.GetName())
		for i := 0; i < formula.ArgCount(); i++ {
			min, max := formula.ArgRange(i)
			fmt.Printf("default: %f min: %f max: %f\n", formula.ArgDefault(i), min, max)
		}
		for i := 0; i < formula.VarCount(); i++ {
			fmt.Printf("name: %s noDraw: %v lineThick: %d color: %+v\n", formula.VarName(i), formula.NoDraw(i), formula.LineThick(i), formula.Color(i))
		}

		start1 := time.Now().UnixNano()
		len := formula.Len()
		fmt.Println("formula.len:", len, "data len:", rv.Len())

		for i := 0; i < formula.Len(); i++ {
			r := formula.Get(i)
			fmt.Printf("%d. %s", i, rv.Get(i).GetDate())
			for _, v := range r {
				fmt.Printf("\t%.02f", v)
			}
			fmt.Println("")
		}

		fmt.Println("time cost: ", (time.Now().UnixNano() - start) / 1000000, (time.Now().UnixNano() - start1) / 1000000, "ms")
	})
})

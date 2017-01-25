package options

import (
	"testing"
	"time"
)

func TestUtilsOptions(t *testing.T) {

	ops := New()
	ops["int64"] = 123
	ops["string"] = "abc"
	ops["bool"] = true
	ops["dur"] = "2s"
	ops["float"] = 123123.23

	if ops.Bool("bool", false) != true {
		t.Fatal("bool not proper value")
	}

	if ops.Int64("int64", 343) != 123 {
		t.Fatal("int64 not proper value")
	}

	if ops.Float64("float", .9) != 123123.23 {
		t.Fatal("bool not proper value")
	}

	if ops.Float64("MOO", .9) != 0.9 {
		t.Fatal("Float64 did not get default")
	}

	if ops.Bool("MOO", false) != false {
		t.Fatal("bool default was not gotten properly")
	}

	if _, err := ops.BoolRequired("MOO"); err == nil {
		t.Fatal("Required failed")
	}

	if _, err := ops.Float64Required("float"); err != nil {
		t.Fatal("Required failed")
	}
	if _, err := ops.Float64Required("floatsdf"); err == nil {
		t.Fatal("Required failed")
	}
	if _, err := ops.Int64Required("int64"); err != nil {
		t.Fatal("Required failed")
	}
	if _, err := ops.Int64Required("int64sdfsdf"); err == nil {
		t.Fatal("Required failed")
	}
	dd := ops.Duration("dur", time.Duration(0))
	if dd.String() != (time.Duration(2) * time.Second).String() {
		t.Fatal("Duration conversion failure")
	}
	if _, err := ops.Float64Required("int64sdfsdf"); err == nil {
		t.Fatal("Required failed")
	}

	ops.Set("monkey", 123456)
	if ops.Int64("monkey", 3) != 123456 {
		t.Fatal("set int64 failed")
	}
}

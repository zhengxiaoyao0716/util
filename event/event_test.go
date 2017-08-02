package event

import (
	"fmt"
	"testing"
)

func TestNewPool(t *testing.T) {
	p := NewPool()
	fmt.Println("Create new pool:", p.types, p.names)

	// bind listeners
	fmt.Println()

	l1 := p.On(Key{"Test", "abc"}, func(e Event) error {
		fmt.Println("\tlisten01:", e)
		return nil
	})
	fmt.Println("Bind a listener:", l1)

	l2 := &Listener{"listener02", Key{"Test", "def"}, func(e Event) error {
		fmt.Println("\tlisten02:", e)
		return nil
	}}
	p.Register(l2)
	fmt.Println("Bind a listener:", l2)

	l3 := &Listener{"listener03", Key{"Test", "123"}, func(e Event) error {
		fmt.Println("\tlisten03:", e)
		return nil
	}}
	l3.Bind(p)
	fmt.Println("Bind a listener:", l3)

	// send events
	fmt.Println()

	e1 := p.Emit(Key{"Test", "abc"}, "event01")
	fmt.Println("Send an event:", e1)
	p.Wait()

	e2 := &Event{Key{"Test", "def"}, map[string]string{"name": "event02"}}
	p.Publish(e2)
	fmt.Println("Send an event:", e2)
	p.Wait()

	e3 := &Event{Key{"Test", "123"}, 123}
	e3.Send(p)
	fmt.Println("Send an event:", e3)
	p.Wait()

	// unbind listeners
	fmt.Println()

	fmt.Println("Before unbind pool is:", p.types, p.names)
	p.Off(Key{"Test", "abc"}, l1.ID)
	p.Remove(l2)
	l3.Unbind(p)
	fmt.Println("After unbinded:", p.types, p.names)

	// resend events

	e1.Send(p)
	e2.Send(p)
	e3.Send(p)
	p.Wait()

	// tpye group
	fmt.Println()

	p.On(Key{"Group01", ""}, func(e Event) error {
		fmt.Println("\tlisten group01:", e)
		return nil
	})
	p.On(Key{"Group02", "child"}, func(e Event) error {
		t.Fatal("Unexpected handler trigger")
		return nil
	})
	fmt.Println("After group listener binded:", p.types, p.names)
	p.Emit(Key{"Group01", ""}, "group01")
	p.Emit(Key{"Group01", "child"}, "group01 child")
	p.Emit(Key{"Group02", ""}, "group02")
	// p.Emit(Key{"Group02", "child"}, "group02")
	p.Wait()

	// restrict pool
	fmt.Println()

	// ERROR: rp := NewRestrictPool([][2]string{{"Test", "abc"}, {"Test", "def"}, {"Group", ""}}...)
	rp := NewRestrictPool([]Key{{"Test", "abc"}, {"Test", "def"}, {"Group", ""}}...)
	// rp := NewRestrictPool(Key{"Test", "abc"}, Key{"Test", "def"}, Key{"Group", ""})
	// rp := NewRestrictPool([2]string{"Test", "abc"}, [2]string{"Test", "def"}, [2]string{"Group", ""})
	fmt.Println("Create restrict pool:", rp.types, rp.names)
	rp.Register(l1)
	rp.Register(l2)
	rp.Register(l3)
	rp.On(Key{"Group", ""}, func(e Event) error {
		fmt.Println("\tlisten group:", e)
		return nil
	})
	rp.Publish(e1)
	rp.Publish(e2)
	rp.Publish(e3)
	rp.Emit(Key{"Group", "child"}, nil)
	p.Wait()
	fmt.Println("Restrict pool:", rp.types, rp.names)
}

package main

import "fmt"

type CellPhone struct {
	Camera       bool
	DualSim      bool
	Torch        bool
	ColorDisplay bool
}

type BuildProcess interface {
	setCamera() BuildProcess
	setDualSim() BuildProcess
	setTorch() BuildProcess
	setColorDisplay() BuildProcess
	GetCellPhone() CellPhone
}

type Nokia struct {
	Phone CellPhone
}

func (n *Nokia) setCamera() BuildProcess {
	n.Phone.Camera = false
	return n
}

func (n *Nokia) setDualSim() BuildProcess {
	n.Phone.DualSim = false
	return n
}

func (n *Nokia) setTorch() BuildProcess {
	n.Phone.Torch = true
	return n
}

func (n *Nokia) setColorDisplay() BuildProcess {
	n.Phone.ColorDisplay = false
	return n
}

func (n *Nokia) GetCellPhone() CellPhone {
	return n.Phone
}

type Samsung struct {
	Phone CellPhone
}

func (s *Samsung) setCamera() BuildProcess {
	s.Phone.Camera = true
	return s
}

func (s *Samsung) setDualSim() BuildProcess {
	s.Phone.DualSim = true
	return s
}

func (s *Samsung) setTorch() BuildProcess {
	s.Phone.Torch = false
	return s
}

func (s *Samsung) setColorDisplay() BuildProcess {
	s.Phone.ColorDisplay = true
	return s
}

func (s *Samsung) GetCellPhone() CellPhone {
	return s.Phone
}

type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}

func (d *Director) Construct() CellPhone {
	d.builder.setCamera().setDualSim().setTorch().setColorDisplay()
	return d.builder.GetCellPhone()
}

func main() {
	diro := Director{}
	nokiaPhone := &Nokia{}
	diro.SetBuilder(nokiaPhone)
	phone := diro.Construct()
	printProduct(phone)

	samsungPhone := &Samsung{}
	diro.SetBuilder(samsungPhone)
	phone = diro.Construct()
	printProduct(phone)
}

func printProduct(phone CellPhone) {
	fmt.Println("Phone has camera? ", getAns(phone.Camera))
	fmt.Println("Phone has Dual Sim? ", getAns(phone.DualSim))
	fmt.Println("Phone has Torch? ", getAns(phone.Torch))
	fmt.Println("Phone has Color Display? ", getAns(phone.ColorDisplay))
}

func getAns(b bool) string {
	if b {
		return "YES"
	}
	return "NO"
}

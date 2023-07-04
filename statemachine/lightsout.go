package statemachine

import (
	"fmt"
	"study-statemachine/light"
)

type LightsOut struct {
	mc *MachineContext
}

var (
	_ EventListener = (*LightsOut)(nil)
	_ State         = (*LightsOut)(nil)
)

func (s LightsOut) Entry() {
	fmt.Println("LightsOut#Entry")

	s.Do() // Doの処理は非同期実行だけど割愛
}

func (s LightsOut) Do() {
	light.BlueLighting()
	light.RedLighting()
}

func (s LightsOut) Exit() {
	// 本当はDoをキャンセルする処理が必用

	fmt.Println("LightsOut#Exit")
}

func (l LightsOut) OnStart() {
	l.mc.changeState(&BlueLighting{mc: l.mc})
}

func (l LightsOut) OnStop() {
	fmt.Println("Power off")

	l.mc.terminate()
}

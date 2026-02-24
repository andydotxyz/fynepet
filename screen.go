package main

import (
	"time"

	"fyne.io/fyne/v2"
	"github.com/andydotxyz/fynepet/pkg/fynepet"
)

var (
	frameBlack = fynepet.Pixels{4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295}

	frameEggDown = fynepet.Pixels{0, 0, 0, 0, 516096, 946176, 1468416, 1677312, 3775488, 2515968, 2515968, 3775488, 1677312, 946176, 2095104, 0}
	frameEggUp   = fynepet.Pixels{0, 0, 245760, 417792, 946176, 626688, 1677312, 1468416, 2515968, 3775488, 3775488, 1468416, 946176, 368640, 1044480, 0}

	frameBaby1 = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 245760, 270336, 544768, 659456, 528384, 626688, 528384, 516096}
	frameBaby2 = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 61440, 67584, 136192, 164864, 144384, 261120}

	frameDead1 = fynepet.Pixels{33280, 545792, 402919440, 635726888, 1108873232, 1166562304, 1166311936, 1144258560, 573833216, 554172416, 284426308, 134348816, 100794408, 25432080, 7874628, 4096}
	frameDead2 = fynepet.Pixels{33280, 268456960, 538447888, 665087016, 406982672, 295851008, 564298240, 740425728, 742653952, 277086208, 789053508, 1074790416, 1080033320, 562040848, 503326788, 4096}

	sleepDark1 = fynepet.Pixels{4294963711, 4294966783, 4294966271, 4294965247, 4294955519, 4294934527, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295}
	sleepDark2 = fynepet.Pixels{4294967295, 4294967295, 4294936575, 4294965247, 4294963199, 4294959103, 4294950911, 4294936575, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295, 4294967295}

	sleepLight1 = fynepet.Pixels{14, 2, 4, 8, 46, 128, 0, 0, 0, 0, 0, 1032192, 1056768, 1892352, 1056768, 2088960}
	sleepLight2 = fynepet.Pixels{0, 0, 120, 8, 16, 32, 64, 120, 0, 0, 0, 1032192, 1056768, 1892352, 1056768, 2088960}

	feedMenu1 = fynepet.Pixels{0, 268450816, 402670080, 2080391680, 2113994496, 2080440064, 402670080, 268450816, 3072, 7680, 27904, 33024, 65280, 65280, 33024, 65280}
	feedMenu2 = fynepet.Pixels{0, 15360, 16896, 16896, 65280, 65280, 16896, 15360, 3072, 268443136, 402681088, 2080407808, 2113994496, 2080440064, 402686208, 268500736}

	lightMenuOn  = fynepet.Pixels{0, 272130560, 407130624, 2084852224, 2118404608, 2084850176, 406341120, 268435456, 0, 3796736, 4524032, 4574720, 4524032, 4524032, 3737600, 0}
	lightMenuOff = fynepet.Pixels{0, 3695104, 4477440, 4477440, 4475392, 4475392, 3687936, 0, 0, 272232192, 407177216, 2084949504, 2118453248, 2084898816, 406390784, 268435456}

	statsMenuHome       = fynepet.Pixels{0, 1006632960, 1509949525, 2113929302, 1711276116, 2113929268, 1006632980, 96, 0, 1090519040, 2130706512, 134217808, 1040187480, 570425428, 704643156, 2130706520}
	statsMenuDiscipline = fynepet.Pixels{1879050240, 1207961600, 1255860974, 1217439914, 1255320238, 1246948008, 1926924974, 32768, 1073774593, 1073741822, 1073741825, 1426063361, 1426063361, 1073741825, 1073741822, 0}
	statsMenuHungry     = fynepet.Pixels{2415919104, 2415919104, 2507625728, 4116014336, 2505393408, 2505523968, 2538684672, 460288, 0, 1819044972, 3197276818, 3196224130, 4269965954, 2084848708, 942155816, 269488144}
	statsMenuHappy      = fynepet.Pixels{1207959552, 1207959552, 1238297088, 2018159104, 1271572992, 1246684672, 1235780096, 560128, 0, 1819044972, 3197276818, 3196224130, 4269965954, 2084848708, 942155816, 269488144}

	waste1 = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 1, 66, 129, 80, 24, 52, 122, 126}
	waste2 = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 128, 66, 129, 18, 24, 52, 94, 126}

	// Food item on the left side of screen (burger-like shape)
	feedBurger = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1006632960, 2080374784, 2080374784, 1006632960, 0}
	// Food item on the left side of screen (cake-like shape)
	feedCake = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 134217728, 939524096, 2080374784, 2080374784, 2080374784, 0}

	// Pet eating frames (mouth open/closed)
	feedEat1 = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 245760, 270336, 544768, 659456, 528384, 659456, 528384, 516096}
	feedEat2 = fynepet.Pixels{0, 0, 0, 0, 0, 0, 0, 0, 245760, 270336, 544768, 659456, 528384, 528384, 528384, 516096}
)

func (u *ui) feedAnimation(isMeal bool) {
	food := feedCake
	if isMeal {
		food = feedBurger
	}

	// Show food approaching from the left (shift food right over 8 frames)
	for i := 0; i < 8; i++ {
		time.Sleep(time.Millisecond * 100)
		shift := 7 - i
		fyne.Do(func() {
			var frame fynepet.Pixels
			for row := 0; row < 16; row++ {
				frame[row] = food[row] >> shift
			}
			// Overlay pet on the right side
			for row := 0; row < 16; row++ {
				frame[row] |= frameBaby1[row]
			}
			u.scr.SetPixels(frame)
		})
	}

	// Eating animation: alternate mouth open/closed 3 times
	for i := 0; i < 6; i++ {
		time.Sleep(time.Millisecond * 150)
		frame := feedEat1
		if i%2 == 1 {
			frame = feedEat2
		}
		fyne.Do(func() {
			u.scr.SetPixels(frame)
		})
	}

	// Brief pause after eating
	time.Sleep(time.Millisecond * 200)
}

func (u *ui) cleanAnimation() {
	i := 0
	for i < 36 {
		time.Sleep(time.Millisecond * 50)

		fyne.Do(func() {
			if i < 32 { // last 4 frames are a small freeze on the device
				u.scr.ScrollLeft()
				for id := 0; id < 16; id++ {
					switch i {
					case 0:
						if id%4 == 2 {
							u.scr.TogglePixel(31, id)
						}
					case 1:
						if id%4 != 0 {
							u.scr.TogglePixel(31, id)
						}
					case 2:
						if id%4 != 2 {
							u.scr.TogglePixel(31, id)
						}
					case 3:
						if id%2 == 0 {
							u.scr.TogglePixel(31, id)
						}
					case 4:
						if id%2 == 1 {
							u.scr.TogglePixel(31, id)
						}
					case 5:
						if id%4 == 0 {
							u.scr.TogglePixel(31, id)
						}
					}
				}
			}
			i++
		})
	}
}

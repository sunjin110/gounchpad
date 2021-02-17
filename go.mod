module gounchpad

go 1.15

// require github.com/faiface/beep v1.0.2

replace github.com/faiface/beep => ./lib/beep

require (
	github.com/faiface/beep v0.0.0-00010101000000-000000000000
	github.com/franela/goblin v0.0.0-20210113153425-413781f5e6c8
	github.com/gdamore/tcell v1.4.0
)

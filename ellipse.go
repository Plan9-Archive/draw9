package draw9

import "image"

func (dst *Image) Ellipse(src *Image, center image.Point, a int, b int, radius int, sp image.Point, alpha int, phi int) {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	dst.ellipseOp(src, center, a, b, radius, sp, alpha, phi, SoverD)
}

func (dst *Image) EllipseOp(src *Image, center image.Point, a int, b int, radius int, sp image.Point, alpha int, phi int, op Op) {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	dst.ellipseOp(src, center, a, b, radius, sp, alpha, phi, op)
}

func (dst *Image) EllipseFill(src *Image, center image.Point, a int, b int, radius int, sp image.Point, alpha int, phi int) {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	dst.ellipseFillOp(src, center, a, b, radius, sp, alpha, phi, SoverD)
}

func (dst *Image) EllipseFillOp(src *Image, center image.Point, a int, b int, radius int, sp image.Point, alpha int, phi int, op Op) {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	dst.ellipseFillOp(src, center, a, b, radius, sp, alpha, phi, op)
}

func (dst *Image) ellipseOp(src *Image, center image.Point, a int, b int, radius int, sp image.Point, alpha int, phi int, op Op) {
	setdrawop(dst.Display, op)
	ar := dst.Display.bufimage(1 + 4 + 4 + 2*4 + 4 + 4 + 4 + 2*4 + 4 + 4)
	ar[0] = 'e'
	bplong(ar[1:], uint32(dst.ID))
	bplong(ar[5:], uint32(src.ID))
	bplong(ar[9:], uint32(center.X))
	bplong(ar[13:], uint32(center.Y))
	bplong(ar[17:], uint32(a))
	bplong(ar[21:], uint32(b))
	bplong(ar[25:], uint32(radius))
	bplong(ar[29:], uint32(sp.X))
	bplong(ar[33:], uint32(sp.Y))
	bplong(ar[37:], uint32(alpha))
	bplong(ar[41:], uint32(phi))
}

func (dst *Image) ellipseFillOp(src *Image, center image.Point, a int, b int, radius int, sp image.Point, alpha int, phi int, op Op) {
	setdrawop(dst.Display, op)
	ar := dst.Display.bufimage(1 + 4 + 4 + 2*4 + 4 + 4 + 4 + 2*4 + 4 + 4)
	ar[0] = 'E'
	bplong(ar[1:], uint32(dst.ID))
	bplong(ar[5:], uint32(src.ID))
	bplong(ar[9:], uint32(center.X))
	bplong(ar[13:], uint32(center.Y))
	bplong(ar[17:], uint32(a))
	bplong(ar[21:], uint32(b))
	bplong(ar[25:], uint32(radius))
	bplong(ar[29:], uint32(sp.X))
	bplong(ar[33:], uint32(sp.Y))
	bplong(ar[37:], uint32(alpha))
	bplong(ar[41:], uint32(phi))
}

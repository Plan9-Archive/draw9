package draw9

import (
	"image"
	"log"
)

func (dst *Image) String(pt image.Point, src *Image, sp image.Point, f *Font, s string) image.Point {
	if dst == nil {
		panic("nil dst")
	}
	if dst.Display == nil {
		panic("nil dst display")
	}
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	return _string(dst, pt, src, sp, f, s, nil, nil, dst.Clipr, nil, image.ZP, SoverD)
}

func (dst *Image) StringOp(pt image.Point, src *Image, sp image.Point, f *Font, s string, op Op) image.Point {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	return _string(dst, pt, src, sp, f, s, nil, nil, dst.Clipr, nil, image.ZP, op)
}

func (dst *Image) Runes(pt image.Point, src *Image, sp image.Point, f *Font, r []rune) image.Point {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	return _string(dst, pt, src, sp, f, "", nil, r, dst.Clipr, nil, image.ZP, SoverD)
}

func (dst *Image) RunesOp(pt image.Point, src *Image, sp image.Point, f *Font, r []rune, op Op) image.Point {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	return _string(dst, pt, src, sp, f, "", nil, r, dst.Clipr, nil, image.ZP, op)
}

func (dst *Image) Bytes(pt image.Point, src *Image, sp image.Point, f *Font, b []byte) image.Point {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	return _string(dst, pt, src, sp, f, "", b, nil, dst.Clipr, nil, image.ZP, SoverD)
}

func (dst *Image) BytesOp(pt image.Point, src *Image, sp image.Point, f *Font, b []byte, op Op) image.Point {
	dst.Display.mu.Lock()
	defer dst.Display.mu.Unlock()
	return _string(dst, pt, src, sp, f, "", b, nil, dst.Clipr, nil, image.ZP, op)
}

func _string(dst *Image, pt image.Point, src *Image, sp image.Point, f *Font, s string, b []byte, r []rune, clipr image.Rectangle, bg *Image, bgp image.Point, op Op) image.Point {
	if dst.Display != nil && dst.Display.debug {
		log.Printf("_string %s %q", f, s)
	}
	var in input
	in.init(s, b, r)
	const Max = 100
	cbuf := make([]uint16, Max)
	var sf *Subfont
	for !in.done {
		max := Max
		n, wid, subfontname := cachechars(f, &in, cbuf, max)
		if dst.Display != nil && dst.Display.debug {
			log.Printf("cachechars %d %d %q", n, wid, subfontname)
		}
		if n > 0 {
			setdrawop(dst.Display, op)
			m := 47 + 2*n
			if bg != nil {
				m += 4 + 2*4
			}
			b := dst.Display.bufimage(m)
			if bg != nil {
				b[0] = 'x'
			} else {
				b[0] = 's'
			}
			bplong(b[1:], uint32(dst.ID))
			bplong(b[5:], uint32(src.ID))
			bplong(b[9:], uint32(f.cacheimage.ID))
			bplong(b[13:], uint32(pt.X))
			bplong(b[17:], uint32(pt.Y+f.Ascent))
			bplong(b[21:], uint32(clipr.Min.X))
			bplong(b[25:], uint32(clipr.Min.Y))
			bplong(b[29:], uint32(clipr.Max.X))
			bplong(b[33:], uint32(clipr.Max.Y))
			bplong(b[37:], uint32(sp.X))
			bplong(b[41:], uint32(sp.Y))
			bpshort(b[45:], uint16(n))
			b = b[47:]
			if bg != nil {
				bplong(b, uint32(bg.ID))
				bplong(b[4:], uint32(bgp.X))
				bplong(b[8:], uint32(bgp.Y))
				b = b[12:]
			}
			for i, c := range cbuf[:n] {
				bpshort(b[2*i:], c)
			}
			pt.X += wid
			bgp.X += wid
			agefont(f)
		}
		if subfontname != "" {
			sf.free()
			var err error
			sf, err = getsubfont(f.Display, subfontname)
			if err != nil {
				if f.Display != nil && f != f.Display.DefaultFont {
					f = f.Display.DefaultFont
					continue
				}
				break
			}
			/*
			 * must not free sf until cachechars has found it in the cache
			 * and picked up its own reference.
			 */
		}
	}
	return pt
}

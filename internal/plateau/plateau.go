package plateau

type Plateau struct {
   MaxX int
   MaxY int
}

// New creates a new Plateau instance.
func New(maxX, maxY int) *Plateau {
   return &Plateau{
      MaxX: maxX,
      MaxY: maxY,
   }
}

func (p *Plateau) IsinBounds(x, y int) bool {
   return x >= 0 && y >= 0 && x <= p.MaxX && y <= p.MaxY
}

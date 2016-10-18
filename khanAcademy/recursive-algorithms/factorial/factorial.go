package factorial

import "math"

func IterativeFactorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}

func RecursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * RecursiveFactorial(n-1)
}

type Palindrome struct {
	Str string
}

func NewPalindrome() *Palindrome {
	return &Palindrome{Str: ""}
}

func (p *Palindrome) firstCharacter() string {
	return p.Str[:1]
}
func (p *Palindrome) lastCharacter() string {
	return p.Str[len(p.Str)-1:]
}
func (p *Palindrome) middleCharacters() string {
	return p.Str[1 : len(p.Str)-1]
}

func (p *Palindrome) IsPalindrome() bool {
	if len(p.Str) <= 1 {
		return true
	}

	if p.firstCharacter() == p.lastCharacter() {
		p.Str = p.middleCharacters()
		return p.IsPalindrome()
	}

	return false
}

type Power struct {
}

func NewPower() *Power {
	return &Power{}
}

func (p *Power) GetPower(x, n float64) float64 {
	if n == 0.0 {
		return 1
	}

	if n < 0.0 {
		return 1 / p.GetPower(x, -n)
	}

	if math.Mod(n, 2) == 0.0 {
		y := p.GetPower(x, n/2)
		return y * y
	}

	if math.Mod(n, 2) != 0.0 {
		return x * p.GetPower(x, n-1)
	}

	return 1
}

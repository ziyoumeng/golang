package stack

func calculate(s string) int {
	opStack :=NewSequentialStack()
	for i := 0; i < len(s); i++ {
		switch s[i]	{
		case '+','-','(',')' :
			if s[i] != ')'{

			}
		}
	}
}

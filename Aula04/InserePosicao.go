Procedimento InserePositivo(L, n, M, valor, pos):
	se n = M entao "Overflow"
	se Busca(L, n, valor) = -1 entao:
		i := n
		enquanto i != pos, faça:
			L[i] := L[i - 1]
			i := i - 1
		l[pos] := valor
		n := n + 1
	senao:
		"elemento já inserido"
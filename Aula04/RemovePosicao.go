Procedimento Remoção(L, n, valor)
	se n = 0, então "Underflow"
	pos := Busca(L, n, valor)
	se pos = -1, entao "Elemento nao existe"
	senao:
		enquanto pos > n - 1, faça:
			L[pos] = L[pos + 1]
			pos := pos + 1
		Limpa L[n]
		n := n - 1
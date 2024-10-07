sem Empleado = 0
sem Autos[20] = ([20] 0 )
cola colaAutos
sem mutex = 1
int verificacion[20] = ([20] -1)
sem AvisarSalida = 0

Process Autos[id=0..19]{
	int verificacionObtenida
	P(mutex)
	colaAutos.push(id)
	V(mutex)

	V(Empleado)
	P(Autos[id])
	verificacionObtenida = verificacion[id]
	V(AvisarSalida)
}

Process Empleado{
	int autoId,verificado
	for i=0..19{
		P(Empleado)
		P(mutex)
		autoId = colaAutos.pop()
		V(mutex)
		verificado = verificar(autoId)
		verificacionObtenida[autoId] = verificado
		V(Autos[autoId])
		P(AvisarSalida)

	}
}
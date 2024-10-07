
Monitor Oficina {
	txt resultado
	cond cv
	int esperando = 0
	cond Empleado
	cola colaSolicitudes;

	Procedure llegada(solicitud: in txt,resultadoTramite: out txt){
		esperando++
		colaSolicitudes.push(solicitud)
		signal(Empleado)
		wait(cv)
		resultadoTramite = resultado
	}

	Procedure chequearPersonas(libre: out boolean){
		if (esperando == 0){
			libre = true
		} else {
			libre = false
		}

	}

	Procedure atenderPersona(solicitud: out txt){
		solicitud = colaSolicitudes.pop()
		esperando --
	}

	Procedure entregarResultado(resultado: in txt){
		resultadoTramite = resultado
		signal(cv)
	}

}


Process Personas[id=0..N-1]{
	txt resultadoTramite,solicitud
	Oficina.llegada(solicitud,resultadoTramite)

}

Process Empleado{
	boolean libre;
	txt solicitud,resultado
	for id=0..N-1{
		Oficina.chequearPersonas(libre)
		while(libre = true){
			delay(600000)
			Oficina.chequearPersonas(libre)
		}
		Oficina.atenderPersona(solicitud)
		resultado = resolverTramite(solicitud)
		Oficina.entregarResultado(resultado)

	}

}
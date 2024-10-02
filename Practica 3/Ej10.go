 

Monitor Parque{

	cond cola;
	int contador = 0;
	bool libre = true;
	cond empleado;
	cond persona;

	bool habilitado = true;

	Procedure llegada(){
		contador++;
		if (habilitado = false){
			wait(cola);
		} else {
			habilitado = false;
			signal(empleado)
			wait(cola);
		}
	
	}

	Procedure empleado_espera(){
		if(contador == 0){wait(empleado)}
	}

	Procedure avisar_juego(){
		signal(cola);
		wait(empleado);
	}

	Procedure terminar_juego(){
		if (contador > 0){
			contador--;
		} else {
			habilitado = true;
		}
		signal(empleado)
	}
}

Process Persona[id:0..N-1]{
	Parque.llegada();
	Usar_Juego();
	Parque.terminar_juego();
}

Process Empleado{
	for[i:0..N-1]{
		Parque.empleado_espera();
		Desinfectar_Juego();
		Parque.avisar_juego();
	}
}
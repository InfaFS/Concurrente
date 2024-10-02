
Monitor Eleccion {
	colaOrdenada fila;
	int esperando = 0;
	int idAux = 0;
	cond cv[N];
	cond cv_autoridad;
	cond fin;

	Procedure llegada(id:in int,tipo: in txt){
		insertar(fila,id,tipo)
		esperando++; //no seria necesario ya que usamos una cola y no cond para ver, por ende si podemos preguntar por el empty
		signal(autoridad)
		wait(cv[id])
	}

	Procedure avisar(){
		if (esperando == 0){wait (cv_autoridad)}
		esperando--;
		sacar(fila,idAux)
		signal(cv[idAux])
		wait(fin)
	}

	Procedure terminar(){
		signal(fin)
	}

}


Process personas[id:0..N-1]{
	Eleccion.llegada(id,tipo)
	Votar()
	Eleccion.terminar()
}

Process autoridad{
	for i:0..N-1{
		Eleccion.avisar()
	}
}
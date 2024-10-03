
Monitor E_S{
	cond cv[20] //Lo hacemos con 20 y no una cola normal ya que tenemos una prioridad
	boolean libre = true
	colaOrdenada cola
	int idAux
	procedure llegada(processId: in int, prioridad: in int){
		if (libre == true){
			libre = false
		} else {
			insertar(cola,processId,prioridad)
			await(cv[processId])
		}
	}

	procedure terminar(){
		if (cola.empty()){
			libre = true
		} else {
			sacar(cola,idAux)
			signal(cv[idAux])
		}
	}


}


Process procesos[id=0..19]{
	while(true){
		//hace actividades
		resultados = Procesar()
		E_S.llegada()
		buffer.escribir(resultados)
		E_S.terminar()
	}
}
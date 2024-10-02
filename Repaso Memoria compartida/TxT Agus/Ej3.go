Monitor Paso{
	cond cv;
	boolean libre = true
	int esperando = 0;

	Procedure llegada(){
		if (libre){
			libre = false
		} else {
			esperando++
			wait(cv)
		}
	}

	Procedure terminar(){
		if (esperando > 0){
			esperando --
			signal(cv)
		} else {
			libre = true
		}
	}

}



Process Autos[id=0..N-1]{
	Paso.llegada()
	//pasar
	Paso.terminar()
}
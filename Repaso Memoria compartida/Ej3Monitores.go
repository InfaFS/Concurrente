

Monitor Paso{

	bool libre = true;
	int esperando = 0;
	cond cola;
	
	Procedure llegada() {
		if(not libre){
			esperando++
			wait(cola)
		} else {
			libre = false;
		}
	}

	Procedure terminar() {
		if (esperando > 0){
			esperando--
			signal(cola)
		} else {
			libre = true;
		}
	}

}



Process escaladores[id:0..29]{
	//escalan
	Paso.llegada()
	//pasan el paso (xd)
	Paso.terminar()
}
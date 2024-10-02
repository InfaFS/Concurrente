
Monitor Equipo[id:0..4]{
	int equipoArribados = 0;
	int totalVendido = 0;
	cond espera;

	Procedure listo(){
		equipoArribados ++;
		if (equipoArribados < 4){
			wait(espera)
		} else {
			signal_all(espera)
			equipoArribados = 0
		}
	}

	Procedure terminar(vendido: out int){
		equipoArribados++;
		if (equipoArribados < 4){
			wait(espera)
		} else {
			signal_all(espera)1
		}
		vendido = totalVendido;
	}

}

Process vendedor[id:0..19]{
	int id_grupo,vendido;
	Equipo[id_grupo].listo()
	//venden
	Equipo[id].terminar(vendido)
}
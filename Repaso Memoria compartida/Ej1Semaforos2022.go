colaAutos colaLlegada;
colaAutos colaLlegadaEstacion[7];

sem coordinador = 0;

sem esperaCoordinador[150] = ([150] 0)
sem esperaEstacion[150] = ([150] 0)

sem estacion[7] = ([7] 0)

sem mutex_colaLlegada  = 1
sem mutex_colaLlegadaEstacion[7] = ([7] 1)
int estacionAsignada[150] = ([150] -1)
int resultadVehiculo[150] = ([150] -1)

int vehiculosPorEstacion[7] = ([7] 0)
sem mutexContadorVehiculos = 1;

Process vehiculo[id=0..149]{
	int estacion;
	
	P(mutex_colaLlegada)
	colaLlegada.push(id)
	V(mutex_colaLlegada)

	V(coordinador)
	P(esperaCoordinador[id])
	estacion = estacionAsignada[id]
	
	P(mutex_colaLlegadaEstacion[estacion])
	vehiculosPorEstacion[estacion]++
	colaLlegadaEstacion[estacion].push(id)
	V(mutex_colaLlegadaEstacion[estacion])

	V(estacion[estacion])
	P(esperaEstacion[id])
	resultado = resultadVehiculo[id] 
}

Process coordinador{
	int idVehiculo
	int estacionMenor
	for i=0..149{
		P(coordinador)
		P(mutex_colaLlegada)
		idVehiculo = colaLlegada.pop()
		V(mutex_colaLlegada)

		P(mutexContadorVehiculos)
		estacionMenor = vehiculosPorEstacion.min()
		V(mutexContadorVehiculos)

		estacionAsignada[idVehiculo] = estacionMenor
		V(esperaCoordinador[idVehiculo])
	}

}

Process estacion[id=0..6]{
	int vehiculoId,resultado;
	while(true){
		P(estacion[id])
	
		P(mutex_colaLlegadaEstacion[id])
		vehiculosPorEstacion[id]--
		vehiculoId = colaLlegadaEstacion[id].pop()
		V(mutex_colaLlegadaEstacion[id])

		resultado = revisar(vehiculoId)
		resultadVehiculo[vehiculoId] = resultado
		V[esperaEstacion[vehiculoId]] 
	}
}


//------------------------------------------------------------------\\
//Basado en la resolucion. Diferencias
//El coordinador incrementa el valor de cuantos vehiculos hay por lugar
//El vehiculo se decrementa el valor cuando sale

colaAutos colaLlegada;
colaAutos colaLlegadaEstacion[7];

sem coordinador = 0;

sem esperaCoordinador[150] = ([150] 0)
sem esperaEstacion[150] = ([150] 0)

sem estacion[7] = ([7] 0)

sem mutex_colaLlegada  = 1
sem mutex_colaLlegadaEstacion[7] = ([7] 1)
int estacionAsignada[150] = ([150] -1)
int resultadVehiculo[150] = ([150] -1)

int vehiculosPorEstacion[7] = ([7] 0)
sem mutexContadorVehiculos = 1;

Process vehiculo[id=0..149]{
	int estacion;
	
	P(mutex_colaLlegada)
	colaLlegada.push(id)
	V(mutex_colaLlegada)

	V(coordinador)
	P(esperaCoordinador[id])
	estacion = estacionAsignada[id]
	
	P(mutex_colaLlegadaEstacion[estacion])
	vehiculosPorEstacion[estacion]++
	colaLlegadaEstacion[estacion].push(id)
	V(mutex_colaLlegadaEstacion[estacion])

	V(estacion[estacion])
	P(esperaEstacion[id])
	resultado = resultadVehiculo[id] 
	
	P(mutexContadorVehiculos)
	vehiculosPorEstacion[id]--
	V(mutexContadorVehiculos)
}

Process coordinador{
	int idVehiculo
	int estacionMenor
	for i=0..149{
		P(coordinador)
		P(mutex_colaLlegada)
		idVehiculo = colaLlegada.pop()
		V(mutex_colaLlegada)

		P(mutexContadorVehiculos)
		estacionMenor = vehiculosPorEstacion.min()
		vehiculosPorEstacion[id]++
		V(mutexContadorVehiculos)

		estacionAsignada[idVehiculo] = estacionMenor
		V(esperaCoordinador[idVehiculo])
	}

}

Process estacion[id=0..6]{
	int vehiculoId,resultado;
	while(true){
		P(estacion[id])
	
		P(mutex_colaLlegadaEstacion[id])
		vehiculoId = colaLlegadaEstacion[id].pop()
		V(mutex_colaLlegadaEstacion[id])

		resultado = revisar(vehiculoId)
		resultadVehiculo[vehiculoId] = resultado
		V[esperaEstacion[vehiculoId]] 
	}
}
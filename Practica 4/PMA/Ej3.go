/*

4 procesos

Cliente -> manda los pedidos al channel pedidos, espera que le den su pedido en su channel privado

Coordinador entre cliente y empleado -> recibe pedido de empleado, si no esta empty el channel pedidos de los clientes
le da al empleado en su channel privador un pedido, sino le indica -1 o vacio, para que haga delay por 1-3min

Empleado -> manda pedido al coordinador, se queda en el receive, si obtiene un pedido que no sea vacio/-1 entonces se lo mandan al channel de los cocineros(uno global),sino se duerme

Cocinero -> esta todo el tiempo en un receive con el channel que tiene el id de la persona y el pedido, lo hace y se lo da en su channel privado a cada persona

*/

//Variables
chan enviarPedido(int,txt)
chan pedidoListo[C] (txt)
chan pedidoEmpleado (int)
chan enviarPedidoEmpleado[3] (int,txt)
chan pedidosCocineros (txt)

Process Cliente[id:0..C-1]{
	send enviarPedido(id,pedido);
	receive pedidoListo[id](pedido_cocinado);

}

Process Admin{
	while(true){
		receive pedidoEmpleado(idE)
		if(!empty(enviarPedido)){
			receive enviarPedido(idCliente,pedidoCliente)
			send enviarPedidoEmpleado[idE](idCliente,pedidoCliente)
		} else {
			send enviarPedidoEmpleado[idE](-1,null)
		}
	}
}

Process Empleado[id:0..2]{
	while(true){
		send pedidoEmpleado(id)
		receive enviarPedidoEmpleado[id] (idCliente,pedidoCliente)
		if (idCliente <> -1){
			send pedidosCocineros(idCliente,pedidoCliente)
		} else {
			delay(random(1..3))
		}

	}
}

Process Cocinero[id:0..1]{
	receive pedidosCocineros(idCliente,pedidoCliente)
	//hace el pedido
	send pedidoListo[idCliente] (pedidoCliente)
}
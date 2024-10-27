chan pedidoVendedor(int);
chan enviarPedido(texto);
chan entregaPedido[C] (texto);
chan recibirPedidoVendedor[V](text);

Process Cliente[id:0..C-1]{
	pedido texto;
	resultado texto;
	send enviarPedido(pedido,idC);
	receive entregaPedido[id](resultado);
}

Process Admin{
	int idV,idC;
	texto pedido;
	while(true){
		receive pedidoVendedor(idV);
		if (!empty(enviarPedido)){
			receive enviarPedido(pedido,idC);
		} else {
			pedido = "NO HAY PEDIDOS";
			idC = -1;
		}
		send recibirPedidoVendedor[idV] (pedido,idC);
	}
}

Process Vendedor[id:0..V-1]{
	texto pedido,res;
	int idC;
	while(true){
		send pedidoVendedor(id);
		receive pedidoVendedor[id] (pedido,idC);
		if (pedido != "NO HAY PEDIDOS"){
			res = resolverPedido(pedido);
			send entregaPedido[idC](res);
		} else {
			//busco respuestos
		}
	}
}


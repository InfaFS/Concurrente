/*

4 procesos

Cliente -> manda los pedidos al channel pedidos, espera que le den su pedido en su channel privado

Coordinador entre cliente y empleado -> recibe pedido de empleado, si no esta empty el channel pedidos de los clientes
le da al empleado en su channel privador un pedido, sino le indica -1 o vacio, para que haga delay por 1-3min

Empleado -> manda pedido al coordinador, se queda en el receive, si obtiene un pedido que no sea vacio/-1 entonces se lo mandan al channel de los cocineros(uno global),sino se duerme

Cocinero -> esta todo el tiempo en un receive con el channel que tiene el id de la persona y el pedido, lo hace y se lo da en su channel privado a cada persona

*/
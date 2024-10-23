--NO sabria como hacer para que hayan entries personalizadas dentro del admin ,PREGUNTAR
procedure Ej8 is

    --Workers le avisan que tienen que realizar un calculo
    Task type Cliente is
        Entry Atendido();
    End Cliente;

    --Definicion del admin que avisa
    Task Admin is
        Entry PedidoCliente();
        Entry PedidoCamion(IDCliente: OUT integer);
        Entry LlegadaCamion(IDCliente: IN integer);
    End Admin;

    Task Type Camion;
 

    arrCamiones: array(1..3) of Camion;

    arrCliente: array(1..P) of Cliente;

    
    --Clientes
    Task Body Cliente is

    begin
        while not atendido loop
            SELECT
                admin.PedidoCliente();
                accept Atendido() do
                    atendido := true;
                end Atendido;
            OR delay 900.0;
                null;
            END SELECT
        end loop

    end Cliente;

    --Admin
    Task Body Admin is

    begin
        loop
            SELECT
                accept PedidoCamion(IDCliente : OUT integer) do
                    --se fija cual es el cliente que mas pedidos tiene, lo haria con un count de entries pero no se como hacerlo en este caso
                    IDCliente := clienteMasPedidos;
                end PedidoCamion;
            OR --Aca se quedaria tildado si por ejemplo ningun camion fue a buscar un pedido
                when (LlegadaCamion'count > 0) =>
                    accept LlegadaCamion(IDCliente: IN integer) do
                        clienteAtender := IDCliente;
                    end LlegadaCamion;
                    --acepta el pedido del cliente
                    arrCliente(IDCliente).Atendido()
            END SELECT;


        end loop 
      
    End Admin;

    Task Body Camion is

    begin
        loop
        admin.PedidoCamion(IDCliente);
        --voy hasta el lugar
        admin.LlegadaCamion(IDCliente);
        end loop;
    end Camion;

begin
    null;
end Ej8;


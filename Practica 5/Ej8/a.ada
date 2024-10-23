
procedure Ej8 is

    --Workers le avisan que tienen que realizar un calculo
    Task type Cliente is
        
    End Cliente;

    --Definicion del admin que avisa
    Task Admin is

    End Admin;

    Task Type Camion is

    End Camion;
 

    arrCamiones: array(1..3) of Camion;

    arrCliente: array(1..P) of Cliente;

    
    --Clientes
    Task Body Cliente is

    begin
        while not atendido loop
            SELECT
                admin.PedidoCliente();
                atendido := true;
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
                    --se fija cual es el cliente que mas pedidos tiene
                    IDCliente := clienteMasPedidos;
                end PedidoCamion;
            OR
                accept LlegadaCamion(IDCliente: IN integer) do
                    clienteAtender := IDCliente;
                end LlegadaCamion;
                --acepta el pedido del cliente
                accept PedidoCliente
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

